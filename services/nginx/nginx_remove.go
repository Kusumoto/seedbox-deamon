package nginx

import docker "github.com/fsouza/go-dockerclient"

// RemoveNginxContainer remove a nginx container
func RemoveNginxContainer(cli *docker.Client, config *NginxConfig) error {
	// Setting container option
	var removeContainerOption = docker.RemoveContainerOptions{
		ID: config.ContainerID,
	}
	// Remove container
	return cli.RemoveContainer(removeContainerOption)
}
