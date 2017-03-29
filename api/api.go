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
		installationAPI := daemonAPI.Party("/installation")
		{
			// http://0.0.0.0:4444/api/installation
			// Method: "POST"
			installationAPI.Post("/", postEndPointInstallation)
			// http://0.0.0.0:4444/api/installation/callback
			// Method: "Post"
			installationAPI.Post("/callback", postEndPointInstallationCallback)
		}
		containerAPI := daemonAPI.Party("/container")
		{
			// http://0.0.0.0:4444/api/container/create
			// Method: "Post"
			containerAPI.Post("/create", postCreateContainer)
			// http://0.0.0.0:4444/api/container/remove
			// Method: "Post"
			containerAPI.Post("/remove", postRemoveContainer)
		}

	}
	app.Listen(":4444")
}

func endpointAPIMiddleware(ctx *iris.Context) {
	println("Request: " + ctx.Path())
	ctx.Next()
}

func endPointNotFoundHandler(ctx *iris.Context) {
	ctx.JSON(iris.StatusNotFound, httpResponseStructure{Status: 404, Message: "endpoint not found!"})
}

func postEndPointInstallation(ctx *iris.Context) {
	resultData := installationResult{}
	resultData.installEndpoint()
	ctx.JSON(resultData.ResponseStatus, resultData)
}

func postEndPointInstallationCallback(ctx *iris.Context) {
	resultData := installationResult{}
	resultData.installEndpoint()
	ctx.JSON(resultData.ResponseStatus, resultData.ResponseAccessToken)
}

func postCreateContainer(ctx *iris.Context) {

}

func postRemoveContainer(ctx *iris.Context) {

}
