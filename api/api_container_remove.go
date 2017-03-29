package api

import (
	"github.com/kusumoto/seedbox-daemon/services/disk"
	"github.com/kusumoto/seedbox-daemon/services/nginx"
	"github.com/kusumoto/seedbox-daemon/services/transmissions"
)

type removeContainerResult struct {
	diskConfig         disk.DiskConfig
	nginxConfig        nginx.NginxConfig
	transmissionConfig transmissions.TransmissionConfig
	accessToken        string
	responseStatus     int
	responseMessage    string
}

func (result *removeContainerResult) removeContainerEndpoint() {
	result.removeNginxContainer()
	result.removeTransmissionsContainer()
	result.removeNetwork()
	result.removeVirtualDisk()
}

func (result *removeContainerResult) removeNetwork() {

}

func (result *removeContainerResult) removeNginxContainer() {

}

func (result *removeContainerResult) removeTransmissionsContainer() {

}

func (result *removeContainerResult) removeVirtualDisk() {

}
