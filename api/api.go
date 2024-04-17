package api

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/BurntSushi/toml"
)

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

func ReadConfFile(conf []byte) {
	if _, err := toml.Decode(string(conf), &config); err != nil {
		fmt.Println("Error parsing TOML config: ", err)
	}
}

func makeRulesRequestHeader(contentType string, token string) map[string]string {
	headers := make(map[string]string)

	headers["Authorization"] = "Bearer " + token
	headers["Content-Type"] = contentType

	return headers
}

func createRequest(requestType, endpoint string, contentType string, body []byte) string {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	baseURL := fmt.Sprintf("%s://%s:%s", config.Protocol, config.Host, config.Port)
	loginURL := fmt.Sprintf("%s/security/user/authenticate", baseURL)

	token, err := getAuthToken(loginURL, config.Username, config.Password)
	if err != nil {
		panic(err)
	}

	headers := makeRulesRequestHeader(contentType, token)

	if body == nil {
		response, err := getResponse(requestType, baseURL+endpoint, headers, nil)
		if err != nil {
			panic(err)
		}
		return string(response)
	} else {
		response, err := getResponse(requestType, baseURL+endpoint, headers, body)
		if err != nil {
			panic(err)
		}
		return string(response)
	}
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

func getResponse(method string, url string, headers map[string]string, body []byte) ([]byte, error) {
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
