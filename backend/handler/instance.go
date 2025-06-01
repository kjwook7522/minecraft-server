package handler

import (
	"net/http"

	"minecraft-server/service"

	"github.com/labstack/echo/v4"
)

type InstanceHandler struct {
	instanceService *service.InstanceService
}

func NewInstanceHandler(instanceService *service.InstanceService) *InstanceHandler {
	return &InstanceHandler{
		instanceService: instanceService,
	}
}

func (h *InstanceHandler) StartServer(c echo.Context) error {
	if err := h.instanceService.StartServer(c.Request().Context()); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Server is starting...",
	})
}

func (h *InstanceHandler) StopServer(c echo.Context) error {
	if err := h.instanceService.StopServer(c.Request().Context()); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Server is stopping...",
	})
}

func (h *InstanceHandler) GetServerStatus(c echo.Context) error {
	status, err := h.instanceService.GetServerStatus(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"status": status,
	})
}
