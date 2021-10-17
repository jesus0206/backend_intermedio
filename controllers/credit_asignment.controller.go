package controllers

import (
	"net/http"
	"yofio/intermedio/controllers/dto/input"
	"yofio/intermedio/controllers/dto/output"

	"github.com/gin-gonic/gin"
	// "jesus.tn79/ficha_electronica/models"
)

func (con Controller) CreditAssignment(c *gin.Context) {
	var json input.CreditAsignmentRequestDto
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, output.ResponseError{Error: "Debe enviar los datos como json.."})
		return
	}
	a, b, c_300, err := con.repo.CreditAssignment(int32(json.Investment))
	if err != nil {
		c.JSON(http.StatusInternalServerError, output.ResponseError{Error: err.Error()})
		return
	}
	var result output.ResponseDto
	result.CreditType300 = a
	result.CreditType500 = b
	result.CreditType700 = c_300
	c.JSON(http.StatusOK, result)
}
