package network

import (
	"context"

	"github.com/docker/docker/client"
)

// RemoveNetwork remove a docker network
func RemoveNetwork(ctx context.Context, cli *client.Client, config *NetworkConfig) error {
	return cli.NetworkRemove(ctx, config.NetworkID)
}
