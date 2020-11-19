package registry

import (
	"urlshortner/interface/controllers"
	"urlshortner/interface/interactors"
	"urlshortner/interface/repositories"
	uc "urlshortner/usecase/controllers"
	ui "urlshortner/usecase/interactors"
	ur "urlshortner/usecase/repositories"
)

func (self registry) newUrlRepository() ur.URL {
	return repositories.NewUrl(self.controller.Application.DBSession)
}

func (self registry) newUrlInteractors() ui.URL {
	return interactors.NewUrl(self.newUrlRepository())
}

func (self registry) newUrlController() uc.URL {
	return controllers.NewUrl(self.controller, self.newUrlInteractors())
}
