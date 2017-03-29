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
		iris.DevLogger(),
		httprouter.New(),
		cors.New(cors.Options{AllowedOrigins: []string{"*"}}))

	// Fire endPointNotFoundHandler when Not Found
	app.OnError(404, endPointNotFoundHandler)

	daemonAPI := app.Party("/api", endpointAPIMiddleware)
	{
		// http://0.0.0.0:4444/api/installation
		// Method: "GET"
		daemonAPI.Get("/installation", postEndPointInstallation)
	}
	app.Listen(":4444")
}

func endpointAPIMiddleware(ctx *iris.Context) {
	println("Request: " + ctx.Path())
	ctx.Next()
}

func endPointNotFoundHandler(ctx *iris.Context) {
	ctx.JSON(404, httpResponseStructure{Status: 404, Message: "Endpoint not found!"})
}

func postEndPointInstallation(ctx *iris.Context) {
	resultData := installationResult{}
	resultData.installEndpoint()
	ctx.JSON(resultData.responseStatus, resultData.responseAccessToken)
}

func getByIDHandler(ctx *iris.Context) {

}

func saveUserHandler(ctx *iris.Context) {
	// your code here...
}
