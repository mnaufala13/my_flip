package app

import (
	"embed"
	"flip/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_newDB(t *testing.T) {
	cfg := config.NewConfig()
	db, err := newDB(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSsl)
	defer func() {
		if err == nil {
			db.Close()
		}
	}()
	assert.NoError(t, err)
}

func TestNewApp(t *testing.T) {
	cfg := config.NewConfig()
	app := NewApp(*cfg, embed.FS{})
	defer func() {
		app.DB.Close()
	}()
}
