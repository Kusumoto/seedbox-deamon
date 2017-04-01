package disk

// DiskConfig holds parameters to configuration disk setting
type DiskConfig struct {
	DiskLocation string `json:"disk_location"`
	DiskName     string `json:"disk_name"`
	DiskSize     int    `json:"disk_size"`
}

// Disk interface for disk control implementation
type Disk interface {
	CreateVirtualDisk(config *DiskConfig) error
	RemoveVirtualDisk(config *DiskConfig) error
}
