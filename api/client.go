package api

import docker "github.com/fsouza/go-dockerclient"

func NewDockerClient() (docker.Client, error) {
	client, err := docker.NewClient(defaultDockerEndpoint)
	return *client, err
}
