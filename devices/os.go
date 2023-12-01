package devices

import (
	"fmt"
	"syscall"
)

var (
	uname *syscall.Utsname
	os    OperatingSystem
)

type OperatingSystem struct {
	Name          string
	Kernel        string
	KernelVersion string
	KernelRelease string
	Machine       string
}

func NewOperatingSystem() (*OperatingSystem, error) {
	uname = new(syscall.Utsname)
	os = OperatingSystem{}
	if err := syscall.Uname(uname); err != nil {
		return nil, fmt.Errorf("error : %w", err)
	}
	os.initialize(*uname)
	return &os, nil
}

func (o *OperatingSystem) initialize(osinfo syscall.Utsname) {
	o.Name = toHumanReadable(osinfo.Nodename[:])
	o.Kernel = toHumanReadable(osinfo.Sysname[:])
	o.KernelRelease = toHumanReadable(osinfo.Release[:])
	o.KernelVersion = toHumanReadable(osinfo.Version[:])
	o.Machine = toHumanReadable(osinfo.Machine[:])
}

func toHumanReadable(data []int8) string {
	var readable string
	for _, v := range data {
		if v == 0 {
			break
		}
		readable += string(byte(v))
	}
	return readable
}
