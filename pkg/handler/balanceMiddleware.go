package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(1, 3)

func (h *Handler) limit(c *gin.Context) {
	if limiter.Allow() == false {
		sendErrorResponse(c, http.StatusTooManyRequests, http.StatusText(429))
		return
	}
}
