package transmissions

type TransmissionConfig struct {
	ContainerID 				string
	ContainerName 				string
	ImageName 					string
	TorrentPath 				string
	ConfigPath					string
	Port 						int
	LimitTorrentWorking 		int
	LimitTorrentSeed 			int
	LimitTorrentUploadSpeed 	int
	LimitTorrentDownloadSpeed 	int
}