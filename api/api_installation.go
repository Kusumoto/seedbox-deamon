package api

import (
	"math/rand"
	"time"
)

type installationResult struct {
	responseStatus      int
	responseMessage     string
	responseAccessToken string
}

func (result *installationResult) InstallEndpoint() {
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
	result.responseAccessToken = string(tokenResult)
}

func (result *installationResult) checkInstallationProcess() {
}
