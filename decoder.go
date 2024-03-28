package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func CreateDecoderWindow() {
	w2 := a.NewWindow("Wizard Menu Decoder")

	decoderNameLabel := widget.NewLabel("Write the name of the new decoder")
	decoderNameLabel.Wrapping = fyne.TextWrapBreak
	decoderNameEntry := widget.NewEntry()
	decoderNameEntry.SetPlaceHolder("Ej. cisco-custom")
	decoderNameBox := container.NewVBox(decoderNameLabel, decoderNameEntry)

	decoderPrematchLabel := widget.NewLabel("Is there any unique identificator/word on your log?")
	decoderPrematchEntry := widget.NewEntry()
	decoderPrematchEntry.SetPlaceHolder("Ej. wazuh-manager Info:")
	decoderPrematchBox := container.NewVBox(decoderPrematchLabel, decoderPrematchEntry)

	content := container.NewVBox(decoderNameBox, decoderPrematchBox)

	w2.SetContent(content)
	w2.Resize(fyne.NewSize(800, 600))
	w2.SetFixedSize(true)
	w2.Show()
}
