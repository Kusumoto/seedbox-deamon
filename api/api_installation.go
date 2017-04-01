package api

import (
	"math/rand"
	"time"

	"github.com/kusumoto/seedbox-daemon/api/models"

	iris "gopkg.in/kataras/iris.v6"
)

// installationResult is model for installation result
type installationResult struct {
	ResponseStatus      int    `json:"response_status"`
	ResponseMessage     string `json:"response_message"`
	ResponseAccessToken string `json:"response_access_token"`
}

// installationRequestData is model for installation request
type installationRequestData struct {
	AccessToken string `json:"access_token"`
}

func (result *installationResult) installEndpoint() {
	result.checkInstallationProcess()
	result.generateAccessToken()
}

func (result *installationResult) installCallbackEndpoint(accessToken *installationRequestData) {
	result.checkInstallationProcess()
	result.updateInstallStatus(accessToken.AccessToken)
}

func (result *installationResult) removeEndpoint(accessToken *installationRequestData) {
	result.checkInstallationProcess()
	result.removeDatabaseFile(accessToken.AccessToken)
}

func (result *installationResult) removeDatabaseFile(accessToken string) {
	if accessToken == "" || accessToken != result.ResponseAccessToken {
		result.ResponseStatus = iris.StatusForbidden
		result.ResponseAccessToken = ""
		result.ResponseMessage = "access denied"
		return
	}
	models.DeleteDatabaseFile()
	result.ResponseStatus = iris.StatusOK
	result.ResponseMessage = "OK"
	result.ResponseAccessToken = ""
}

func (result *installationResult) updateInstallStatus(accessToken string) {
	if accessToken == "" || accessToken != result.ResponseAccessToken || result.ResponseStatus != 200 {
		result.ResponseStatus = iris.StatusForbidden
		result.ResponseAccessToken = ""
		result.ResponseMessage = "access denied"
		return
	}
	updateInstallationResult := models.UpdateInstallation(accessToken)
	if updateInstallationResult.IsError {
		result.ResponseStatus = iris.StatusInternalServerError
		result.ResponseAccessToken = ""
		result.ResponseMessage = updateInstallationResult.ErrorMessage
	} else {
		result.ResponseStatus = iris.StatusOK
		result.ResponseAccessToken = ""
		result.ResponseMessage = "OK"
	}
}

func (result *installationResult) generateAccessToken() {
	if result.ResponseStatus != iris.StatusOK || result.ResponseAccessToken != "" {
		return
	}
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	tokenResult := make([]byte, defaultAccessTokenLength)
	for i := 0; i < defaultAccessTokenLength; i++ {
		tokenResult[i] = chars[rand.Intn(len(chars))]
	}
	saveTokenResult := models.InsertAccessToken(string(tokenResult))
	if saveTokenResult.IsError {
		result.ResponseStatus = iris.StatusInternalServerError
		result.ResponseMessage = saveTokenResult.ErrorMessage
	} else {
		result.ResponseStatus = iris.StatusOK
		result.ResponseAccessToken = saveTokenResult.AccessToken
	}
}

func (result *installationResult) checkInstallationProcess() {
	installResult := models.CheckInstallation()
	if installResult.IsError {
		result.ResponseStatus = iris.StatusInternalServerError
		result.ResponseMessage = installResult.ErrorMessage
	} else if installResult.IsAlreadyInstalled {
		result.ResponseStatus = iris.StatusForbidden
		result.ResponseMessage = "this daemon has already installed."
	} else {
		result.ResponseStatus = iris.StatusOK
		result.ResponseMessage = "OK"
		result.ResponseAccessToken = installResult.AccessToken
	}
}
