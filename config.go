package config

import (
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
	filePath.Dir(absPath)
}


func SetConfigPath(){
	
}

func ReloadConfig(){

}

func Get(key string) interface{}  {
	return nil
}

func init()  {
	
}

