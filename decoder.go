package main

import (
	"errors"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type decoderInfo struct {
	log, decoderName, prematch string
	variables                  []string
}

var w2 fyne.Window

func CreateDecoderWindow(pLog string, pVariables []string) {
	w2 = a.NewWindow("Wizard Menu Decoder")

	decoderNameLabel := widget.NewLabel("Write the name of the new decoder")
	decoderNameLabel.Wrapping = fyne.TextWrapBreak
	decoderNameEntry := widget.NewEntry()
	decoderNameEntry.SetPlaceHolder("Ej. cisco-custom")
	decoderNameBox := container.NewVBox(decoderNameLabel, decoderNameEntry)

	decoderPrematchLabel := widget.NewLabel("Is there any unique identificator/word on your log?")
	decoderPrematchEntry := widget.NewEntry()
	decoderPrematchEntry.SetPlaceHolder("Ej. wazuh-manager Info:")
	decoderPrematchBox := container.NewVBox(decoderPrematchLabel, decoderPrematchEntry)

	submitButton := widget.NewButton("Submit", func() {
		if decoderNameEntry.Text != "" && decoderPrematchEntry.Text != "" {
			data := decoderInfo{
				log:         pLog,
				decoderName: decoderNameEntry.Text,
				prematch:    decoderPrematchEntry.Text,
				variables:   pVariables,
			}
			decoderXMLGenerator(data)
		}
	})

	form := container.NewVBox(decoderNameBox, decoderPrematchBox, submitButton)
	content := container.NewHBox(layout.NewSpacer(), form, layout.NewSpacer())

	w2.SetContent(content)
	w2.Resize(fyne.NewSize(800, 600))
	w2.SetFixedSize(true)
	w2.Show()
}

func decoderXMLGenerator(data decoderInfo) {
	xmlFile, err := os.Create("decoder.xml")
	if err != nil {
		dialog.ShowError(errors.New("decoder xml not created"), w2)
	} else {
		var xml string
		var regex string
		var order string
		for _, value := range data.variables {
			if regex == "" {
				regex += value + "[=:](\"\\.+\")"
			} else {
				regex += " " + value + "[=:](\"\\.+\")"
			}
			order += value + ","
			if order != "" {
				order = order[:len(order)-1]
			}
		}
		xml = "<decoder name=\"" + data.decoderName + "\">\n" +
			"\t<prematch>^" + data.prematch + "</prematch>\n" +
			"\t<regex type=\"pcre2\">" + regex + "</regex>\n" +
			"\t<order>" + order + "</order>\n" +
			"</decoder>"
		xmlFile.WriteString(xml)
		dialog.ShowCustomConfirm("Success!", "Generate Rule", "Cancel", widget.NewLabel("Do you want to generate the rule for this decoder?"), func(b bool) {
			if b {
				w2.Close()
				CreateRuleWindow(data.decoderName, data.variables)
			} else {
				w2.Close()
			}
		}, w2)
	}
}
