package api

import (
	"math/rand"
	"time"

	iris "gopkg.in/kataras/iris.v6"
)

type installationResult struct {
	ResponseStatus      int    `json:"response_status"`
	ResponseMessage     string `json:"response_message"`
	ResponseAccessToken string `json:"response_access_token"`
}

func (result *installationResult) installEndpoint() {
	result.checkInstallationProcess()
	result.generateAccessToken()
}

func (result *installationResult) generateAccessToken() {

	if result.ResponseStatus != iris.StatusOK {
		return
	}

	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	tokenResult := make([]byte, defaultAccessTokenLength)
	for i := 0; i < defaultAccessTokenLength; i++ {
		tokenResult[i] = chars[rand.Intn(len(chars))]
	}
	result.ResponseMessage = "success"
	result.ResponseStatus = iris.StatusOK
	result.ResponseAccessToken = string(tokenResult)
}

func (result *installationResult) checkInstallationProcess() {
	result.ResponseMessage = "success"
	result.ResponseStatus = iris.StatusOK
}
