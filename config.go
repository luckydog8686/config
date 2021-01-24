package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func AppName()(string,string){
	appName,err := os.Executable()
	if err != nil {
		panic(err)
	}
	absPath,err:=filepath.Abs(appName)
	if err != nil {
		panic(err)
	}
	dir := filepath.Dir(absPath)
	exeName := filepath.Base(absPath)
	return dir,exeName
}


func SetConfig(appname string){
	viper.AddConfigPath(".")
	viper.AddConfigPath(fmt.Sprintf("/etc/%s/",appname))
	viper.AddConfigPath("$HOME/.appname")
	viper.AddConfigPath("../configs")
	viper.AddConfigPath("./configs")
	viper.SetConfigName(appname)
	viper.SetConfigType("json")
}

func loadConfig(cf map[string]string){
	_,app := AppName()
	SetConfig(app)
	cfPath,ok := cf["config_path"]
	if ok{
		viper.AddConfigPath(cfPath)
	}
	cfType,ok :=cf["config_type"]
	if ok {
		viper.SetConfigType(cfType)
	}
	cfName,ok := cf["config_name"]
	if ok{
		viper.SetConfigName(cfName)
	}
}

func Get(key string) interface{}  {
	return nil
}

func init()  {
	
}

