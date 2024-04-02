package main

import (
	"errors"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
)

var a fyne.App = app.New()
var w fyne.Window = a.NewWindow("Wazuh R&D")

func main() {
	w.Resize(fyne.NewSize(800, 600))
	w.CenterOnScreen()

	// Check if ./api/conf.toml exists
	if _, err := os.Stat("./api/conf.toml"); os.IsNotExist(err) {
		w.Show()
		dialog.ShowError(errors.New("Configuration file ./api/conf.toml not found"), w)
		w.ShowAndRun()
		return
	}
	w.SetContent(MakeGui())
	w.ShowAndRun()
}
