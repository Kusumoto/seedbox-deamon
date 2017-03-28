package network

import docker "github.com/fsouza/go-dockerclient"

// NetworkConfig holds parameters to configuration network setting
type NetworkConfig struct {
	NetworkID    string
	NetoworkName string
	IPRange      string
	SubnetMast   string
	Gateway      string
}

// Network interface for Network control implementation
type Network interface {
	CreateNetwork(cli *docker.Client, config *NetworkConfig) (*docker.Network, error)
	RemoveNetwork(cli *docker.Client, config *NetworkConfig) error
}
