package main

import (
	"errors"
	"go_gui/api"
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func SearchWindow() {
	id := widget.NewEntry()
	idName := widget.NewFormItem("ID", id)
	dialog.ShowForm("Search by Rule ID", "Yes", "Cancel", []*widget.FormItem{idName}, func(b bool) {
		if b {
			if id.Text == "" || id.Text == "0" {
				dialog.ShowError(errors.New("Invalid rule ID, try again"), w)
				return
			}
			idInt, err := strconv.Atoi(id.Text)
			if err != nil {
				panic(err)
			}
			response := api.SearchRequestedID(idInt)
			if response.ID == -1 && response.Description == "null" && response.Level == -1 {
				dialog.ShowError(errors.New("That rule doesn't exist. Try an existing one"), w)
				return
			}
			createSearchContent(response)
		}
	}, w)
}

func createSearchContent(rule api.MatchingRule) {
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	strID := strconv.Itoa(rule.ID)
	strLevel := strconv.Itoa(rule.Level)
	w4 := a.NewWindow("Information for ID: " + strID)
	IDLabel := widget.NewLabel("ID: " + strID)
	LevelLabel := widget.NewLabel("Level: " + strLevel)
	DescriptionLabel := widget.NewLabel("Description: " + rule.Description)
	FileNameLabel := widget.NewLabel("File name: " + rule.FileName)
	DirNameLabel := widget.NewLabel("Relative Directory: " + rule.DirName)
	var StatusLabel *canvas.Text
	if rule.Status == "enabled" {
		StatusLabel = canvas.NewText("  Status: "+rule.Status, green)
	} else {
		StatusLabel = canvas.NewText("  Status: "+rule.Status, red)
	}
	StatusLabel.Alignment = fyne.TextAlignCenter

	CancelButton := widget.NewButtonWithIcon("Cancel", theme.ContentClearIcon(), func() {
		w4.Close()
		w.RequestFocus()
	})

	SearchButton := widget.NewButtonWithIcon("Search another rule", theme.SearchIcon(), func() {
		w4.Close()
		SearchWindow()
	})

	contentContainer := container.NewVBox(IDLabel, LevelLabel, DescriptionLabel, FileNameLabel, DirNameLabel, StatusLabel)
	buttonsContainer := container.NewHBox(CancelButton, layout.NewSpacer(), SearchButton)

	containers := container.NewVBox(contentContainer, layout.NewSpacer(), buttonsContainer)

	content := container.New(layout.NewCenterLayout(), containers)

	w4.Resize(fyne.NewSize(600, 400))
	w4.SetContent(content)
	w4.CenterOnScreen()
	w4.SetFixedSize(true)
	w4.Show()
}
