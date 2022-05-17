package router

// Router contains the logic for setting up the handler for a specific path or a logical group of paths.
type Router interface {

	// Route setups a handler for the request.
	Route()
}
