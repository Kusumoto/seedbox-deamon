package api

import (
	docker "github.com/fsouza/go-dockerclient"
	"github.com/kusumoto/seedbox-daemon/services/disk"
	"github.com/kusumoto/seedbox-daemon/services/network"
	"github.com/kusumoto/seedbox-daemon/services/nginx"
	"github.com/kusumoto/seedbox-daemon/services/transmissions"
)

type createNewContainerResult struct {
	DiskConfig         disk.DiskConfig                  `json:"disk_config"`
	NginxConfig        nginx.NginxConfig                `json:"nginx_config"`
	TransmissionConfig transmissions.TransmissionConfig `json:"transmission_config"`
	NetworkConfig      network.NetworkConfig            `json:"network_config"`
	AccessToken        string
	ResponseStatus     int    `json:"response_status"`
	ResponseMessage    string `json:"response_message"`
}

func (result *createNewContainerResult) createNewContainerEndpoint(cli docker.Client) {
	result.createVirtualDisk(cli)
	result.createNetwork(cli)
	result.createNginxContainer(cli)
	result.createTransmissionsContainer(cli)
}

func (result *createNewContainerResult) createNetwork(cli docker.Client) {
	networkResult, err := network.CreateNetwork(&cli, &result.NetworkConfig)
	result.NetworkConfig.NetworkID = networkResult.ID
}

func (result *createNewContainerResult) createVirtualDisk(cli docker.Client) {
	err := disk.CreateVirtualDisk(&result.DiskConfig)
}

func (result *createNewContainerResult) createNginxContainer(cli docker.Client) {
	nginxResult, err := nginx.CreateNignxContainer(&cli, &result.NginxConfig, &result.NetworkConfig)
	result.NginxConfig.ContainerID = nginxResult.ID
}

func (result *createNewContainerResult) createTransmissionsContainer(cli docker.Client) {
	transmissionsResult, err := transmissions.CreateTransmissionContainer(&cli, &result.TransmissionConfig, &result.NetworkConfig)
	result.TransmissionConfig.ContainerID = transmissionsResult.ID
}
