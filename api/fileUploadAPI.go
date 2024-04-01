package api

import (
	"encoding/json"
	"os"
	"strings"
	"time"
)

type data struct {
	AffectedItems      []string `json:"affected_items"`
	TotalAffectedItems int      `json:"total_affected_items"`
	TotalFailedItems   int      `json:"total_failed_items"`
	FailedItems        []string `json:"failed_items"`
}

type drGeneralResponse struct {
	Data    data   `json:"data"`
	Message string `json:"message"`
	Error   int    `json:"error"`
}

func UploadFileAfterCreation(b bool) {
	readConfFile()
	var response drGeneralResponse
	var decoderXMLFilePath string = "./decoder.xml"

	nowString := time.Now().Format("2006-01-02T15:04:05.999999999Z07:00")
	nowString = strings.Replace(nowString, ":", "_", -1)
	nowString = strings.Replace(nowString, ".", "_", -1)

	decoderXMLContent, err := os.ReadFile(decoderXMLFilePath)
	if err != nil {
		panic(err)
	}

	decoderStringResponse := createRequest("PUT", "/decoders/files/"+nowString+"-decoder.xml", "application/octet-stream", decoderXMLContent)

	if err := json.Unmarshal([]byte(decoderStringResponse), &response); err != nil {
		panic(err)
	}

	if b {
		var ruleXMLFilePath string = "./rules.xml"
		ruleXMLContent, err := os.ReadFile(ruleXMLFilePath)
		if err != nil {
			panic(err)
		}
		ruleStringResponse := createRequest("PUT", "/rules/files/"+nowString+"-rule.xml", "application/octet-stream", ruleXMLContent)

		if err := json.Unmarshal([]byte(ruleStringResponse), &response); err != nil {
			panic(err)
		}
	}
}
