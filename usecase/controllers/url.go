package controllers

import "urlshortner/infrastructure/application"

// URL is url controller interface
type URL interface {
	Post(ctx *application.Context) error
	Get(ctx *application.Context) error
	Redirect(ctx *application.Context) error
}
