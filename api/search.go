package api

import (
	"encoding/json"
	"fmt"
)

type AffectedItem struct {
	FileName        string      `json:"filename"`
	RelativeDirName string      `json:"relative_dirname"`
	ID              int         `json:"id"`
	Level           int         `json:"level"`
	Status          string      `json:"status"`
	Details         interface{} `json:"details"`
	PciDss          []string    `json:"pci_dss"`
	Gpg13           []string    `json:"gpg13"`
	Gdpr            []string    `json:"gpdr"`
	Hipaa           []string    `json:"hipaa"`
	Nist80053       []string    `json:"nist_800_53"`
	TSC             []string    `json:"tsc"`
	Mitre           []string    `json:"mitre"`
	Groups          []string    `json:"groups"`
	Description     string      `json:"description"`
}

type Data struct {
	AffectedItems      []AffectedItem `json:"affected_items"`
	TotalAffectedItems int            `json:"total_affected_items"`
	TotalFailedItems   int            `json:"total_failed_items"`
	FailedItems        []interface{}  `json:"failed_items"`
}

type RuleResponse struct {
	Data    Data   `json:"data"`
	Message string `json:"message"`
	Error   int    `json:"error"`
}

func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

func SearchRequestedID(id int) {
	var str RuleResponse
	readConfFile()
	response := createRequest("GET", "/rules?relative_dirname=etc%2Frules", "application/json", nil)
	if err := json.Unmarshal([]byte(response), &str); err != nil {
		panic(err)
	}

	for _, value := range str.Data.AffectedItems {
		if value.ID == id {
			res, _ := PrettyStruct(value)
			fmt.Println(res)
		}
	}

}
