package config

import (
	"github.com/spf13/viper"
	"testing"
)

func TestAppName(t *testing.T) {
	t.Log(AppName())
	dir, filename := AppName()
	t.Log(dir)
	t.Log(filename)
	loadConfig(nil)
	t.Log(viper.Get("testapp"))
}
