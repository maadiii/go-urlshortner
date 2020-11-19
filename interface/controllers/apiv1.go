package controllers

import (
	"urlshortner/infrastructure/application"
	"urlshortner/usecase/controllers"
)

type apiv1 struct {
	application.Controller
	UrlController controllers.URL
}

// NewApiv1Controller creates and returns apiv1 controller
func NewApiV1(c application.Controller) controllers.ApiV1 {
	return &apiv1{Controller: c}
}

func (self *apiv1) WithUrlController(c controllers.URL) controllers.ApiV1 {
	self.UrlController = c
	return self
}

func (self *apiv1) GetUrlController() controllers.URL {
	return self.UrlController
}
