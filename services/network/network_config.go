package network

import docker "github.com/fsouza/go-dockerclient"

// NetworkConfig holds parameters to configuration network setting
type NetworkConfig struct {
	NetworkID    string `json:"network_id"`
	NetoworkName string `json:"network_name"`
	IPRange      string `json:"network_iprange"`
	SubnetMast   string `json:"network_subnetmast"`
	Gateway      string `json:"network_gateway"`
}

// Network interface for Network control implementation
type Network interface {
	CreateNetwork(cli *docker.Client, config *NetworkConfig) (*docker.Network, error)
	RemoveNetwork(cli *docker.Client, config *NetworkConfig) error
}
