package transmissions

import docker "github.com/fsouza/go-dockerclient"

// StartTransmission start the transmission container
func StartTransmission(cli *docker.Client, config *TransmissionConfig) error {
	return cli.StartContainer(config.ContainerID, &docker.HostConfig{})
}

// StopTransmission stop the transmission container
func StopTransmission(cli *docker.Client, config *TransmissionConfig) error {
	return cli.StopContainer(config.ContainerID, 3000)
}
