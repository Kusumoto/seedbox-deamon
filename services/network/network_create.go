package network

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

// CreateNetwork create a docker network for encapsulate a container
func CreateNetwork(ctx context.Context, cli *client.Client, config *NetworkConfig) (types.NetworkCreateResponse, error) {
	ipmConfig := []network.IPAMConfig{
		network.IPAMConfig{
			Subnet:  config.SubnetMast,
			IPRange: config.IPRange,
			Gateway: config.Gateway,
		},
	}
	networkOption := types.NetworkCreate{
		CheckDuplicate: true,
		Driver:         "bridge",
		Attachable:     true,
		IPAM: &network.IPAM{
			Config: ipmConfig,
		},
	}
	return cli.NetworkCreate(ctx, config.NetoworkName, networkOption)
}
