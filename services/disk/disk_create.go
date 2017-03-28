package disk

import (
	"os/exec"
	"strconv"
)

// CreateVirtualDisk create a virtual disk from specification
func CreateVirtualDisk(config *DiskConfig) error {
	totalCount := config.DiskSize / 512
	err := exec.Command("dd", "if=/dev/zero", "of="+config.DiskLocation+"/"+config.DiskName+".img", "bs=512k", "count="+strconv.Itoa(totalCount)).Run()
	if err != nil {
		return err
	}
	err = exec.Command("mkfs.ext4", config.DiskLocation+"/"+config.DiskName+".img").Run()
	if err != nil {
		return err
	}
	return nil
}
