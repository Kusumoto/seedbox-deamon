package models

import (
	"fmt"
)

// InstallationModel is model for installation data
type InstallationModel struct {
	IsError            bool
	IsAlreadyInstalled bool
	AccessToken        string
	ErrorMessage       string
}

// InsertAccessToken is insert installation token
func InsertAccessToken(accessToken string) *InstallationModel {
	models := InstallationModel{
		IsError:            false,
		IsAlreadyInstalled: false,
		AccessToken:        accessToken,
		ErrorMessage:       "",
	}
	db := InitDatabase()
	stmt, err := db.Prepare("INSERT into installation (is_already_installed, access_token) VALUES (?,?)")
	if err != nil {
		fmt.Println(err)
		models = InstallationModel{
			IsError:            true,
			IsAlreadyInstalled: false,
			AccessToken:        "",
			ErrorMessage:       err.Error(),
		}
		return &models
	}
	defer stmt.Close()
	_, err = stmt.Exec(models.IsAlreadyInstalled, models.AccessToken)
	if err != nil {
		fmt.Println(err)
		models = InstallationModel{
			IsError:            true,
			IsAlreadyInstalled: false,
			AccessToken:        "",
			ErrorMessage:       err.Error(),
		}
		return &models
	}
	defer db.Close()
	return &models
}

// UpdateInstallation is update status installation of this daemon
func UpdateInstallation(accessToken string) *InstallationModel {
	models := InstallationModel{
		IsError:            false,
		IsAlreadyInstalled: true,
		AccessToken:        accessToken,
		ErrorMessage:       "",
	}
	db := InitDatabase()
	stmt, err := db.Prepare("UPDATE installation SET is_already_installed = ? WHERE access_token = ?")
	if err != nil {
		fmt.Println(err)
		models = InstallationModel{
			IsError:            true,
			IsAlreadyInstalled: false,
			AccessToken:        "",
			ErrorMessage:       err.Error(),
		}
		return &models
	}
	defer stmt.Close()
	_, err = stmt.Exec(models.IsAlreadyInstalled, models.AccessToken)
	if err != nil {
		fmt.Println(err)
		models = InstallationModel{
			IsError:            true,
			IsAlreadyInstalled: false,
			AccessToken:        "",
			ErrorMessage:       err.Error(),
		}
		return &models
	}
	defer db.Close()
	return &models
}

// CheckInstallation is check daemon installation
func CheckInstallation() *InstallationModel {
	models := InstallationModel{
		IsError:            false,
		IsAlreadyInstalled: false,
		AccessToken:        "",
		ErrorMessage:       "",
	}
	db := InitDatabase()
	rows, err := db.Query("SELECT * from installation")
	if err != nil {
		fmt.Println(err)
		models = InstallationModel{
			IsError:            true,
			IsAlreadyInstalled: false,
			AccessToken:        "",
			ErrorMessage:       err.Error(),
		}
		return &models
	}
	for rows.Next() {
		var is_already_installed bool
		var access_token string
		err = rows.Scan(&is_already_installed, &access_token)
		if err != nil {
			fmt.Println(err)
			models = InstallationModel{
				IsError:            true,
				IsAlreadyInstalled: false,
				AccessToken:        "",
				ErrorMessage:       err.Error(),
			}
			return &models
		}
		models = InstallationModel{
			IsError:            false,
			IsAlreadyInstalled: is_already_installed,
			AccessToken:        access_token,
			ErrorMessage:       "",
		}
	}
	defer db.Close()
	return &models
}
