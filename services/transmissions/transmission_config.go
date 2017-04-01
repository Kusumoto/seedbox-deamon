package transmissions

import (
	docker "github.com/fsouza/go-dockerclient"
	"github.com/kusumoto/seedbox-daemon/services/network"
)

// TransmissionConfig holds parameters to configuration transmission container setting
type TransmissionConfig struct {
	ContainerID               string `json:"transmission_container_id"`
	ContainerName             string `json:"transmission_container_name"`
	ImageName                 string `json:"transmission_image_name"`
	SrcTorrentPath            string `json:"transmission_src_torrent_path"`
	SrcConfigPath             string `json:"transmission_src_config_path"`
	SrcIncompleteTorrentPath  string `json:"transmission_src_incomplete_torrent_path"`
	Username                  string `json:"transmission_username"`
	Password                  string `json:"transmission_password"`
	Port                      int    `json:"transmission_port"`
	LimitTorrentWorking       int    `json:"transmission_limit_torrent_working"`
	LimitTorrentSeed          int    `json:"transmission_limit_torrent_seed"`
	LimitTorrentUploadSpeed   int    `json:"transmission_limit_torrent_upload_seed"`
	LimitTorrentDownloadSpeed int    `json:"transmission_limit_torrent_download_seed"`
}

// Transmission interface for transmission control implementation
type Transmission interface {
	CreateTransmissionContainer(cli *docker.Client, config *TransmissionConfig, netconfig *network.NetworkConfig) (*docker.Container, error)
	PullTransmissionImage(cli *docker.Client) error
	RemoveTransmissionContainer(cli *docker.Client, config *TransmissionConfig) error
	StartTransmission(cli *docker.Client, config *TransmissionConfig) error
	StopTransmission(cli *docker.Client, config *TransmissionConfig) error
}

// TransmissionImageName holds constant define transmission image name
const TransmissionImageName = "dperson/transmission"

// TransmissionImageTag holds constant define transmission image tag
const TransmissionImageTag = "latest"
