package api

import (
	"github.com/kusumoto/seedbox-daemon/services/disk"
	"github.com/kusumoto/seedbox-daemon/services/nginx"
	"github.com/kusumoto/seedbox-daemon/services/transmissions"
)

type createNewContainerResult struct {
	diskConfig         disk.DiskConfig
	nginxConfig        nginx.NginxConfig
	transmissionConfig transmissions.TransmissionConfig
}

func (result *createNewContainerResult) CreateNewContainerEndpoint() {
}

func (result *createNewContainerResult) createVirtualDisk() {
}

func (result *createNewContainerResult) createNginxContainer() {
}

func (result *createNewContainerResult) createTransmissionsContainer() {
}
