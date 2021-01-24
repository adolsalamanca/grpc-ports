package _interface

import (
	"fmt"
	"github.com/adolsalamanca/ports/client/domain/entity"
	"github.com/adolsalamanca/ports/client/domain/repository"
	"github.com/adolsalamanca/ports/client/infrastructure/common"
	"github.com/labstack/echo/v4"
	"net/http"
)

const (
	ErrNotFound = common.Error("Ports were not found")
	ErrNotValidRequestBody = common.Error("The request of the body is invalid")
	ErrInternalServerErr = common.Error("Could not process the request due to a server problem")
)

type PortHandler struct {
	repository repository.PortRepository
}

// Instantiates a new handler to interact with port REST API
func NewPortHandler(repository repository.PortRepository) *PortHandler {
	return &PortHandler{repository: repository}
}

// GetPorts retrieves ports data from the handler's repository
func (h PortHandler) GetPorts(c echo.Context) error {
	ports, err := h.repository.GetAllPorts()
	// TODO: Switch errors to understand if the response should be different here
	if err != nil {
		return c.JSON(http.StatusNotFound, ErrNotFound)
	}

	return c.JSON(http.StatusOK, ports)
}

// StorePorts stores ports data into handler's repository
func (h PortHandler) StorePorts(c echo.Context) error {
	var ports entity.Port
	err := c.Bind(&ports)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrNotValidRequestBody)
	}

	var portArr []entity.PortInfo
	for _, v := range ports {
		portArr = append(portArr, v)
	}

	err = h.repository.StorePorts(portArr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrInternalServerErr)
	}

	return c.JSON(http.StatusOK, fmt.Sprintf("Stored %d port elements" , len(portArr)))
}