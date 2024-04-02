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

var green color.NRGBA = color.NRGBA{R: 0, G: 180, B: 0, A: 255}

func SelectIDorDecoderWindow() {
	var opcion string
	selector := widget.NewSelect([]string{"Rule", "Decoder"}, func(opt string) {
		opcion = opt
	})
	selectorItem := widget.NewFormItem("Selecciona una opcion: ", selector)

	dialog.ShowForm("Selector de Opción", "Continue", "Cancel", []*widget.FormItem{selectorItem}, func(bool) {
		if opcion == "Rule" {
			response := api.SearchForAllIDs()
			RuleSearchListWindow(response)
		}
	}, w)
}

func RuleSearchListWindow(list []api.MatchingRule) {
	w5 := a.NewWindow("Rules")
	w5.Resize(fyne.NewSize(800, 600))
	w5.SetFixedSize(true)

	listView := widget.NewList(func() int {
		return len(list)
	}, func() fyne.CanvasObject {
		return widget.NewLabel("")
	}, func(id widget.ListItemID, object fyne.CanvasObject) {
		object.(*widget.Label).SetText(strconv.Itoa(list[id].ID))
	})

	contentText := widget.NewLabel("Please select a rule ID")
	contentText.Wrapping = fyne.TextWrapWord

	IDLabel := widget.NewLabel("")
	IDLabel.Wrapping = fyne.TextWrapWord
	LevelLabel := widget.NewLabel("")
	LevelLabel.Wrapping = fyne.TextWrapWord
	DescriptionLabel := widget.NewLabel("")
	DescriptionLabel.Wrapping = fyne.TextWrapWord
	FileNameLabel := widget.NewLabel("")
	FileNameLabel.Wrapping = fyne.TextWrapWord
	DirNameLabel := widget.NewLabel("")
	DirNameLabel.Wrapping = fyne.TextWrapWord
	var StatusLabel widget.Label
	StatusLabel.Alignment = fyne.TextAlignCenter
	StatusLabel.Wrapping = fyne.TextWrapWord
	listView.OnSelected = func(id widget.ListItemID) {
		strID := strconv.Itoa(list[id].ID)
		strLevel := strconv.Itoa(list[id].Level)

		contentText.SetText("")
		IDLabel.SetText("ID: " + strID)
		LevelLabel.SetText("Level: " + strLevel)
		DescriptionLabel.SetText("Description: " + list[id].Description)
		FileNameLabel.SetText("File name: " + list[id].FileName)
		DirNameLabel.SetText("Relative Directory: " + list[id].DirName)
		var StatusText string
		if list[id].Status == "enabled" {
			StatusText = "  Status: " + list[id].Status
		} else {
			StatusText = "  Status: " + list[id].Status
		}
		StatusLabel.SetText(StatusText)
	}

	contentContainer := container.NewVBox(IDLabel, LevelLabel, DescriptionLabel, FileNameLabel, DirNameLabel)
	split := container.NewHSplit(listView, container.NewStack(contentText, contentContainer))
	split.Offset = 0.15
	w5.SetContent(split)
	w5.CenterOnScreen()
	w5.Show()
}

func SearchWindow() {
	id := widget.NewEntry()
	idName := widget.NewFormItem("ID", id)
	dialog.ShowForm("Search by Rule ID", "Yes", "Cancel", []*widget.FormItem{idName}, func(b bool) {
		if b {
			progressBar := dialog.NewCustomWithoutButtons("Searching...", widget.NewProgressBarInfinite(), w)
			progressBar.Resize(fyne.NewSize(300, 100))
			progressBar.Show()
			if id.Text == "" || id.Text == "0" {
				progressBar.Hide()
				dialog.ShowError(errors.New("Invalid rule ID, try again"), w)
				return
			}
			idInt, err := strconv.Atoi(id.Text)
			if err != nil {
				panic(err)
			}
			response := api.SearchRequestedID(idInt)
			if response.ID == -1 && response.Description == "null" && response.Level == -1 {
				progressBar.Hide()
				dialog.ShowError(errors.New("That rule doesn't exist. Try an existing one"), w)
				return
			}
			createSearchContent(response)
			progressBar.Hide()
		}
	}, w)
}

func createSearchContent(rule api.MatchingRule) {
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
