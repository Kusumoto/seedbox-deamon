package api

import docker "github.com/fsouza/go-dockerclient"

// NewDockerClient is create a new docker instant
func NewDockerClient() (docker.Client, error) {
	client, err := docker.NewClient(defaultDockerEndpoint)
	return *client, err
}
