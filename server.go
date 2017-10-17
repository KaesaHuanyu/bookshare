package main

import (
	"github.com/labstack/echo"
	"bookshare/handlers"
	"gopkg.in/mgo.v2"
	"github.com/sirupsen/logrus"
	"os"
	"golang.org/x/crypto/acme/autocert"
)

var (
	MONGO = ""
)

func main() {
	e := echo.New()
	defer e.Close()
	session, _ := mgo.Dial(MONGO)
	//if err != nil {
	//	panic(err)
	//}
	defer session.Close()
	log := logrus.New()
	log.Out = os.Stdout
	h := handlers.New(session, log)
	e.AutoTLSManager.HostPolicy = autocert.HostWhitelist("127.0.0.1")
	e.AutoTLSManager.Cache = autocert.DirCache("var/www/.cache")
	e.GET("/v1/users/register", h.Register)
	e.POST("/v1/books/release", h.Release)
	e.GET("/v1/books", h.Query)
	e.GET("/v1/books/:id", h.Info)
	log.Fatal(e.StartAutoTLS(":4396"))
}
