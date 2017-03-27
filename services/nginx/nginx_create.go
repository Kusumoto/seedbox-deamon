package nginx

import (
	"context"
	"strings"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	seedboxNet "github.com/kusumoto/seedbox-daemon/services/network"
)

// CreateNignxContainer create a lasted nginx container and configuration from API endpoint
func CreateNignxContainer(ctx context.Context, cli *client.Client, config *NginxConfig, netconfig *seedboxNet.NetworkConfig) (container.ContainerCreateCreatedBody, error) {
	// Setting container data binding
	torrentBindPath := []string{config.SrcTorrentPath, config.DestTorrentPath}
	nginxSiteEnableConfigBindPath := []string{config.NginxConfig, "/etc/nginx/sites-enabled"}
	nginxSiteAvailableConfigBindPath := []string{config.NginxConfig, "/etc/nginx/sites-available"}
	containerHostConfig := container.HostConfig{
		PortBindings: nat.PortMapping{Port: "80", Binding: nat.PortBinding{HostPort: config.Port}},
		Binds:        []string{strings.Join(torrentBindPath, ":"), strings.Join(nginxSiteEnableConfigBindPath, ":"), strings.Join(nginxSiteAvailableConfigBindPath, ":")},
	}
	containerConfig := container.Config{
		Image: NginxImageName,
	}
	// Setting container network
	networkEndpointSetting := make(map[string]*network.EndpointSettings)
	networkEndpointSetting["torrent_net"] = &network.EndpointSettings{
		Gateway:   netconfig.Gateway,
		NetworkID: netconfig.NetworkID,
	}
	networkConfig := network.NetworkingConfig{
		EndpointsConfig: networkEndpointSetting,
	}
	return cli.ContainerCreate(ctx, &containerConfig, &containerHostConfig, &networkConfig, config.ContainerName)
}
