package main

import (
	"agenda/storage"
	"agenda/ui"
)

func main() {
	storage.LoadCompromissos("compromissos.json")
	ui.RunUI()
}
