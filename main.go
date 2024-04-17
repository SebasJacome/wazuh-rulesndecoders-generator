package main

import (
	"errors"
	"go_gui/api"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var a fyne.App = app.New()
var w fyne.Window = a.NewWindow("Wazuh R&D")

func main() {
	w.Resize(fyne.NewSize(800, 600))
	w.CenterOnScreen()

	if resourceConfToml.StaticContent == nil {
		panic(errors.New("Error: No data from configuration file"))
	}
	api.ReadConfFile(resourceConfToml.StaticContent)
	w.SetContent(MakeGui())
	w.ShowAndRun()
}
