package models

import (
	"os"
)

// CheckInstallationStatus is check status the installation for this daemon
func CheckInstallationStatus() bool {
	if _, err := os.Stat(defaultDatabaseFile); os.IsExist(err) {
		return false
	}
	installationStatus := CheckInstallation()
	if !installationStatus.IsAlreadyInstalled {
		return false
	}
	return true
}

// CheckAccessToken is check a access token in request
func CheckAccessToken(accessToken string) bool {
	installationStatus := CheckInstallation()
	if installationStatus.AccessToken != accessToken {
		return false
	}
	return true
}
