package main

import configs "github.com/devlucascardoso/api-rest-golang/config"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}