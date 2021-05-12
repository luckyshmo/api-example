package service

import (
	"time"

	"github.com/luckyshmo/api-example/models"
	"github.com/sirupsen/logrus"
)

type DataParser struct {
}

func NewDataParserService() *DataParser {
	return &DataParser{}
}

func (s *DataParser) ParseData(data []models.Data) error {
	logrus.Info(data)
	time.Sleep(time.Second) //some Data process
	return nil
}
