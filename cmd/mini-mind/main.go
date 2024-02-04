package main

import (
	"github.com/adriano-henrique/mini-mind-be/api"
	"github.com/adriano-henrique/mini-mind-be/internal/database"
)

func main() {
	database.DatabaseConnect()
	api.BuildRoutes()
}
