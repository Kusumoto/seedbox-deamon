package nginx

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

// NginxImageName holds constant define nginx image name
const NginxImageName = "nginx:stable"
