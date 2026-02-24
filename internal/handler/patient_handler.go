package handler

import (
	"agnos-hospital/internal/service"
	"net/http"
	"github.com/gin-gonic/gin"
)

type PatientHandler struct {
	service *service.PatientService
}

func NewPatientHandler(service *service.PatientService) *PatientHandler {
	return &PatientHandler{service: service}
}

func (h *PatientHandler) Search(c *gin.Context) {
	hospitalID := c.GetInt("hospital_id")

	nationalID := c.Query("national_id")
	passportID := c.Query("passport_id")

	patients, err := h.service.Search(hospitalID, nationalID, passportID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, patients)
}