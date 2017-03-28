package disk

import (
	"os/exec"
)

// CreateVirtualDisk create a virtual disk from specification
func CreateVirtualDisk(config *DiskConfig) error {
	err := exec.Command("dd", "if=/dev/zero", "of="+config.DiskLocation+"/"+config.DiskName+".img", "bs=512k", "count=200").Run()
	if err != nil {
		return err
	}
	err = exec.Command("mkfs.ext4", config.DiskLocation+"/"+config.DiskName+".img").Run()
	if err != nil {
		return err
	}
	return nil
}
