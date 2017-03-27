package services

type DockerAPIClient interface {
	DockerContainerClient
	DockerImageClient
	ClientVersion() string
}

type DockerContainerClient interface {
	ContainerCreate()
	ContainerRemove()
	ContainerLog()
	ContainerUpdate()
	ContainerRun()
}

type DockerImageClient interface {
	ImagePull()
	ImageRemove()
}

type DiskAPI interface {
	DiskCreate()
	DiskRemove()
}
