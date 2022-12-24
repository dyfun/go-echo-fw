package handlers

import (
	"github.com/dyfun/go-echo-fw/cmd/api/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func PostHandler(c echo.Context) error {
	data, err := service.GetAll()
	if err != nil {
		c.String(http.StatusBadRequest, "Unable")
	}

	result := make(map[string]any)
	result["status"] = "Success"
	result["data"] = data

	return c.JSON(http.StatusOK, result)
}

func SingleHandler(c echo.Context) error {
	id := c.Param("id")
	newId, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	data, err := service.GetById(newId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	result := make(map[string]any)
	result["code"] = 200
	result["data"] = data

	return c.JSON(http.StatusOK, result)
}
