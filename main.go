package main

import (
	"embed"
	"flip/cmd"
)

//go:embed template/index.html template/history.html
var templates embed.FS

func main() {
	cmd.Execute(templates)
}
