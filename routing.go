package S3Proxy

import "github.com/drone/routes"

func SetUpRoutes() *routes.RouteMux {
    // Using routes to give Regex functionality not offered by net/http
    mux := routes.New()
    // Route: Index
    mux.Get("/", IndexHandler)
    // Route: Status
    mux.Get("/_status", StatusHandler)
    // Route: Default
    mux.Get("/.*", DefaultHandler)

    return mux
}
