package network

import docker "github.com/fsouza/go-dockerclient"

// RemoveNetwork remove a docker network
func RemoveNetwork(cli *docker.Client, config *NetworkConfig) error {
	return cli.RemoveNetwork(config.NetoworkName)
}
