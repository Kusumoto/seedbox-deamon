package network

import (
	docker "github.com/fsouza/go-dockerclient"
)

// CreateNetwork create a docker network for encapsulate a container
func CreateNetwork(cli *docker.Client, config *NetworkConfig) (*docker.Network, error) {
	var netoworkOptions = docker.CreateNetworkOptions{
		CheckDuplicate: true,
		Driver:         "bridge",
		Name:           config.NetoworkName,
		IPAM: docker.IPAMOptions{
			Config: []docker.IPAMConfig{docker.IPAMConfig{
				Subnet:  config.SubnetMast,
				IPRange: config.IPRange,
				Gateway: config.Gateway,
			},
			},
		},
	}
	return cli.CreateNetwork(netoworkOptions)
}
