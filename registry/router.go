package registry

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (self registry) Route(router *httprouter.Router) {
	handle := self.controller.Handle
	urlv1Controller := self.rootController.ApiV1().GetUrlController()

	router.Handler(http.MethodPost, "/v1/urls", handle(urlv1Controller.Post))
  router.Handler(http.MethodGet, "/v1/urls/:short_url", handle(urlv1Controller.Get))

	// use custome verv
  router.Handler("REDIRECT", "/v1/urls/:short_url", handle(urlv1Controller.Redirect))
}
