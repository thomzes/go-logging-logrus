package gologginglogrus

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestOutput(t *testing.T) {
	logger := logrus.New()

	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	logger.SetOutput(file)
	logger.Info("Hello Logging")
	logger.Warn("Hello logging")
	logger.Error("Hello logging")
}
