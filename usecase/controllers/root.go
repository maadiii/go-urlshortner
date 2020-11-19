package controllers

import "urlshortner/infrastructure/application"

// Root is root controller interface
type Root interface {
	BaseController() application.Controller
	WithApiV1(ApiV1) Root
	ApiV1() ApiV1
}
