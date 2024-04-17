package api

import (
	"encoding/json"
	"regexp"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

type ruleFile struct {
	FileName    string `json:"filename"`
	RelativeDir string `json:"relative_dirname"`
	Status      string `json:"status"`
}

type JSONRuleFileResponse struct {
	Data struct {
		AffectedItems      []ruleFile `json:"affected_items"`
		TotalAffectedItems int        `json:"total_affected_items"`
		TotalFailedItems   []string   `json:"failed_items"`
	} `json:"data"`
	Message  string `json:"message"`
	ErrorNum int    `json:"error"`
}

func RequestRuleIDs(w3 fyne.Window) map[string]bool {
	var response JSONRuleFileResponse
	ruleIDs := make(map[string]bool)
	stringResponse := createRequest("GET", "/rules/files", "application/json", nil)
	if err := json.Unmarshal([]byte(stringResponse), &response); err != nil {
		panic(err)
	}

	for _, itemI := range response.Data.AffectedItems {
		if itemI.RelativeDir == "etc/rules" {
			ruleInfo := createRequest("GET", "/rules/files/"+itemI.FileName, "application/json", nil)

			r := regexp.MustCompile(`"@id": "[1-9][0-9]*"`)
			matches := r.FindAllString(ruleInfo, -1)
			if matches == nil {
				dialog.ShowInformation("Corrupted Rule File", "The file "+itemI.FileName+" is corrupted or it has no rules. Do something about it", w3)
			}

			for _, itemJ := range matches {
				parts := strings.Split(itemJ, " ")
				if len(parts) == 2 {
					parts[1] = strings.Trim(parts[1], `"`)
					ruleIDs[parts[1]] = true
				}
			}
		}
	}

	return ruleIDs
}
