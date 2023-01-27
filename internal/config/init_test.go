package config

import (
	"github.com/spf13/viper"
	"testing"
)

func TestInit(t *testing.T) {
	Init("config_user.yaml")
	ans := viper.GetString("test")
	if ans != "test" {
		t.Error("failed to load viper")
	}
}
