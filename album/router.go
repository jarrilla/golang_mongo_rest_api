package album

import (
	"net/http"

	"github.com/gorilla/mux"

	"musicstore/logger"
)

var controller = &Controller{Repository: Repository{}}

// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		controller.Index,
	},
	Route{
		"AddAlbum",
		"POST",
		"/",
		controller.AddAlbum,
	},
	Route{
		"UpdateAlbum",
		"PUT",
		"/",
		controller.UpdateAlbum,
	},
	Route{
		"DeleteAlbum",
		"DELETE",
		"/{id}",
		controller.DeleteAlbum,
	},
}

// NewRouter configures a new router to the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
