package handlers

import (
	"net/http"
	"personal-income-tax-calculation-service/internal/logic"
	"personal-income-tax-calculation-service/pkg/apis"

	"github.com/gin-gonic/gin"
)

func RenderTaxPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func TaxCaculate(c *gin.Context) {
	var req apis.TaxRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, apis.TaxResponse{
			Message: "invalid params",
		})
		return
	}

	tax := logic.CalculateTax(req.Income)

	c.JSON(http.StatusOK, apis.TaxResponse{
		Message: "ok",
		Tax:     tax,
	})
}
