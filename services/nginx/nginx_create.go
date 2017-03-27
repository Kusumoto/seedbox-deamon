package nginx

import (
	"context"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	seedboxNet "github.com/kusumoto/seedbox-daemon/services/network"
)

// CreateNignxContainer create a lasted nginx container and configuration from API endpoint
func CreateNignxContainer(cxt context.Context, cli *client.Client, config *NginxConfig, netconfig *seedboxNet.NetworkConfig) error {
	containerHostConfig := container.HostConfig{}
	hostConfig := container.Config{}
	networkConfig := network.NetworkingConfig{
		EndpointsConfig: map["torrent_net"]network.EndpointSettings{
			Gateway:   netconfig.Gateway,
			NetworkID: netconfig.NetworkID,
		},
	}
}
