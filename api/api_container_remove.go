package api

import (
	docker "github.com/fsouza/go-dockerclient"
	"github.com/kusumoto/seedbox-daemon/api/models"
	"github.com/kusumoto/seedbox-daemon/services/disk"
	"github.com/kusumoto/seedbox-daemon/services/network"
	"github.com/kusumoto/seedbox-daemon/services/nginx"
	"github.com/kusumoto/seedbox-daemon/services/transmissions"
	iris "gopkg.in/kataras/iris.v6"
)

type removeContainerResult struct {
	DiskConfig         disk.DiskConfig                  `json:"disk_config"`
	NginxConfig        nginx.NginxConfig                `json:"nginx_config"`
	TransmissionConfig transmissions.TransmissionConfig `json:"transmission_config"`
	NetworkConfig      network.NetworkConfig            `json:"network_config"`
	AccessToken        string                           `json:"access_token"`
	ResponseStatus     int                              `json:"response_status"`
	ResponseMessage    string                           `json:"response_message"`
}

func (result *removeContainerResult) removeContainerEndpoint(cli docker.Client) {
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
	result.removeNginxContainer(cli)
	result.removeTransmissionsContainer(cli)
	result.removeNetwork(cli)
	result.removeVirtualDisk(cli)
}

func (result *removeContainerResult) removeNetwork(cli docker.Client) {
	if result.ResponseStatus != iris.StatusOK {
		return
	}
	err := network.RemoveNetwork(&cli, &result.NetworkConfig)
	if err != nil {
		result.ResponseStatus = iris.StatusInternalServerError
		result.ResponseMessage = err.Error()
		return
	}
}

func (result *removeContainerResult) removeNginxContainer(cli docker.Client) {
	if result.ResponseStatus != iris.StatusOK {
		return
	}
	err := nginx.RemoveNginxContainer(&cli, &result.NginxConfig)
	if err != nil {
		result.ResponseStatus = iris.StatusInternalServerError
		result.ResponseMessage = err.Error()
		return
	}
}

func (result *removeContainerResult) removeTransmissionsContainer(cli docker.Client) {
	if result.ResponseStatus != iris.StatusOK {
		return
	}
	err := transmissions.RemoveTransmissionContainer(&cli, &result.TransmissionConfig)
	if err != nil {
		result.ResponseStatus = iris.StatusInternalServerError
		result.ResponseMessage = err.Error()
		return
	}
}

func (result *removeContainerResult) removeVirtualDisk(cli docker.Client) {
	if result.ResponseStatus != iris.StatusOK {
		return
	}
	err := disk.RemoveVirtualDisk(&result.DiskConfig)
	if err != nil {
		result.ResponseStatus = iris.StatusInternalServerError
		result.ResponseMessage = err.Error()
		return
	}
}
