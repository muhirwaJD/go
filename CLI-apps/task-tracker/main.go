package main

import (
	"music/cmd"
	"music/database"
)

func main() {
	database.Init(".env")
	cmd.Execute()
}
