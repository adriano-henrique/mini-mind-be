package main

import (
	"github.com/adriano-henrique/mini-mind-be/internal/routes"
)

func main() {
	r := routes.BuildRoutes()
	r.Run()
}
