package api

import (
	"github.com/kusumoto/seedbox-daemon/services/disk"
	"github.com/kusumoto/seedbox-daemon/services/nginx"
	"github.com/kusumoto/seedbox-daemon/services/transmissions"
)

type createNewContainerResult struct {
	DiskConfig         disk.DiskConfig                  `json:"DiskConfig"`
	NginxConfig        nginx.NginxConfig                `json:"NginxConfig"`
	TransmissionConfig transmissions.TransmissionConfig `json:"TransmissionConfig"`
	AccessToken        string
	ResponseStatus     int    `json:"responseStatus"`
	ResponseMessage    string `json:"responseMessage"`
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
