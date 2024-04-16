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

func SelectIDorDecoderListWindow() {
	var opcion string
	selector := widget.NewSelect([]string{"Rule", "Decoder"}, func(opt string) {
		opcion = opt
	})
	selectorItem := widget.NewFormItem("Selecciona una opcion: ", selector)

	dialog.ShowForm("Selector de Opción", "Continue", "Cancel", []*widget.FormItem{selectorItem}, func(b bool) {
		if b {
			if opcion == "Rule" {
				response := api.SearchForAllIDs()
				RuleSearchListWindow(response)
			} else if opcion == "Decoder" {
				response := api.SearchForAllDecoders()
				DecoderSearchListWindow(response)
			} else {
				dialog.ShowError(errors.New("Selecciona una opción válida"), w)
			}
		} else {
			return
		}
	}, w)
}

func SelectIDorDecoderSpecificWindow() {
	var opcion string
	selector := widget.NewSelect([]string{"Rule", "Decoder"}, func(opt string) {
		opcion = opt
	})
	selectorItem := widget.NewFormItem("Selecciona una opcion: ", selector)

	dialog.ShowForm("Selector de Opción", "Continue", "Cancel", []*widget.FormItem{selectorItem}, func(b bool) {
		if b {
			if opcion == "Rule" {
				SearchRuleWindow()
			} else if opcion == "Decoder" {
				SearchDecoderWindow()
			} else {
				dialog.ShowError(errors.New("Selecciona una opción válida"), w)
			}
		} else {
			return
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

func DecoderSearchListWindow(list []api.MatchingDecoder) {
	w6 := a.NewWindow("Decoder")
	w6.Resize(fyne.NewSize(800, 600))
	w6.SetFixedSize(true)

	listView := widget.NewList(func() int {
		return len(list)
	}, func() fyne.CanvasObject {
		return widget.NewLabel("")
	}, func(id widget.ListItemID, object fyne.CanvasObject) {
		object.(*widget.Label).SetText(list[id].Name)
	})

	contentText := widget.NewLabel("Please select a rule ID")
	contentText.Wrapping = fyne.TextWrapWord

	NameLabel := widget.NewLabel("")
	NameLabel.Wrapping = fyne.TextWrapWord
	FileNameLabel := widget.NewLabel("")
	FileNameLabel.Wrapping = fyne.TextWrapWord
	DirNameLabel := widget.NewLabel("")
	DirNameLabel.Wrapping = fyne.TextWrapWord
	ParentNameLabel := widget.NewLabel("")
	ParentNameLabel.Wrapping = fyne.TextWrapWord
	RegexNameLabel := widget.NewLabel("")
	RegexNameLabel.Wrapping = fyne.TextWrapWord
	PrematchNameLabel := widget.NewLabel("")
	PrematchNameLabel.Wrapping = fyne.TextWrapWord

	var StatusLabel widget.Label
	StatusLabel.Alignment = fyne.TextAlignCenter
	StatusLabel.Wrapping = fyne.TextWrapWord
	listView.OnSelected = func(id widget.ListItemID) {

		contentText.SetText("")
		NameLabel.SetText("Name: " + list[id].Name)
		FileNameLabel.SetText("File name: " + list[id].FileName)
		DirNameLabel.SetText("Relative Directory: " + list[id].RelativeDirName)
		ParentNameLabel.SetText("Parent: " + list[id].Parent)
		RegexNameLabel.SetText("Regex: " + list[id].Regex)
		PrematchNameLabel.SetText("Prematch: " + list[id].Prematch)
		var StatusText string
		if list[id].Status == "enabled" {
			StatusText = "  Status: " + list[id].Status
		} else {
			StatusText = "  Status: " + list[id].Status
		}
		StatusLabel.SetText(StatusText)
	}

	contentContainer := container.NewVBox(NameLabel, FileNameLabel, DirNameLabel, ParentNameLabel, RegexNameLabel, PrematchNameLabel)
	split := container.NewHSplit(listView, container.NewStack(contentText, contentContainer))
	split.Offset = 0.15
	w6.SetContent(split)
	w6.CenterOnScreen()
	w6.Show()
}

func SearchRuleWindow() {
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
			createRuleSearchContent(response)
			progressBar.Hide()
		}
	}, w)
}

func SearchDecoderWindow() {
	name := widget.NewEntry()
	decoderName := widget.NewFormItem("Decoder Name", name)
	dialog.ShowForm("Search by Decoder Name", "Yes", "Cancel", []*widget.FormItem{decoderName}, func(b bool) {
		if b {
			progressBar := dialog.NewCustomWithoutButtons("Searching...", widget.NewProgressBarInfinite(), w)
			progressBar.Resize(fyne.NewSize(300, 100))
			progressBar.Show()
			if name.Text == "" || name.Text == "0" {
				progressBar.Hide()
				dialog.ShowError(errors.New("Invalid decoder name, try again"), w)
				return
			}
			response := api.SearchRequestedName(name.Text)
			if response.Name == "" && response.FileName == "" && response.Regex == "" {
				progressBar.Hide()
				dialog.ShowError(errors.New("That rule doesn't exist. Try an existing one"), w)
				return
			}
			createDecoderSearchContent(response)
			progressBar.Hide()
		}
	}, w)
}
func createRuleSearchContent(rule api.MatchingRule) {
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
		SearchRuleWindow()
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

func createDecoderSearchContent(decoder api.MatchingDecoder) {
	w7 := a.NewWindow("Information for decoder: " + decoder.Name)

	NameLabel := widget.NewLabel("Name: " + decoder.Name)
	FileNameLabel := widget.NewLabel("File name: " + decoder.FileName)
	DirNameLabel := widget.NewLabel("Relative Directory: " + decoder.RelativeDirName)
	ParentNameLabel := widget.NewLabel("Parent: " + decoder.Parent)
	RegexNameLabel := widget.NewLabel("Regex: " + decoder.Regex)
	PrematchNameLabel := widget.NewLabel("Prematch: " + decoder.Prematch)
	var StatusLabel *canvas.Text
	if decoder.Status == "enabled" {
		StatusLabel = canvas.NewText("  Status: "+decoder.Status, green)
	} else {
		StatusLabel = canvas.NewText("  Status: "+decoder.Status, red)
	}
	StatusLabel.Alignment = fyne.TextAlignCenter

	CancelButton := widget.NewButtonWithIcon("Cancel", theme.ContentClearIcon(), func() {
		w7.Close()
		w.RequestFocus()
	})

	SearchButton := widget.NewButtonWithIcon("Search another decoder", theme.SearchIcon(), func() {
		w7.Close()
		SearchDecoderWindow()
	})

	contentContainer := container.NewVBox(NameLabel, ParentNameLabel, RegexNameLabel, PrematchNameLabel, FileNameLabel, DirNameLabel, StatusLabel)
	buttonsContainer := container.NewHBox(CancelButton, layout.NewSpacer(), SearchButton)

	containers := container.NewVBox(contentContainer, layout.NewSpacer(), buttonsContainer)

	content := container.New(layout.NewCenterLayout(), containers)

	w7.Resize(fyne.NewSize(600, 400))
	w7.SetContent(content)
	w7.CenterOnScreen()
	w7.SetFixedSize(true)
	w7.Show()
}
