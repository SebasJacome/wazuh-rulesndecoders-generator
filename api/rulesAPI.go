package api

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/pelletier/go-toml/v2"
)

// Configuration

var config = struct {
	Protocol string
	Host     string
	Port     string
	Username string
	Password string
}{
	Protocol: "https",
	Host:     "",
	Port:     "",
	Username: "",
	Password: "",
}

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

type ruleInfo struct {
	ID          string `json:"@id"`
	Level       string `json:"@level"`
	ProgramName string `json:"program_name"`
	Description string `json:"description"`
	Group       string `json:"group"`
}

type group struct {
	Name string   `json:"@name"`
	Rule ruleInfo `json:"rule"`
}

type AffectedItem struct {
	Group group `json:"group"`
}

type JSONRuleInfoResponse struct {
	Data struct {
		AffectedItems      []AffectedItem `json:"affected_items"`
		TotalAffectedItems int            `json:"total_affected_items"`
		TotalFailedItems   []string       `json:"failed_items"`
	} `json:"data"`
	Message  string `json:"message"`
	ErrorNum int    `json:"error"`
}

func RequestRuleIDs() {
	readConfFile()
	var response JSONRuleFileResponse
	var ruleInfoResponse JSONRuleInfoResponse
	stringResponse := createRequest("GET", "/rules/files")

	if err := json.Unmarshal([]byte(stringResponse), &response); err != nil {
		panic(err)
	}

	for _, itemI := range response.Data.AffectedItems {
		if itemI.RelativeDir == "etc/rules" {
			ruleInfo := createRequest("GET", "/rules/files/"+itemI.FileName)
			fmt.Println(ruleInfo)
			if err := json.Unmarshal([]byte(ruleInfo), &ruleInfoResponse); err != nil {
				panic(err)
			}
			for _, itemJ := range ruleInfoResponse.Data.AffectedItems {
				fmt.Println(itemJ.Group.Rule.ID)
			}
		}
	}

}

func readConfFile() {
	data, err := os.ReadFile("./conf.toml")
	if err != nil {
		panic(err)
	}
	if err := toml.Unmarshal([]byte(data), &config); err != nil {
		panic(err)
	}
}

func createRequest(requestType, endpoint string) string {
	// Disable insecure https warnings (for self-signed SSL certificates)
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	baseURL := fmt.Sprintf("%s://%s:%s", config.Protocol, config.Host, config.Port)
	loginURL := fmt.Sprintf("%s/security/user/authenticate", baseURL)

	token, err := getAuthToken(loginURL, config.Username, config.Password)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	// Add the Bearer token to the headers
	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + token
	headers["Content-Type"] = "application/json"

	response, err := getResponse(requestType, baseURL+endpoint, headers, nil)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return string(response)
}

func getAuthToken(url, user, password string) (string, error) {
	auth := user + ":" + password
	basicAuth := base64.StdEncoding.EncodeToString([]byte(auth))
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", "Basic "+basicAuth)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	token, ok := result["data"].(map[string]interface{})["token"].(string)
	if !ok {
		return "", fmt.Errorf("token not found in response")
	}

	return token, nil
}

func getResponse(method, url string, headers map[string]string, body []byte) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error obtaining response: status code %d", resp.StatusCode)
	}

	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return response, nil
}
