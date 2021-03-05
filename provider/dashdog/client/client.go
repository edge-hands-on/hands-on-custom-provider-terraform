package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type DatadogClient struct {
	host   string
	apiKey string
	appKey string
}

func New(host string, apiKey string, appKey string) DatadogClient {
	client := DatadogClient{host, apiKey, appKey}
	return client
}

func (client DatadogClient) ValidateApiKey() bool {
	url := fmt.Sprintf("https://%s/api/v1/validate", client.host)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("DD-API-KEY", client.apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("StatusCode: %d", resp.StatusCode)
	log.Printf("ResponseBody: %s", string(bodyBytes))

	return resp.StatusCode >= 200 && resp.StatusCode <= 299
}

func (client DatadogClient) GetDashboard(id string) (interface{}, error) {
	url := fmt.Sprintf("https://%s/api/v1/dashboard/%s", client.host, id)

	log.Printf("Dashboard id: %s", id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("DD-API-KEY", client.apiKey)
	req.Header.Add("DD-APPLICATION-KEY", client.appKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	jsonMap := make(map[string]json.RawMessage)
	err = json.Unmarshal(bodyBytes, &jsonMap)
	if err != nil {
		panic(err)
	}

	log.Printf("StatusCode: %d", resp.StatusCode)
	log.Printf("ResponseBody: %s", string(bodyBytes))

	isOk := resp.StatusCode >= 200 && resp.StatusCode <= 299
	if isOk {
		return strings.Trim(string(jsonMap["id"]), "\""), nil
	}

	return nil, errors.New("error getting dashboard")
}

func (client DatadogClient) CreateDashboard(jsonData []byte) (interface{}, error) {
	url := fmt.Sprintf("https://%s/api/v1/dashboard", client.host)

	log.Printf("Dashboard json: %s", string(jsonData))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("DD-API-KEY", client.apiKey)
	req.Header.Add("DD-APPLICATION-KEY", client.appKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	jsonMap := make(map[string]json.RawMessage)
	err = json.Unmarshal(bodyBytes, &jsonMap)
	if err != nil {
		panic(err)
	}

	log.Printf("StatusCode: %d", resp.StatusCode)
	log.Printf("ResponseBody: %s", string(bodyBytes))

	isOk := resp.StatusCode >= 200 && resp.StatusCode <= 299
	if isOk {
		return strings.Trim(string(jsonMap["id"]), "\""), nil
	}

	return nil, errors.New("error creating dashboard")
}

func (client DatadogClient) UpdateDashboard(id string, jsonData []byte) error {
	url := fmt.Sprintf("https://%s/api/v1/dashboard/%s", client.host, id)

	log.Printf("Dashboard %s json: %s", id, string(jsonData))

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("DD-API-KEY", client.apiKey)
	req.Header.Add("DD-APPLICATION-KEY", client.appKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("StatusCode: %d", resp.StatusCode)
	log.Printf("ResponseBody: %s", string(bodyBytes))

	isOk := resp.StatusCode >= 200 && resp.StatusCode <= 299
	if isOk {
		return nil
	}

	return errors.New("error updating dashboard")
}

func (client DatadogClient) DeleteDashboard(id string) error {
	url := fmt.Sprintf("https://%s/api/v1/dashboard/%s", client.host, id)

	log.Printf("Dashboard id: %s", id)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("DD-API-KEY", client.apiKey)
	req.Header.Add("DD-APPLICATION-KEY", client.appKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("StatusCode: %d", resp.StatusCode)
	log.Printf("ResponseBody: %s", string(bodyBytes))

	isOk := resp.StatusCode >= 200 && resp.StatusCode <= 299
	if isOk {
		return nil
	}

	return errors.New("error deleting dashboard")
}
