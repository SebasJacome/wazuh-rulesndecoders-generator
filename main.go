package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var a fyne.App = app.New()
var w fyne.Window = a.NewWindow("Wazuh R&D")

func main() {

	w.SetContent(MakeGui())
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}
