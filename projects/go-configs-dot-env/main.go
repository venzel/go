package main

import (
	"dotenv/configs"
	"fmt"
	"reflect"
)

func main() {
	cfg, err := configs.InitConfig(".")

	if err != nil {
		panic(err)
	}

	fmt.Println(cfg.APIName)
	fmt.Println(cfg.APIEnv)

	typeApiPort := reflect.TypeOf(cfg.APIPort)

	fmt.Println(typeApiPort)
}
