package transmissions

import (
	docker "github.com/fsouza/go-dockerclient"
)

// PullTransmissionImage pull a lasted transmission image from repository
func PullTransmissionImage(cli *docker.Client) error {
	// Setting image option
	var imageOption = docker.PullImageOptions{
		Registry: TransmissionImageName,
		Tag:      TransmissionImageTag,
	}
	// Pull image
	return cli.PullImage(imageOption, docker.AuthConfiguration{})
}
