package api

import (
	"math/rand"
	"time"
)

type installationResult struct {
	responseStatus      int    `json:"responseStatus"`
	responseMessage     string `json:"responseMessage"`
	responseAccessToken string `json:"responseAccessToken"`
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
	result.responseMessage = "Success"
	result.responseStatus = 200
	result.responseAccessToken = string(tokenResult)
}

func (result *installationResult) checkInstallationProcess() {
	result.responseMessage = "Success"
	result.responseStatus = 200
}
