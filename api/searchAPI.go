package api

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type RuleAffectedItem struct {
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

type RuleData struct {
	AffectedItems      []RuleAffectedItem `json:"affected_items"`
	TotalAffectedItems int                `json:"total_affected_items"`
	TotalFailedItems   int                `json:"total_failed_items"`
	FailedItems        []interface{}      `json:"failed_items"`
}

type RuleResponse struct {
	Data    RuleData `json:"data"`
	Message string   `json:"message"`
	Error   int      `json:"error"`
}

type MatchingRule struct {
	ID          int
	Description string
	FileName    string
	Level       int
	DirName     string
	Status      string
}

type MatchingDecoder struct {
	Name            string
	FileName        string
	RelativeDirName string
	Status          string
	Parent          string
	Regex           string
	Prematch        string
}

type DecoderDetails struct {
	Parent   string       `json:"parent"`
	Prematch PrematchInfo `json:"prematch"`
	Regex    RegexInfo    `json:"regex"`
	Order    string       `json:"order"`
}

type PrematchInfo struct {
	Pattern string `json:"pattern"`
	Offset  string `json:"offset"`
}

type RegexInfo struct {
	Pattern string `json:"pattern"`
	Offset  string `json:"offset"`
}

type DecoderAffectedItem struct {
	FileName        string         `json:"filename"`
	RelativeDirName string         `json:"relative_dirname"`
	Status          string         `json:"status"`
	Name            string         `json:"name"`
	Position        int            `json:"position"`
	Level           int            `json:"level"`
	Details         DecoderDetails `json:"details"`
}

type DecoderData struct {
	AffectedItems      []DecoderAffectedItem `json:"affected_items"`
	TotalAffectedItems int                   `json:"total_affected_items"`
	TotalFailedItems   int                   `json:"total_failed_items"`
	FailedItems        []interface{}         `json:"failed_items"`
}

type DecoderResponse struct {
	Data    DecoderData `json:"data"`
	Message string      `json:"message"`
	Error   int         `json:"error"`
}

func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

func SearchRequestedID(id int) MatchingRule {
	var str RuleResponse
	var result MatchingRule
	readConfFile()
	response := createRequest("GET", "/rules?relative_dirname=etc%2Frules", "application/json", nil)
	if err := json.Unmarshal([]byte(response), &str); err != nil {
		panic(err)
	}

	for _, value := range str.Data.AffectedItems {
		if value.ID == id {
			result.Description = value.Description
			result.ID = value.ID
			result.Level = value.Level
			result.FileName = value.FileName
			result.Status = value.Status
			result.DirName = value.RelativeDirName
			return result
		}
	}
	result.ID = -1
	result.Level = -1
	result.Description = "null"
	return result
}

func SearchRequestedName(name string) MatchingDecoder {
	var str DecoderResponse
	var result MatchingDecoder
	readConfFile()
	response := createRequest("GET", "/decoders?relative_dirname=etc%2Fdecoders", "application/json", nil)
	if err := json.Unmarshal([]byte(response), &str); err != nil {
		panic(err)
	}

	for _, value := range str.Data.AffectedItems {
		if value.Name == name {
			result.Name = value.Name
			result.FileName = value.FileName
			result.Status = value.Status
			result.RelativeDirName = value.RelativeDirName
			result.Parent = value.Details.Parent
			result.Prematch = value.Details.Prematch.Pattern
			result.Regex = value.Details.Regex.Pattern
			return result
		}
	}
	result.Name = ""
	result.FileName = ""
	result.Regex = ""
	return result
}

func SearchRequestedParameters(values []string) (bool, string) {
	var str RuleResponse
	var FoundValues []bool
	var errorMessage string = ""
	var all_found bool = true
	readConfFile()
	response := createRequest("GET", "/rules?relative_dirname=etc%2Frules", "application/json", nil)
	if err := json.Unmarshal([]byte(response), &str); err != nil {
		panic(err)
	}
	for i := range len(values) {
		fmt.Println(i)
		FoundValues = append(FoundValues, false)
	}
	for valueIndex, value := range values {
		for _, existingRule := range str.Data.AffectedItems {
			strval, err := strconv.Atoi(value)
			if err != nil {
				panic(err)
			}
			if existingRule.ID == strval {
				FoundValues[valueIndex] = true
			}
		}
	}
	for index, value := range FoundValues {
		if !value {
			errorMessage += values[index] + ","
			all_found = false
		}
	}
	if all_found {
		return true, ""
	} else {
		errorMessage = "Referenced Rule IDs " + errorMessage[:len(errorMessage)-1] + " did not match any existing rule ID"
		return false, errorMessage
	}
}

func SearchRequestedDecoder(values []string) (bool, string) {
	var str DecoderResponse
	var FoundValues []bool
	var errorMessage string = ""
	var all_found bool = true
	readConfFile()
	response := createRequest("GET", "/decoders?relative_dirname=etc%2Fdecoders", "application/json", nil)
	if err := json.Unmarshal([]byte(response), &str); err != nil {
		panic(err)
	}
	for i := range len(values) {
		fmt.Println(i)
		FoundValues = append(FoundValues, false)
	}
	for valueIndex, value := range values {
		for _, existingDecoder := range str.Data.AffectedItems {
			if existingDecoder.Name == value {
				FoundValues[valueIndex] = true
			}
		}
	}
	for index, value := range FoundValues {
		if !value {
			errorMessage += values[index] + ","
			all_found = false
		}
	}
	if all_found {
		return true, ""
	} else {
		errorMessage = "Decoder_as " + errorMessage[:len(errorMessage)-1] + " did not match any existing decoder name"
		return false, errorMessage
	}
}

func SearchForAllIDs() []MatchingRule {
	var str RuleResponse
	var result MatchingRule
	var results []MatchingRule
	readConfFile()
	response := createRequest("GET", "/rules?relative_dirname=etc%2Frules", "application/json", nil)
	if err := json.Unmarshal([]byte(response), &str); err != nil {
		panic(err)
	}

	for _, value := range str.Data.AffectedItems {
		result.Description = value.Description
		result.ID = value.ID
		result.Level = value.Level
		result.FileName = value.FileName
		result.Status = value.Status
		result.DirName = value.RelativeDirName
		results = append(results, result)
	}
	return results
}

func SearchForAllDecoders() []MatchingDecoder {
	var str DecoderResponse
	var result MatchingDecoder
	var results []MatchingDecoder
	readConfFile()
	response := createRequest("GET", "/decoders?relative_dirname=etc%2Fdecoders", "application/json", nil)
	if err := json.Unmarshal([]byte(response), &str); err != nil {
		panic(err)
	}

	fmt.Println(str)

	for _, value := range str.Data.AffectedItems {
		result.Name = value.Name
		result.FileName = value.FileName
		result.Status = value.Status
		result.RelativeDirName = value.RelativeDirName
		result.Parent = value.Details.Parent
		result.Prematch = value.Details.Prematch.Pattern
		result.Regex = value.Details.Regex.Pattern
		results = append(results, result)
	}

	return results
}
