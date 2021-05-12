package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luckyshmo/api-example/config"
	"github.com/sirupsen/logrus"
	"golang.org/x/time/rate"
)

const defaultMaxCount = 1000000

func getRequestNumberFromConfig() int {
	n, err := strconv.Atoi(config.Get().MaxRequestCount)
	if err != nil {
		logrus.Warn("Using deafault requse max count")
		return defaultMaxCount
	}
	return n
}

var limiter = rate.NewLimiter(1, getRequestNumberFromConfig())

func (h *Handler) limit(c *gin.Context) {
	if limiter.Allow() == false {
		sendErrorResponse(c, http.StatusTooManyRequests, http.StatusText(429))
		return
	}
}
