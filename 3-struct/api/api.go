package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-course/bibin/config"
	"io"
	"net/http"
)

type BinApi struct {
	apiKey      string
	contentType string
	baseUrl     string
}

func NewBinApi(config *config.Config) *BinApi {
	return &BinApi{
		apiKey:      config.ApiKey,
		contentType: "application/json",
		baseUrl:     "https://api.jsonbin.io/v3/b",
	}
}

type CreateResponse struct {
	Metadata struct {
		Id      string `json:"id"`
		Private bool   `json:"private"`
	} `json:"metadata"`
}

type GetResponse struct {
	Record interface{} `json:"record"`
}

func (api *BinApi) Create(data *[]byte) (CreateResponse, error) {
	var result CreateResponse
	client := http.Client{}

	request, err := http.NewRequest("POST", api.baseUrl, bytes.NewBuffer(*data))
	if err != nil {
		return result, err
	}

	request.Header.Set("Content-Type", api.contentType)
	request.Header.Set("X-Master-Key", api.apiKey)

	response, err := client.Do(request)
	if err != nil {
		return result, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return result, err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return result, fmt.Errorf("ошибка при сохранении данных по api, код ошибки: %d, текст ошибки: %s", response.StatusCode, string(body))
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (api *BinApi) Get(id string) (string, error) {
	client := http.Client{}

	request, err := http.NewRequest("GET", api.baseUrl+"/"+id, nil)
	if err != nil {
		return "", err
	}

	request.Header.Set("Content-Type", api.contentType)
	request.Header.Set("X-Master-Key", api.apiKey)

	response, err := client.Do(request)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return "", fmt.Errorf("ошибка при получении данных по api, код ошибки: %d, текст ошибки: %s", response.StatusCode, string(body))
	}

	var getResponse GetResponse

	err = json.Unmarshal(body, &getResponse)
	if err != nil {
		return "", err
	}

	data, err := json.MarshalIndent(getResponse.Record, "", "  ")
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (api *BinApi) Update(id string, data *[]byte) error {
	client := http.Client{}

	request, err := http.NewRequest("PUT", api.baseUrl+"/"+id, bytes.NewBuffer(*data))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", api.contentType)
	request.Header.Set("X-Master-Key", api.apiKey)

	response, err := client.Do(request)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return fmt.Errorf("ошибка при обновлении данных по api, код ошибки: %d, текст ошибки: %s", response.StatusCode, string(body))
	}

	return nil
}

func (api *BinApi) Delete(id string) error {
	client := http.Client{}

	request, err := http.NewRequest("DELETE", api.baseUrl+"/"+id, nil)
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", api.contentType)
	request.Header.Set("X-Master-Key", api.apiKey)

	response, err := client.Do(request)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return fmt.Errorf("ошибка при удалении данных по api, код ошибки: %d, текст ошибки: %s", response.StatusCode, string(body))
	}

	return nil
}
