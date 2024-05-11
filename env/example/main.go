package main

import "github.com/gohryt/asphyxia/env"

type Configuration struct {
	Name string `name:"NAME"`
}

func main() {
	configuration, err := env.Parse[Configuration]()
	if err != nil {
		panic(err)
	}

	_ = configuration
}
