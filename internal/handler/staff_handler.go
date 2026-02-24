package handler

import (
	"agnos-hospital/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StaffHandler struct {
	authService *service.AuthService
}

func NewStaffHandler(authService *service.AuthService) *StaffHandler {
	return &StaffHandler{authService: authService}
}

func (h *StaffHandler) Register(c *gin.Context) {
	var req struct {
		Username   string `json:"username"`
		Password   string `json:"password"`
		HospitalID int    `json:"hospital_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	err := h.authService.Register(req.Username, req.Password, req.HospitalID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "staff created"})
}

func (h *StaffHandler) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	token, err := h.authService.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}