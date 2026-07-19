package core_http_server

import (
	"net/http"
)

type ApiVersion string

var (
	ApiVersion1 = ApiVersion("v1")
	ApiVersion2 = ApiVersion("v2")
	ApiVersion3 = ApiVersion("v3")
)

type APIVersionRouter struct {
	*http.ServeMux
	apiVersion ApiVersion
}

func NewAPIVersionRouter(apiVersion ApiVersion) *APIVersionRouter {
	return &APIVersionRouter{
		ServeMux:   http.NewServeMux(),
		apiVersion: apiVersion,
	}
}

/*func (r *APIVersionRouter) RegistrRoutes(routes ...Route) {
	for _, route := range routes {
		pattern := fmt.Sprintf("/%s%s", r.apiVersion, route.Path)
		fmt.Println(pattern)
		r.Handle(pattern, route.Handler)
	}

}*/

func (r *APIVersionRouter) RegistrRoutes(routes ...Route) {
	// Группируем маршруты по пути
	routesByPath := make(map[string]map[string]http.Handler)
	for _, route := range routes {
		if _, ok := routesByPath[route.Path]; !ok {
			routesByPath[route.Path] = make(map[string]http.Handler)
		}
		routesByPath[route.Path][route.Method] = route.Handler
	}

	// Регистрируем один обработчик для каждого пути
	for path, handlers := range routesByPath {
		r.HandleFunc(path, func(rw http.ResponseWriter, r *http.Request) {
			if handler, ok := handlers[r.Method]; ok {
				handler.ServeHTTP(rw, r)
			} else {
				http.Error(rw, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			}
		})
	}
}
