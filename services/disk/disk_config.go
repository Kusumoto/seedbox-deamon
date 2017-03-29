package disk

// DiskConfig holds parameters to configuration disk setting
type DiskConfig struct {
	DiskLocation string
	DiskName     string
	DiskSize     int
}

// Disk interface for disk control implementation
type Disk interface {
	CreateVirtualDisk(config *DiskConfig) error
	RemoveVirtualDisk(config *DiskConfig) error
}
