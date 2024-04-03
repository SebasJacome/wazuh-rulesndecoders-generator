package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func UploadWindow() {
	myWindow := a.NewWindow("XML File Upload")

	fileLabel := widget.NewLabel("Select an XML file:")
	fileEntry := widget.NewEntry()
	fileOpen := widget.NewButtonWithIcon("Open File", theme.FileIcon(), func() {
		openFileDialog(myWindow, fileEntry)
	})

	uploadButton := widget.NewButton("Upload", func() {
		filename := fileEntry.Text
		if filename == "" {
			dialog.ShowInformation("Error", "Please select a file first", myWindow)
			return
		}

		fileBytes, err := os.ReadFile(filename)
		if err != nil {
			dialog.ShowError(err, myWindow)
			return
		}

		fileType, err := parseXMLFile(fileBytes)
		if err != nil {
			dialog.ShowError(err, myWindow)
			return
		}
		fmt.Println("File type:", fileType)

		dialog.ShowInformation("Success", "File uploaded successfully as "+getFilename(filename)+"."+"contents: "+string(fileBytes), myWindow)

	})

	content := container.New(layout.NewVBoxLayout(),
		fileLabel,
		container.New(layout.NewHBoxLayout(), fileEntry, fileOpen),
		uploadButton,
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.Show()
}

func parseXMLFile(fileBytes []byte) (string, error) {
	fileContent := string(fileBytes)
	// Check for well-formedness
	decoder := xml.NewDecoder(strings.NewReader(fileContent))
	for {
		_, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", errors.New("malformed XML file")
		}
	}

	// Check for root element
	rootElementRegex := regexp.MustCompile(`^<(decoder|group)\s+name="[\w\-]+">`)
	matches := rootElementRegex.FindStringSubmatch(fileContent)
	if len(matches) < 2 {
		return "", errors.New("invalid XML file, unable to determine root element (rule or decoder)")
	}

	allowedDecoderTags := []string{
		"name", "parent", "accumulate", "program_name", "prematch", "regex", "order",
		"fts", "ftscomment", "plugin_decoder", "use_own_name", "json_null_field",
		"json_array_structure", "var", "type",
	}

	allowedRuleTags := []string{
		"rule", "match", "regex", "decoded_as", "category", "field", "srcip", "dstip",
		"srcport", "dstport", "data", "extra_data", "user", "system_name", "program_name",
		"protocol", "hostname", "time", "weekday", "id", "url", "location", "action",
		"status", "srcgeoip", "dstgeoip", "if_sid", "if_group", "if_level", "if_matched_sid",
		"if_matched_group", "same_id", "different_id", "same_srcip", "different_srcip",
		"same_dstip", "different_dstip", "same_srcport", "different_srcport", "same_dstport",
		"different_dstport", "same_location", "different_location", "same_srcuser", "different_srcuser",
		"same_user", "different_user", "same_field", "different_field", "same_protocol",
		"different_protocol", "same_action", "different_action", "same_data", "different_data",
		"same_extra_data", "different_extra_data", "same_status", "different_status",
		"same_system_name", "different_system_name", "same_url", "different_url", "same_srcgeoip",
		"different_srcgeoip", "same_dstgeoip", "different_dstgeoip", "description", "list", "info",
		"options", "check_diff", "group", "mitre", "var",
	}

	switch matches[1] {
	case "decoder":
		// Check for disallowed tags in decoders
		tagRegex := regexp.MustCompile(`<(\w+)`)
		tags := tagRegex.FindAllString(fileContent, -1)
		for _, tag := range tags {
			tag = strings.TrimPrefix(tag, "<")
			if !stringInSlice(tag, allowedDecoderTags) {
				return "", fmt.Errorf("disallowed tag '%s' found in the decoder", tag)
			}
		}
		return "decoder", nil
	case "group":
		// Check for disallowed tags in rules
		tagRegex := regexp.MustCompile(`<(\w+)`)
		tags := tagRegex.FindAllString(fileContent, -1)
		for _, tag := range tags {
			tag = strings.TrimPrefix(tag, "<")
			if !stringInSlice(tag, allowedRuleTags) {
				return "", fmt.Errorf("disallowed tag '%s' found in the rules", tag)
			}
		}
		return "rules", nil
	default:
		return "", errors.New("unknown XML file type")
	}
}

func stringInSlice(str string, slice []string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

func getFilename(path string) string {
	parts := strings.Split(path, "/")
	return parts[len(parts)-1]
}

func openFileDialog(window fyne.Window, fileEntry *widget.Entry) {
	fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
		if reader == nil {
			return
		}
		if err == nil {
			fileEntry.SetText(reader.URI().Path())
		}
	}, window)
	fd.Show()
}
