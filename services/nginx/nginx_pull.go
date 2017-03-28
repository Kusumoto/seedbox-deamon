package nginx

import (
	"context"

	docker "github.com/fsouza/go-dockerclient"
)

// PullNginxImage pull a lasted nginx image from repository
func PullNginxImage(ctx context.Context, cli *docker.Client) error {
	// Setting image option
	var imageOption = docker.PullImageOptions{
		Registry: NginxImageName,
		Tag:      NginxImageTag,
	}
	// Pull image
	return cli.PullImage(imageOption, docker.AuthConfiguration{})
}
