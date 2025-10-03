package main

import "go-musthave-shortener-tpl/internal/app"

func main() {

	application := app.New()

	if err := application.Run(); err != nil {
		panic(err)
	}

}
