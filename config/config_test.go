package config

import (
	"testing"
)

func TestConnectDB(t *testing.T) {
	ConnectDB()
	if DB == nil {
		t.Error("Database connection failed")
	}
}
