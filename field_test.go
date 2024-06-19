package gologginglogrus

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "Thomas Ardiansah").Info("Hello log")

	logger.WithField("username", "Thomas").
		WithField("name", "Thomas Ajadah").
		Info("Hello World!")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "Thomas",
		"name":     "Thomas Ardiansah",
	}).Info("Hello World!")
}
