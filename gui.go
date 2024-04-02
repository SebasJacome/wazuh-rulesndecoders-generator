//go:generate fyne bundle -o bundled.go assets
package main

import (
	"errors"
	"fmt"
	"go_gui/utils"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var red color.NRGBA = color.NRGBA{R: 180, G: 0, B: 0, A: 255}
var logEntries = widget.NewEntry()
var variablesEntries = widget.NewEntry()

func MakeGui() fyne.CanvasObject {
	return container.NewBorder(makeToolBar(), nil, nil, nil, createMainContent())
}
func makeToolBar() fyne.CanvasObject {
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.SearchIcon(), SearchWindow),
		widget.NewToolbarAction(theme.ZoomInIcon(), SelectIDorDecoderWindow),
		widget.NewToolbarSeparator(),
		widget.NewToolbarSpacer(),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.DeleteIcon(), func() {
			logEntries.SetText("")
			variablesEntries.SetText("")
		}),
	)
	logo := canvas.NewImageFromResource(resourceLogosUPPng)
	logo.FillMode = canvas.ImageFillContain
	return container.NewStack(toolbar, logo)
}

func createMainContent() fyne.CanvasObject {
	errorLabel := canvas.NewText("Cannot Be Empty", red)
	errorLabel2 := canvas.NewText("Cannot Be Empty", red)
	errorLabel.Hide()
	errorLabel2.Hide()

	logLabel := widget.NewLabel("Enter the log from where you want to extract the information")
	logEntries.SetPlaceHolder("Info - New agent connected: { \"name\": \"ExampleAgent\", \"ip\": \"192.168.1.100\", \"id\": \"001\"...")
	logEntries.Validator = func(input string) error {
		if len(logEntries.Text) == 0 {
			return errors.New("Cannot be empty")
		} else {
		}
		return nil
	}
	logEntries.OnChanged = func(str string) {
		if logEntries.Text == "" {
			errorLabel.Show()
			errorLabel.TextStyle.Italic = true
		} else {
			errorLabel.Hide()
		}
	}

	variablesLabel := widget.NewLabel("Type the variables that you want to read (multiple variables have to be separated by commas)")
	variablesEntries.SetPlaceHolder("name,ip,id...")
	variablesEntries.Validator = func(input string) error {
		if len(variablesEntries.Text) == 0 {
			return errors.New("Cannot be empty")
		} else {
		}
		return nil
	}

	variablesEntries.OnChanged = func(str string) {
		if variablesEntries.Text == "" {
			errorLabel2.Show()
			errorLabel2.TextStyle.Italic = true
		} else {
			errorLabel2.Hide()
		}
	}

	data := struct {
		log, variables string
	}{
		log:       "",
		variables: "",
	}

	button := widget.NewButton("Submit", func() {
		if variablesEntries.Text != "" && logEntries.Text != "" {
			data.log = logEntries.Text
			data.variables = variablesEntries.Text
			processData(data)
		}
	})

	inputVBox := container.NewVBox(layout.NewSpacer(), logLabel, logEntries, errorLabel, variablesLabel, variablesEntries, errorLabel2, layout.NewSpacer(), button, layout.NewSpacer(), layout.NewSpacer())
	return container.NewHBox(layout.NewSpacer(), inputVBox, layout.NewSpacer())

}

func processData(data struct {
	log, variables string
}) {
	fmt.Println("Processing the variables...")
	processedVars := utils.SplitString(data.variables)
	result := utils.CompareLogAndVars(data.log, processedVars)
	if result == "" {
		fmt.Println("All the variables were found inside the log...")
		fmt.Println("End of validation")
		CreateDecoderWindow(data.log, processedVars, w)
	} else {
		fmt.Println("There is one or more variables that were not found in the log")
		fmt.Println("The variables must be contained by the log")
		dialog.ShowInformation("Invalid input", "Variables "+result+" were not found in the log", w)
	}
}
