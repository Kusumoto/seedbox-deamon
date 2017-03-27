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
	
	// Setter network configuration
	networkEndpointSetting := make(map[string]*network.EndpointSettings)
	networkEndpointSetting["torrent_net"] = network.EndpointSettings{
		Gateway:   netconfig.Gateway,
		NetworkID: netconfig.NetworkID,
	}
	networkConfig := network.NetworkingConfig{
		EndpointsConfig: networkEndpointSetting,
	},
}
