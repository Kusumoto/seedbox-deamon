package api

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/cors"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

type httpResponseStructure struct {
	Status  int    `json:"Status"`
	Message string `json:"Message"`
}

// StartAPIServer start the API server
func StartAPIServer() {
	app := iris.New()
	app.Adapt(
		// iris.DevLogger(),
		httprouter.New(),
		cors.New(cors.Options{AllowedOrigins: []string{"*"}}))

	// Route: /
	// Method: "GET"
	// Hello Wolrd!
	app.Get("/", func(ctx *iris.Context) {
		ctx.JSON(200, httpResponseStructure{
			Message: "seedbox-daemon up and running.",
			Status:  200,
		})
	})

	// Route: /api/installer
	// Method: "GET"
	// Initization and pair with web portal control panel
	app.Get("/api/installer", func(ctx *iris.Context) {
		ctx.JSON(200, httpResponseStructure{
			Message: "seedbox-daemon initization and pair with control panel.",
			Status:  200,
		})
	})

	// Route: /api/status
	// Method: "GET"
	// Inside docker and container infomation
	app.Get("/api/status", func(ctx *iris.Context) {
		ctx.JSON(200, httpResponseStructure{
			Message: "seedbox-daemon inside docker and container infomation.",
			Status:  200,
		})
	})

	// Start the server at 0.0.0.0:4444
	app.Listen(":4444")
}
