package transmissions

import (
	"strings"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/kusumoto/seedbox-daemon/services/network"
)

// CreateNignxContainer create a lasted nginx container and configuration from API endpoint
func CreateTransmissionContainer(cli *docker.Client, config *TransmissionConfig, netconfig *network.NetworkConfig) (*docker.Container, error) {
	// Setting container data binding
	torrentBindPath := []string{config.SrcTorrentPath, config.DestTorrentPath}
	nginxSiteEnableConfigBindPath := []string{config.NginxConfig, "/etc/nginx/sites-enabled"}
	nginxSiteAvailableConfigBindPath := []string{config.NginxConfig, "/etc/nginx/sites-available"}
	networkEndpointSetting := make(map[string]*docker.EndpointConfig)
	// Setting container network
	networkEndpointSetting["torrent_net"] = &docker.EndpointConfig{
		Gateway:   netconfig.Gateway,
		NetworkID: netconfig.NetworkID,
	}
	// Setting container data
	containerOptions := docker.CreateContainerOptions{
		Config: &docker.Config{},
		NetworkingConfig: &docker.NetworkingConfig{
			EndpointsConfig: networkEndpointSetting,
		},
		HostConfig: &docker.HostConfig{
			Binds: []string{strings.Join(torrentBindPath, ":"), strings.Join(nginxSiteEnableConfigBindPath, ":"), strings.Join(nginxSiteAvailableConfigBindPath, ":")},
		},
	}
	// Create Container
	return cli.CreateContainer(containerOptions)
}
