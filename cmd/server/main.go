package main

import configs "github.com/devlucascardoso/api-rest-golang/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}