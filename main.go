package main

import (
	"github.com/dogg5432/cs_ocpp2.0/config"
	"github.com/dogg5432/cs_ocpp2.0/serve"
)

func main(){
	if err := config.Load(); err != nil {
		panic(err)
	}
	// if err := database.Connect(); err != nil {
	// 	panic(err)
	// }
	serve.Run()
}