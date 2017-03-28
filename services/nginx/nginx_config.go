package nginx

import (
	docker "github.com/fsouza/go-dockerclient"
	"github.com/kusumoto/seedbox-daemon/services/network"
)

// NginxConfig holds parameters to configuration nginx container setting
type NginxConfig struct {
	ContainerID     string
	ContainerName   string
	ImageName       string
	SrcTorrentPath  string
	DestTorrentPath string
	NginxConfig     string
	Port            string
	LimitConnection int
}

// Nginx interface for nginx control implementation
type Nginx interface {
	CreateNignxContainer(cli *docker.Client, config *NginxConfig, netconfig *network.NetworkConfig) (*docker.Container, error)
	PullNginxImage(cli *docker.Client) error
	RemoveNginxContainer(cli *docker.Client, config *NginxConfig) error
	StartNignx(cli *docker.Client, config *NginxConfig) error
	StopNginx(cli *docker.Client, config *NginxConfig) error
}

// NginxImageName holds constant define nginx image name
const NginxImageName = "nginx"

// NginxImageTag holds constant define nginx image tag
const NginxImageTag = "stable"
