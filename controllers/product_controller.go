package controllers

import (
	"fmt"
	"log"
	"net/http"
	"product-management-system/config"
	"product-management-system/models"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/rabbitmq/amqp091-go"
)

func PublishToQueue(queueName string, message string) error {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a RabbitMQ channel: %v", err)
	}
	defer ch.Close()

	err = ch.Publish(
		"",        // Exchange
		queueName, // Routing key (queue name)
		false,     // Mandatory
		false,     // Immediate
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish a message: %v", err)
	}

	log.Printf("Message published to queue %s: %s", queueName, message)
	return nil
}

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		log.Printf("Invalid request payload: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// SQL query to insert the product into the database
	query := `
		INSERT INTO products (user_id, product_name, product_description, product_images, product_price)
		VALUES ($1, $2, $3, $4, $5) RETURNING id
	`

	// Execute the query and handle errors
	log.Printf("Executing query: %s", query)
	log.Printf("Parameters: user_id=%d, product_name=%s, product_description=%s, product_images=%v, product_price=%f",
		product.UserID, product.ProductName, product.ProductDescription, product.ProductImages, product.ProductPrice)

	if err := config.DB.QueryRow(
		query,
		product.UserID,
		product.ProductName,
		product.ProductDescription,
		pq.Array(product.ProductImages),
		product.ProductPrice,
	).Scan(&product.ID); err != nil {
		log.Printf("Error inserting product: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to insert product into database: %v", err)})
		return
	}

	// Publish product images to RabbitMQ
	imageURLs := strings.Join(product.ProductImages, ",")
	if err := PublishToQueue("image_processing", imageURLs); err != nil {
		log.Printf("Error publishing to RabbitMQ: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to publish message to RabbitMQ"})
		return
	}

	// Success response
	c.JSON(http.StatusOK, gin.H{"id": product.ID, "message": "Product created successfully"})
}

func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	query := `
		SELECT id, user_id, product_name, product_description, product_images, compressed_product_images, product_price
		FROM products WHERE id = $1
	`
	err := config.DB.QueryRow(
		query, id,
	).Scan(
		&product.ID,
		&product.UserID,
		&product.ProductName,
		&product.ProductDescription,
		pq.Array(&product.ProductImages),
		pq.Array(&product.CompressedProductImages),
		&product.ProductPrice,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func GetProducts(c *gin.Context) {
	var products []models.Product
	filters := []string{}
	args := []interface{}{}
	argID := 1

	if userID := c.Query("user_id"); userID != "" {
		filters = append(filters, fmt.Sprintf("user_id = $%d", argID))
		args = append(args, userID)
		argID++
	}
	if productName := c.Query("product_name"); productName != "" {
		filters = append(filters, fmt.Sprintf("product_name ILIKE $%d", argID))
		args = append(args, "%"+productName+"%")
		argID++
	}
	if minPrice := c.Query("min_price"); minPrice != "" {
		minPriceValue, err := strconv.ParseFloat(minPrice, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid min_price value"})
			return
		}
		filters = append(filters, fmt.Sprintf("product_price >= $%d", argID))
		args = append(args, minPriceValue)
		argID++
	}
	if maxPrice := c.Query("max_price"); maxPrice != "" {
		maxPriceValue, err := strconv.ParseFloat(maxPrice, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid max_price value"})
			return
		}
		filters = append(filters, fmt.Sprintf("product_price <= $%d", argID))
		args = append(args, maxPriceValue)
		argID++
	}

	query := `
		SELECT id, user_id, product_name, product_description, product_images, compressed_product_images, product_price
		FROM products
	`
	if len(filters) > 0 {
		query += " WHERE " + strings.Join(filters, " AND ")
	}

	rows, err := config.DB.Query(query, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		err := rows.Scan(
			&product.ID,
			&product.UserID,
			&product.ProductName,
			&product.ProductDescription,
			pq.Array(&product.ProductImages),
			pq.Array(&product.CompressedProductImages),
			&product.ProductPrice,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse products"})
			return
		}
		products = append(products, product)
	}

	c.JSON(http.StatusOK, products)
}
