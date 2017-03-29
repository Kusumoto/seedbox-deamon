package api

import docker "github.com/fsouza/go-dockerclient"

func InstallationResponse() {
	client, err := docker.NewClient(defaultDockerEndpoint)
}
