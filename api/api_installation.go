package api

import (
	"math/rand"
	"time"
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
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	tokenResult := make([]byte, defaultAccessTokenLength)
	for i := 0; i < defaultAccessTokenLength; i++ {
		tokenResult[i] = chars[rand.Intn(len(chars))]
	}
	result.ResponseMessage = "Success"
	result.ResponseStatus = 200
	result.ResponseAccessToken = string(tokenResult)
}

func (result *installationResult) checkInstallationProcess() {
	result.ResponseMessage = "Success"
	result.ResponseStatus = 200
}
