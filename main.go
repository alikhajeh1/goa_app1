package main

import (
	"github.com/alikhajeh1/goa_app1/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("API")

	// Setup middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	cs := NewSwaggerController(service)
  app.MountSwaggerController(service, cs)

	// Mount "numbers" controller
	c := NewNumbersController(service)
	app.MountNumbersController(service, c)

	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
