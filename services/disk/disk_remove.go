package disk

import (
	"os/exec"
)

// RemoveVirtualDisk remove a virtual disk from specification
func RemoveVirtualDisk(config *DiskConfig) error {
	err := exec.Command("rm", "-rf", config.DiskLocation+"/"+config.DiskName).Run()
	if err != nil {
		return err
	}
	return nil
}
