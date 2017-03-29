package api

import (
	"github.com/kusumoto/seedbox-daemon/services/disk"
	"github.com/kusumoto/seedbox-daemon/services/nginx"
	"github.com/kusumoto/seedbox-daemon/services/transmissions"
)

type createNewContainerResult struct {
	DiskConfig         disk.DiskConfig                  `json:"disk_config"`
	NginxConfig        nginx.NginxConfig                `json:"nginx_config"`
	TransmissionConfig transmissions.TransmissionConfig `json:"transmission_config"`
	AccessToken        string
	ResponseStatus     int    `json:"response_status"`
	ResponseMessage    string `json:"response_message"`
}

func (result *createNewContainerResult) createNewContainerEndpoint() {
	result.createVirtualDisk()
	result.createNetwork()
	result.createNginxContainer()
	result.createTransmissionsContainer()
}

func (result *createNewContainerResult) createNetwork() {
}

func (result *createNewContainerResult) createVirtualDisk() {
}

func (result *createNewContainerResult) createNginxContainer() {
}

func (result *createNewContainerResult) createTransmissionsContainer() {
}
