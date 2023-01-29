package db

import (
	"testing"
	"tiktok/internal/config"
)

func TestInit(t *testing.T) {
	config.Init("config_user.yaml")
	Init()
}
