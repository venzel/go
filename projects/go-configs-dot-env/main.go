package main

import (
	"dotenv/configs"
	"fmt"
	"reflect"
)

func main() {
	cfg, _ := configs.InitConfig(".")

	fmt.Println(cfg.APIName)
	fmt.Println(cfg.APIPort)
	fmt.Println(cfg.DBHost)
	fmt.Println(cfg.DBPort)
	fmt.Println(cfg.DBUser)
	fmt.Println(cfg.DBPassword)
	fmt.Println(cfg.DBName)

	typeApiPort := reflect.TypeOf(cfg.APIPort)

	fmt.Println(typeApiPort)
}
