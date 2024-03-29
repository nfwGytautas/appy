package main

import (
	"errors"
	"net/http"

	"github.com/nfwGytautas/appy"
	appy_default_drivers "github.com/nfwGytautas/appy/defaults"
)

func main() {
	options := appy.AppyOptions{
		Environment: appy.DefaultEnvironment(),
		Logger: appy.LoggerOptions{
			Name:     "Appy",
			Provider: appy_default_drivers.DefaultLogger(),
		},
		HTTP: &appy.HttpOptions{
			Provider: appy_default_drivers.DefaultHttpServer(),
			Address:  "127.0.0.1:8080",
			SSL:      nil, // HTTP
		},
	}

	// Create
	app, err := appy.New(options)
	if err != nil {
		panic(err)
	}

	// Add an endpoint handler
	app.Http().RootGroup().GET("/hello", func(c appy.HttpContext) appy.HttpResult {
		return c.Ok(http.StatusNoContent, nil)
	})

	app.Http().RootGroup().GET("/fail", func(c appy.HttpContext) appy.HttpResult {
		return c.Fail(http.StatusBadRequest, "Bad request")
	})

	app.Http().RootGroup().GET("/error", func(c appy.HttpContext) appy.HttpResult {
		return c.Error(errors.New("something went wrong"))
	})

	app.Http().RootGroup().GET("/services", func(c appy.HttpContext) appy.HttpResult {
		// Access appy services from http request
		c.App.Logger.Debug("Logger service accessed")
		return c.Ok(http.StatusNoContent, nil)
	})

	// Run
	err = app.Run()
	if err != nil {
		panic(err)
	}
}
