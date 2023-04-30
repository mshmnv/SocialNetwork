package metrics

import "syscall"

// DiskUsage contains usage data and provides user-friendly access methods
type DiskUsage struct {
	stat *syscall.Statfs_t
}

// NewDiskUsage returns an object holding the disk usage of volumePath ("/") or nil in case of error (invalid path, etc)
func NewDiskUsage() *DiskUsage {

	var stat syscall.Statfs_t
	err := syscall.Statfs("/", &stat) // /dev/sda
	if err != nil {
		return nil
	}
	return &DiskUsage{&stat}
}

// Free returns total free bytes on file system
func (du *DiskUsage) Free() uint64 {
	return du.stat.Bfree * uint64(du.stat.Bsize)
}

// Available return total available bytes on file system to an unprivileged user
func (du *DiskUsage) Available() uint64 {
	return du.stat.Bavail * uint64(du.stat.Bsize)
}

// Size returns total size of the file system
func (du *DiskUsage) Size() uint64 {
	return du.stat.Blocks * uint64(du.stat.Bsize)
}
