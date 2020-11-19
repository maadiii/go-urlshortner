package registry

import (
	"urlshortner/infrastructure/application"
	"urlshortner/infrastructure/datastore"
	"urlshortner/interface/controllers"
	uc "urlshortner/usecase/controllers"

	"github.com/julienschmidt/httprouter"
)

var GlobalRegistry Registry

// Registry interface
type Registry interface {
	GetController() application.Controller
	Route(router *httprouter.Router)
}

type registry struct {
	controller     application.Controller
	rootController uc.Root
}

// New creates and returns registry
func Init() error {
	app, err := application.New()
	if err != nil {
		return err
	}

	ctrl, err := application.NewController(app)
	if err != nil {
		return err
	}

	reg := registry{controller: ctrl}
	reg.init()
	GlobalRegistry = reg

	return nil
}

// NewTestRegistry creates and return registry for test goals
func InitForTest() error {
	app, err := application.New()
	if err != nil {
		return err
	}

	session, err := datastore.NewTestSession()
	if err != nil {
		return err
	}
	app.DBSession = *session

	c, err := application.NewController(app)
	if err != nil {
		return err
	}

	reg := registry{controller: c}
	reg.init()
	GlobalRegistry = reg

	return nil
}

func (self *registry) init() {
	apiv1 := controllers.NewApiV1(self.controller)
	apiv1 = apiv1.WithUrlController(self.newUrlController())

	root := controllers.NewRoot(self.controller)
	root = root.WithApiV1(apiv1)

	self.rootController = root
}

func (self registry) GetController() application.Controller {
	return self.controller
}
