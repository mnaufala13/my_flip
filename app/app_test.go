package app

import (
	"flip/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_newDB(t *testing.T) {
	cfg := config.NewConfig()
	db, err := newDB(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	defer func() {
		if err == nil {
			db.Close()
		}
	}()
	assert.NoError(t, err)
}

func TestNewApp(t *testing.T) {
	cfg := config.NewConfig()
	app := NewApp(*cfg)
	defer func() {
		app.DB.Close()
	}()
}
