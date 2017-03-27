package nginx

// NginxConfig holds parameters to configuration nginx container setting
type NginxConfig struct {
	ContainerID     string
	ContainerName   string
	ImageName       string
	TorrentPath     string
	Port            int
	LimitConnection int
}
