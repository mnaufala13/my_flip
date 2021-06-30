package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewConfig(t *testing.T) {
	os.Setenv("DBHost", "DBHost")
	os.Setenv("DBPort", "DBPort")
	os.Setenv("DBUser", "DBUser")
	os.Setenv("DBPassword", "DBPassword")
	os.Setenv("DBName", "DBName")
	os.Setenv("FlipHost", "FlipHost")
	os.Setenv("FlipSecret", "FlipSecret")
	cfg := NewConfig()
	assert.EqualValues(t, "DBHost", cfg.DBHost)
	assert.EqualValues(t, "DBPort", cfg.DBPort)
	assert.EqualValues(t, "DBUser", cfg.DBUser)
	assert.EqualValues(t, "DBPassword", cfg.DBPassword)
	assert.EqualValues(t, "DBName", cfg.DBName)
	assert.EqualValues(t, "FlipHost", cfg.FlipHost)
	assert.EqualValues(t, "FlipSecret", cfg.FlipSecret)
}
