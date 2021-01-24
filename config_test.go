package config

import (
	"testing"
)

func TestAppName(t *testing.T) {
	t.Log(AppName())
	dir, filename := AppName()
	t.Log(dir)
	t.Log(filename)
	loadConfig(nil)
}
