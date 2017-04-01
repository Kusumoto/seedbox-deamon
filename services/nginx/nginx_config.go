package nginx

import (
	docker "github.com/fsouza/go-dockerclient"
	"github.com/kusumoto/seedbox-daemon/services/network"
)

// NginxConfig holds parameters to configuration nginx container setting
type NginxConfig struct {
	ContainerID     string `json:"nginx_container_id"`
	ContainerName   string `json:"nginx_container_name"`
	ImageName       string `json:"nginx_image_name"`
	SrcTorrentPath  string `json:"nginx_src_torrent_path"`
	DestTorrentPath string `json:"nginx_dest_torrent_path"`
	NginxConfig     string `json:"nginx_config"`
	Port            string `json:"nginx_port"`
	LimitConnection int    `json:"nginx_limit_connection"`
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
