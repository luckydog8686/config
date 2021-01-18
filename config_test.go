package config

import (
	"github.com/luckydog8686/logs"
	"path/filepath"
	"testing"
)

func TestAppName(t *testing.T) {
	t.Log(AppName())
	appName:= AppName()
	logs.Info(filepath.Dir(appName))
	logs.Info(filepath.IsAbs(appName))
	logs.Info(filepath.Base(appName))
	logs.Info(filepath.Abs(appName))
	logs.Info(filepath.Ext(appName))
	logs.Info(filepath.Join(appName,"kdfaksdf"))
}
