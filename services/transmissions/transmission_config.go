package transmissions

import (
	docker "github.com/fsouza/go-dockerclient"
	"github.com/kusumoto/seedbox-daemon/services/network"
)

type TransmissionConfig struct {
	ContainerID               string
	ContainerName             string
	ImageName                 string
	SrcTorrentPath            string
	DestTorrentPath           string
	ConfigPath                string
	Port                      int
	LimitTorrentWorking       int
	LimitTorrentSeed          int
	LimitTorrentUploadSpeed   int
	LimitTorrentDownloadSpeed int
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
const TransmissionImageName = "nginx"

// TransmissionImageTag holds constant define transmission image tag
const TransmissionImageTag = "stable"
