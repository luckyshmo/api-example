package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luckyshmo/api-example/models"
)

// @Summary parseData
// @Tags data parse
// @Description parseData
// @Accept  json
// @Produce  json
// @Param input body models.Data
// @Success 200 {integer} integer 1
// @Failure default {object} errorResponse
// @Router /data [post]
func (h *Handler) receiveData(c *gin.Context) {

	var input []models.Data

	if err := c.BindJSON(&input); err != nil {
		sendErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err := h.services.ParseData(input)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendStatusResponse(c, http.StatusOK, map[string]interface{}{
		"parsed": len(input), //JSON body
	})
}
