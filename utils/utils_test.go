package utils

import (
	"testing"
	"time"
)

func TestInitCache(t *testing.T) {
	InitCache("localhost:6379")
	if redisClient == nil {
		t.Error("Failed to initialize Redis client")
	}
}

func TestLogger(t *testing.T) {
	InitLogger()
	if Logger == nil {
		t.Error("Logger not initialized")
	}
	Logger.Info("Logger test message")
}

func TestCache(t *testing.T) {
	InitCache("localhost:6379")

	err := SetCache("test_key", "test_value")
	if err != nil {
		t.Fatalf("Failed to set cache: %v", err)
	}

	value, err := GetCache("test_key")
	if err != nil {
		t.Fatalf("Failed to get cache: %v", err)
	}

	if value != "test_value" {
		t.Errorf("Expected 'test_value', got '%s'", value)
	}
}
func TestCacheEdgeCases(t *testing.T) {
	InitCache("localhost:6379")

	// Test empty key
	_, err := GetCache("")
	if err == nil {
		t.Errorf("Expected error for empty key, got none")
	} else {
		t.Logf("Received expected error for empty key: %v", err)
	}

	// Test invalid key (key that doesn't exist)
	_, err = GetCache("nonexistent_key")
	if err == nil {
		t.Errorf("Expected error for nonexistent key, got none")
	} else {
		t.Logf("Received expected error for nonexistent key: %v", err)
	}
}
func TestCacheTTL(t *testing.T) {
	err := SetCache("key_with_ttl", "value_with_ttl")
	if err != nil {
		t.Fatalf("Failed to set cache with TTL: %v", err)
	}

	val, err := GetCache("key_with_ttl")
	if err != nil || val != "value_with_ttl" {
		t.Errorf("Expected 'value_with_ttl', got '%s'", val)
	}
}
func TestCacheInvalidKey(t *testing.T) {
	InitCache("localhost:6379")

	err := SetCache("", "some_value")
	if err == nil {
		t.Errorf("Expected error for empty cache key, got none")
	} else {
		t.Logf("Received expected error for empty cache key: %v", err)
	}
}
func TestCacheExpiration(t *testing.T) {
	InitCache("localhost:6379")

	// Set a cache key with a value
	SetCache("temp_key", "temp_value")

	// Ensure the key exists immediately
	value, err := GetCache("temp_key")
	if err != nil || value != "temp_value" {
		t.Fatalf("Expected temp_value, got %v (error: %v)", value, err)
	}

	// Wait for the cache to expire
	time.Sleep(11 * time.Second)

	// Check that the key has expired
	_, err = GetCache("temp_key")
	if err == nil {
		t.Errorf("Expected cache miss (error: %v)", err)
	} else {
		t.Logf("Cache expired successfully: %v", err)
	}
}
