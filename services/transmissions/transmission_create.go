package transmissions

import (
	"strings"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/kusumoto/seedbox-daemon/services/network"
)

// CreateTransmissionContainer create a lasted transmission container and configuration from API endpoint
func CreateTransmissionContainer(cli *docker.Client, config *TransmissionConfig, netconfig *network.NetworkConfig) (*docker.Container, error) {
	// Setting container data binding
	torrentBindPath := []string{config.SrcTorrentPath, "/var/lib/transmission-daemon/downloads"}
	torrentSettingBindPath := []string{config.SrcConfigPath, "/var/lib/transmission-daemon/info"}
	torrentIncompleteBindPath := []string{config.SrcIncompleteTorrentPath, "/var/lib/transmission-daemon/incomplete"}
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
			Binds: []string{strings.Join(torrentBindPath, ":"), strings.Join(torrentSettingBindPath, ":"), strings.Join(torrentIncompleteBindPath, ":")},
		},
	}
	// Create Container
	return cli.CreateContainer(containerOptions)
}
