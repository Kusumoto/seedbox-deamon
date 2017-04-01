package models

import "fmt"

// ContainerModel is model for container query result
type ContainerModel struct {
	IsError       bool
	ErrorMessage  string
	ID            int
	ContainerID   string
	ContainerName string
	ImageName     string
	Author        string
	NetworkID     string
}

// InsertContainerData is insert container data
func (model *ContainerModel) InsertContainerData() error {
	db := InitDatabase()
	stmt, err := db.Prepare("INSERT into containers (container_id, container_name, image_name, network_id, author) VALUES (?,?,?,?,?)")
	if err != nil {
		fmt.Println(err)
		model = ContainerModel{
			IsError:      true,
			ErrorMessage: err.Error(),
		}
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(model.ContainerID, model.ContainerName, model.ImageName, model.NetworkID, model.Author)
	if err != nil {
		fmt.Println(err)
		models = ContainerModel{
			IsError:      true,
			ErrorMessage: err.Error(),
		}
		return
	}
	defer db.Close()
}
