package nginx

import docker "github.com/fsouza/go-dockerclient"

// StartNignx start the nginx container
func StartNignx(cli *docker.Client, config *NginxConfig) error {
	return cli.StartContainer(config.ContainerID, &docker.HostConfig{})
}

// StopNginx stop the nginx container
func StopNginx(cli *docker.Client, config *NginxConfig) error {
	return cli.StopContainer(config.ContainerID, 3000)
}
