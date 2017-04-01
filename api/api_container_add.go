package api

import (
	docker "github.com/fsouza/go-dockerclient"
	"github.com/kusumoto/seedbox-daemon/api/models"
	"github.com/kusumoto/seedbox-daemon/services/disk"
	"github.com/kusumoto/seedbox-daemon/services/network"
	"github.com/kusumoto/seedbox-daemon/services/nginx"
	"github.com/kusumoto/seedbox-daemon/services/transmissions"
	"gopkg.in/kataras/iris.v6"
)

// createNewContainerResult is model for response container creation result
type createNewContainerResult struct {
	DiskConfig         disk.DiskConfig                  `json:"disk_config"`
	NginxConfig        nginx.NginxConfig                `json:"nginx_config"`
	TransmissionConfig transmissions.TransmissionConfig `json:"transmission_config"`
	NetworkConfig      network.NetworkConfig            `json:"network_config"`
	AccessToken        string                           `json:"access_token"`
	ResponseStatus     int                              `json:"response_status"`
	ResponseMessage    string                           `json:"response_message"`
}

func (result *createNewContainerResult) createNewContainerEndpoint(cli docker.Client) {
	result.ResponseStatus = iris.StatusOK
	if !models.CheckInstallationStatus() {
		result.ResponseStatus = iris.StatusForbidden
		result.ResponseMessage = "this daemon has been not install"
		return
	}
	if !models.CheckAccessToken(result.AccessToken) {
		result.ResponseStatus = iris.StatusForbidden
		result.ResponseMessage = "access denied"
		return
	}
	result.createVirtualDisk(cli)
	result.createNetwork(cli)
	result.createNginxContainer(cli)
	result.createTransmissionsContainer(cli)
}

func (result *createNewContainerResult) createNetwork(cli docker.Client) {
	if result.ResponseStatus != iris.StatusOK {
		return
	}
	networkResult, err := network.CreateNetwork(&cli, &result.NetworkConfig)
	if err != nil {
		result.ResponseStatus = iris.StatusInternalServerError
		result.ResponseMessage = err.Error()
		return
	}
	result.ResponseStatus = iris.StatusOK
	result.NetworkConfig.NetworkID = networkResult.ID
}

func (result *createNewContainerResult) createVirtualDisk(cli docker.Client) {
	if result.ResponseStatus != iris.StatusOK {
		return
	}
	err := disk.CreateVirtualDisk(&result.DiskConfig)
	if err != nil {
		result.ResponseStatus = iris.StatusInternalServerError
		result.ResponseMessage = err.Error()
		return
	}
	result.ResponseStatus = iris.StatusOK
}

func (result *createNewContainerResult) createNginxContainer(cli docker.Client) {
	if result.ResponseStatus != iris.StatusOK {
		return
	}
	nginxResult, err := nginx.CreateNignxContainer(&cli, &result.NginxConfig, &result.NetworkConfig)
	if err != nil {
		result.ResponseStatus = iris.StatusInternalServerError
		result.ResponseMessage = err.Error()
		return
	}
	result.ResponseStatus = iris.StatusOK
	result.NginxConfig.ContainerID = nginxResult.ID
}

func (result *createNewContainerResult) createTransmissionsContainer(cli docker.Client) {
	if result.ResponseStatus != iris.StatusOK {
		return
	}
	transmissionsResult, err := transmissions.CreateTransmissionContainer(&cli, &result.TransmissionConfig, &result.NetworkConfig)
	if err != nil {
		result.ResponseStatus = iris.StatusInternalServerError
		result.ResponseMessage = err.Error()
		return
	}
	result.ResponseStatus = iris.StatusOK
	result.TransmissionConfig.ContainerID = transmissionsResult.ID
}
