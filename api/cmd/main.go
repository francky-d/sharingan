package main

import (
	_ "github.com/joho/godotenv/autoload"
	"gitlab.jems-group.com/fdjacoto/sharingan/backend/internal/core"
)

func main() {
	core.App{}.Start()
}
