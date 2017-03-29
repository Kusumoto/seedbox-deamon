package transmissions

import docker "github.com/fsouza/go-dockerclient"

// RemoveTransmissionContainer remove a transmission container
func RemoveTransmissionContainer(cli *docker.Client, config *TransmissionConfig) error {
	// Setting container option
	var removeContainerOption = docker.RemoveContainerOptions{
		ID: config.ContainerID,
	}
	// Remove container
	return cli.RemoveContainer(removeContainerOption)
}
