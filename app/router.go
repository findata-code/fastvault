package app

type Router struct {
	route map[string]map[string]Handler
}

func NewRouter() *Router {
	route := make(map[string]map[string]Handler)
	return &Router{
		route: route,
	}
}

func (rt *Router) Get(method, path string) Handler {
	defer func () {
		if r := recover(); r != nil {}
	}()

	return rt.route[method][path]
}

func (rt *Router) Link(method, path string, handler Handler) {
	if rt.route[method] == nil {
		rt.route[method] = make(map[string]Handler)
	}

	rt.route[method][path] = handler
}
