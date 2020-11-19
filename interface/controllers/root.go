package controllers

import (
	"urlshortner/infrastructure/application"
	"urlshortner/usecase/controllers"
)

type root struct {
	application.Controller
	apiv1 controllers.ApiV1
}

// NewRootController creates and returns root controller
func NewRoot(c application.Controller) controllers.Root {
	return &root{Controller: c}
}

func (self root) BaseController() application.Controller {
	return self.Controller
}

func (self *root) ApiV1() controllers.ApiV1 {
	return self.apiv1
}

func (self *root) WithApiV1(apiv1 controllers.ApiV1) controllers.Root {
	self.apiv1 = apiv1
	return self
}
