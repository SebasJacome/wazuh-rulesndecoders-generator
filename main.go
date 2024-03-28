package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var a fyne.App = app.New()

func main() {
	w := a.NewWindow("Wazuh R&D")

	w.SetContent(MakeGui())
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}
