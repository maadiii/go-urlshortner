package controllers

import (
	"net/http"
	"urlshortner/domain/interactors"
	"urlshortner/infrastructure/application"
	"urlshortner/usecase/controllers"
	ui "urlshortner/usecase/interactors"

	"github.com/julienschmidt/httprouter"
)

type url struct {
	application.Controller
	interactor ui.URL
}

func NewUrl(c application.Controller, i ui.URL) controllers.URL {
	return url{c, i}
}

func (self url) Redirect(ctx *application.Context) error {
	shortUrlParam := httprouter.ParamsFromContext(ctx.Request.Context()).ByName("short_url")
	output, err := self.interactor.Get(interactors.ShortedUrl{Url: shortUrlParam})
	if err != nil {
		return err
	}

	http.Redirect(ctx.Response, ctx.Request, output.OriginalUrl, http.StatusMovedPermanently)

	return nil
}

func (self url) Get(ctx *application.Context) error {
	shortUrlParam := httprouter.ParamsFromContext(ctx.Request.Context()).ByName("short_url")
	output, err := self.interactor.Get(interactors.ShortedUrl{Url: shortUrlParam})
	if err != nil {
		return err
	}

	return ctx.Finish(http.StatusOK, output)
}

func (self url) Post(ctx *application.Context) error {
	var input interactors.OriginalUrl
	if err := ctx.DecodeModel(&input); err != nil {
		return err
	}

	if output, err := self.interactor.Add(&input); err != nil {
		return err
	} else {

		return ctx.Finish(http.StatusOK, output)
	}
}
