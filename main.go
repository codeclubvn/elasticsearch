package main

import (
	"elasticsearch/conf"
	"elasticsearch/route"
)

func main() {
	conf.SetEnv()
	app := route.NewService()
	if err := app.Start(); err != nil {
		panic(err)
	}
}
