package controllers

// ApiV1 is api v1 controller interface
type ApiV1 interface {
	WithUrlController(URL URL) ApiV1
	GetUrlController() URL
}
