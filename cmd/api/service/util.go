package service

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

type Data struct {
	Id     int
	UserId int
	Title  string
	Body   string
}

type Payload struct {
	Data []Data
}

func Raw() ([]Data, error) {
	raw, err := os.ReadFile("data.json")
	if err != nil {
		return nil, err
	}

	var payload Payload
	err = json.Unmarshal(raw, &payload.Data)

	if err != nil {
		return nil, err
	}

	return payload.Data, nil
}

func GetAll() ([]Data, error) {
	data, err := Raw()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func GetById(id int) (any, error) {
	data, err := Raw()
	if err != nil {
		return nil, err
	}

	if id > len(data) {
		result := make([]string, 0)
		err1 := echo.NewHTTPError(http.StatusBadRequest, "Please provide valid id")
		return result, err1
	}

	return data[id], nil
}
