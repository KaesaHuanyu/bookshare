package handlers

import (
	"gopkg.in/mgo.v2"
	"github.com/sirupsen/logrus"
	"github.com/labstack/echo"
)

type (
	handler struct {
		session *mgo.Session
		log *logrus.Logger
	}
)

func New(session *mgo.Session, log *logrus.Logger) *handler {
	return &handler{
		session: session,
		log: log,
	}
}

func (h *handler) Register(c echo.Context) error {
	return c.JSONPretty(0, nil, "")
}

func (h *handler) Release(c echo.Context) error {
	return c.JSONPretty(0, nil, "")
}

func (h *handler) Query(c echo.Context) error {
	return c.JSONPretty(0, nil, "")
}

func (h *handler) Info(c echo.Context) error {
	return c.JSONPretty(0, nil, "")
}