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

var w3 fyne.Window

type ruleInfo struct {
	id, level, description, decoderName string
	variables                           []string
}

func CreateRuleWindow(pDecoderName string, pVariables []string) {
	w3 = a.NewWindow("Wizard Menu Rule")

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
		if ruleIDEntry.Text != "" && ruleLevelEntry.Text != "" && ruleDescriptionEntry.Text != "" {
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

	form := container.NewVBox(ruleIDBox, ruleLevelBox, ruleDescriptionBox, submitButton)
	content := container.NewHBox(layout.NewSpacer(), form, layout.NewSpacer())
	w3.SetContent(content)
	w3.Resize(fyne.NewSize(800, 600))
	w3.SetFixedSize(true)
	w3.Show()
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
		dialog.ShowInformation("Success!", "Rule xml file created succesfully", w3)
	}
}

// <rule id="222000" level="3">
//   <decoded_as>fortigate-custom</decoded_as>
//   <description>Fortigate message: $(syscheck).</description>
// </rule>
