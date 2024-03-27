package main

import (
	"errors"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main(){
    a := app.New()
    w := a.NewWindow("Wazuh R&D Generator")
    
    red := color.NRGBA{R:180, G:0, B:0, A:255}
    errorLabel := canvas.NewText("Cannot Be Empty", red)
    errorLabel2 := canvas.NewText("Cannot Be Empty", red)
    errorLabel.Hide()
    errorLabel2.Hide()

    logLabel:= widget.NewLabel("Enter the log from where you want to extract the information")
    logEntries := widget.NewEntry()
    logEntries.SetPlaceHolder("Info - New agent connected: { \"name\": \"ExampleAgent\", \"ip\": \"192.168.1.100\", \"id\": \"001\"...")
    logEntries.Validator = func (input string) error {
        if len(logEntries.Text) == 0 {
            return errors.New("Cannot be empty")
        } else{
        }
        return nil
    }
    logEntries.OnChanged = func (str string) {
        if logEntries.Text == "" {
            errorLabel.Show()
            errorLabel.TextStyle.Italic = true
        } else {
            errorLabel.Hide()
        }
    }

    variablesLabel := widget.NewLabel("Type the variables that you want to read (multiple variables have to be separated by commas)")
    variablesEntries := widget.NewEntry()
    variablesEntries.SetPlaceHolder("name, ip, id...")
    variablesEntries.Validator = func (input string) error {
        if len(variablesEntries.Text) == 0 {
            return errors.New("Cannot be empty")
        } else{
        }
        return nil
    }

    variablesEntries.OnChanged = func (str string) {
        if variablesEntries.Text == "" {
            errorLabel2.Show()
            errorLabel2.TextStyle.Italic = true
        } else {
            errorLabel2.Hide()
        }
    }

    data := struct {
        log, variables string
    } {
        log : "",
        variables : "",
    }

    information := widget.NewLabel("")
    information.Hide()

    button := widget.NewButton("Submit", func () {
       if variablesEntries.Text != "" && logEntries.Text != "" {
           data.log = logEntries.Text 
           data.variables = variablesEntries.Text
           information.SetText("This is the information that we are processing: \n" + data.log + "\n" + data.variables) 
           information.Show()
       }
    })

    inputVBox := container.NewVBox(logLabel, logEntries, errorLabel, variablesLabel, variablesEntries, errorLabel2, layout.NewSpacer(), button, layout.NewSpacer(), information, layout.NewSpacer())
    content := container.NewHBox(layout.NewSpacer(), inputVBox, layout.NewSpacer())

    w.SetContent(content)

    w.Resize(fyne.NewSize(800, 600))
    w.SetFixedSize(true)

    w.Show()
    a.Run()
}

func processData(data struct {
    log, variables string
}) {
}
