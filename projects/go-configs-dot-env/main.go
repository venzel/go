package main

import (
	"dotenv/configs"
	"fmt"
	"reflect"
)

func main() {
	cfg, _ := configs.InitConfig(".")

	fmt.Println(cfg.Validate())

	fmt.Println(cfg.APIName)
	fmt.Println(cfg.APIEnv)

	typeApiPort := reflect.TypeOf(cfg.APIPort)

	fmt.Println(typeApiPort)
}
