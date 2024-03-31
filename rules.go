package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"go_gui/api"
	"go_gui/utils"
)

var w3 fyne.Window

type ruleInfo struct {
	id, level, description, decoderName string
	variables                           []string
}

func CreateRuleWindow(pDecoderName string, pVariables []string) {
	w3 = a.NewWindow("Wizard Menu Rule")
	fmt.Println("I've created the window")
	ruleIDLabel := widget.NewLabel("Write the ID of your new rule")
	ruleIDEntry := widget.NewEntry()
	ruleIDEntry.SetPlaceHolder("Ej. 202232")
	ruleIDBox := container.NewVBox(ruleIDLabel, ruleIDEntry)

	ruleLevelLabel := widget.NewLabel("Write the Level of your new rule")
	ruleLevelEntry := widget.NewEntry()
	ruleLevelEntry.SetPlaceHolder("Ej. 10 (Rule level can go from 1 to 15)")
	ruleLevelBox := container.NewVBox(ruleLevelLabel, ruleLevelEntry)

	ruleDescriptionLabel := widget.NewLabel("Write description of your new rule")
	ruleDescriptionEntry := widget.NewEntry()
	ruleDescriptionEntry.SetPlaceHolder("Fortinet: Login Failed")
	ruleDescriptionBox := container.NewVBox(ruleDescriptionLabel, ruleDescriptionEntry)

	submitButton := widget.NewButton("Submit", func() {
		if ruleIDEntry.Text != "" && ruleDescriptionEntry.Text != "" && ruleLevelEntry.Text != "" {
			ruleLevel, err := strconv.Atoi(ruleLevelEntry.Text)
			if err != nil {
				dialog.ShowError(errors.New("Invalid rule level. Please enter a number."), w3)
				return
			}

			if ruleLevel > 16 {
				dialog.ShowError(errors.New("Rule level cannot be greater than 16."), w3)
				return
			}

			var ruleIDs map[string]bool = api.RequestRuleIDs(w3)
			ruleID, err := strconv.Atoi(ruleIDEntry.Text)
			if ruleID < 100000 || ruleID > 120000 {
				dialog.ShowError(errors.New("The rule must be between the range 100000-120000"), w3)
			}

			if utils.CompareExistingIDs(ruleIDs, ruleIDEntry.Text) {
				dialog.ShowError(errors.New("That rule already exists, pick another ID"), w3)
				return
			}

			data := ruleInfo{
				id:          ruleIDEntry.Text,
				level:       ruleLevelEntry.Text,
				description: ruleDescriptionEntry.Text,
				decoderName: pDecoderName,
				variables:   pVariables,
			}
			ruleXMLGenerator(data)
		}
	})

	fmt.Println("I've created the content")

	form := container.NewVBox(ruleIDBox, ruleLevelBox, ruleDescriptionBox, submitButton)
	content := container.NewHBox(layout.NewSpacer(), form, layout.NewSpacer())
	w3.SetContent(content)
	w3.Resize(fyne.NewSize(800, 600))
	w3.SetFixedSize(true)
	w3.Show()
	fmt.Println("I've shown the window")
}

func ruleXMLGenerator(data ruleInfo) {
	xmlFile, err := os.Create("rules.xml")
	if err != nil {
		dialog.ShowError(errors.New("rule xml not created"), w3)
	} else {
		var xml string = "<rule id=\"" + data.id + " level=\"" + data.level + ">\n" +
			"\t<decoded_as>" + data.decoderName + "</decoded_as>\n" +
			"\t<description>" + data.description + "</description>\n" +
			"</rule>\n" +
			"</group>"
		xmlFile.WriteString(xml)
		dialog.ShowConfirm("Successful creation!", "Do you want to upload both files to your Wazuh Server?", func(bool) {
			// Implementación de función para subir archivos a wazuh usando su API
			w3.Close()
		}, w3)
	}
}
