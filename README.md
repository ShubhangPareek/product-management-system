
# Product Management System with Asynchronous Image Processing

## Overview

This project is a scalable backend system built in **Go** for managing products. It uses **PostgreSQL** for data storage, **RabbitMQ** for asynchronous message queuing, and **Redis** for caching.

---

## Features

1. **Core API Endpoints**:
   - **POST /products**: Adds a product and queues image URLs for processing.
   - **GET /products/:id**: Retrieves product details by ID.
   - **GET /products**: Fetches all products with optional filters.

2. **Asynchronous Image Processing**:
   - Processes image URLs via RabbitMQ.

3. **Caching**:
   - Caches `GET /products/:id` responses using Redis.

4. **Logging**:
   - Structured logging for all operations.

---

## Technologies Used

- **Golang**: Backend development.
- **PostgreSQL**: Relational database.
- **RabbitMQ**: Message queue for asynchronous tasks.
- **Redis**: Cache layer for improved performance.

---

## Prerequisites

- **Go**: [Install Go](https://golang.org/doc/install).
- **PostgreSQL**: [Install PostgreSQL](https://www.postgresql.org/download/).
- **RabbitMQ**: [Install RabbitMQ](https://www.rabbitmq.com/download.html).
- **Redis**: [Install Redis](https://redis.io/download).

---

## Setup Instructions

### Step 1: Clone the Repository
```bash
git clone https://github.com/your-username/product-management-system.git
cd product-management-system
```

### Step 2: Set Up Environment Variables
1. Copy `.env.example` to `.env`:
   ```bash
   cp .env.example .env
   ```
2. Update the `.env` file with your PostgreSQL, Redis, and RabbitMQ credentials.

---

### Step 3: Set Up PostgreSQL

1. Create the database:
   ```sql
   CREATE DATABASE product_management;
   ```
2. Execute the schema:
   ```sql
   CREATE TABLE users (
       id SERIAL PRIMARY KEY,
       name VARCHAR(255) NOT NULL
   );

   CREATE TABLE products (
       id SERIAL PRIMARY KEY,
       user_id INT REFERENCES users(id) ON DELETE CASCADE,
       product_name VARCHAR(255) NOT NULL,
       product_description TEXT,
       product_images TEXT[],
       compressed_product_images TEXT[],
       product_price NUMERIC(10, 2) NOT NULL
   );
   ```

3. Insert a test user:
   ```sql
   INSERT INTO users (id, name) VALUES (1, 'Test User');
   ```

---

### Step 4: Install Dependencies

```bash
go mod tidy
```

---

### Step 5: Run the Application

1. **Start the Server**:
   ```bash
   go run cmd/main.go
   ```
   Access the API at `http://localhost:8080`.

2. **Start the RabbitMQ Consumer**:
   ```bash
   go run services/image_processor.go
   ```

---

## API Usage

### POST /products
- **Payload**:
  ```json
  {
      "user_id": 1,
      "product_name": "Sample Product",
      "product_description": "This is a test product",
      "product_images": ["http://example.com/image1.jpg", "http://example.com/image2.jpg"],
      "product_price": 123.45
  }
  ```
- **Response**:
  ```json
  {
      "id": 1,
      "message": "Product created successfully"
  }
  ```

---

### GET /products/:id
- **Response**:
  ```json
  {
      "id": 1,
      "user_id": 1,
      "product_name": "Sample Product",
      "product_description": "This is a test product",
      "product_images": ["http://example.com/image1.jpg"],
      "compressed_product_images": ["http://example.com/compressed_image1.jpg"],
      "product_price": 123.45
  }
  ```

---

### GET /products
- **Query Parameters**:
  - `user_id=1`
  - `product_name=Sample`
  - `min_price=100`
  - `max_price=200`
- **Response**:
  ```json
  [
      {
          "id": 1,
          "user_id": 1,
          "product_name": "Sample Product",
          "product_description": "This is a test product",
          "product_images": ["http://example.com/image1.jpg"],
          "compressed_product_images": ["http://example.com/compressed_image1.jpg"],
          "product_price": 123.45
      }
  ]
  ```

---

## Testing

### Unit Tests
```bash
go test ./... -v
```

### Code Coverage
```bash
go test ./... -cover
```

---

## Author

- **Your Name**
- GitHub: [your-username](https://github.com/your-username)
- LinkedIn: [Your LinkedIn](https://linkedin.com/in/your-profile)

---

## License
This project is licensed under the MIT License.

![Uploading image.png…]()
