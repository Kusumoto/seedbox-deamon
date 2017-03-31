package api

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/cors"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

type httpResponseStructure struct {
	Status  int    `json:"response_status"`
	Message string `json:"response_message"`
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
	// Fire endPointInternalServerErrorHandler when Internal Server Error
	app.OnError(500, endPointInternalServerErrorHandler)

	daemonAPI := app.Party("/api", endpointAPIMiddleware)
	{
		installationAPI := daemonAPI.Party("/installation")
		{
			// http://0.0.0.0:4444/api/installation
			// Method: "Get"
			installationAPI.Get("/", postEndPointInstallation)
			// http://0.0.0.0:4444/api/installation/callback
			// Method: "Post"
			installationAPI.Post("/callback", postEndPointInstallationCallback)
			// http://0.0.0.0:4444/api/installation/reset
			// Method: "Post"
			installationAPI.Post("/reset", postEndPointInstallationReset)
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

func endPointInternalServerErrorHandler(ctx *iris.Context) {
	println("Error:" + ctx.Err().Error())
	ctx.JSON(iris.StatusInternalServerError, httpResponseStructure{Status: iris.StatusInternalServerError, Message: "internal server error!"})
}

func endPointNotFoundHandler(ctx *iris.Context) {
	ctx.JSON(iris.StatusNotFound, httpResponseStructure{Status: iris.StatusNotFound, Message: "endpoint not found!"})
}

func postEndPointInstallationReset(ctx *iris.Context) {
	resultData := installationResult{}
	requestData := installationRequestData{}
	err := ctx.ReadJSON(&requestData)
	if err != nil {
		ctx.JSON(iris.StatusInternalServerError, httpResponseStructure{Status: iris.StatusInternalServerError, Message: err.Error()})
		return
	}
	resultData.removeEndpoint(&requestData)
	ctx.JSON(resultData.ResponseStatus, resultData)
}

func postEndPointInstallation(ctx *iris.Context) {
	resultData := installationResult{}
	resultData.installEndpoint()
	ctx.JSON(resultData.ResponseStatus, resultData)
}

func postEndPointInstallationCallback(ctx *iris.Context) {
	requestData := installationRequestData{}
	resultData := installationResult{}
	err := ctx.ReadJSON(&requestData)
	if err != nil {
		ctx.JSON(iris.StatusInternalServerError, httpResponseStructure{Status: iris.StatusInternalServerError, Message: err.Error()})
		return
	}
	resultData.installCallbackEndpoint(&requestData)
	ctx.JSON(resultData.ResponseStatus, resultData)
}

func postCreateContainer(ctx *iris.Context) {
	resultData := createNewContainerResult{}
	client, err := NewDockerClient()
	if err != nil {
		ctx.JSON(iris.StatusInternalServerError, httpResponseStructure{Status: iris.StatusInternalServerError, Message: err.Error()})
		return
	}
	err = ctx.ReadJSON(&resultData)
	if err != nil {
		ctx.JSON(iris.StatusInternalServerError, httpResponseStructure{Status: iris.StatusInternalServerError, Message: err.Error()})
		return
	}
	resultData.createNewContainerEndpoint(client)
	ctx.JSON(resultData.ResponseStatus, resultData)
}

func postRemoveContainer(ctx *iris.Context) {

}
