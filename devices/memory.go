package devices

import (
	"errors"
	"fmt"
	"syscall"
)

var (
	syscallMemory syscall.Sysinfo_t
	memory        Memory
	errorMemory   = errors.New("cannot load memory information")
)

type Ram struct {
	Total, Free, Usage string
}

type Swap struct {
	Total, Free, Usage string
}

type Memory struct {
	Ram
	Swap
}

func (m *Memory) initialize(mem syscall.Sysinfo_t) {
	ramUsage := syscallMemory.Totalram - syscallMemory.Freeram
	swapUsage := syscallMemory.Totalswap - syscallMemory.Freeswap

	m.Ram.Total = m.toReadableSize(mem.Totalram)
	m.Ram.Free = m.toReadableSize(mem.Freeram)
	m.Ram.Usage = m.toReadableSize(ramUsage)
	m.Swap.Total = m.toReadableSize(mem.Totalswap)
	m.Swap.Free = m.toReadableSize(mem.Freeswap)
	m.Swap.Usage = m.toReadableSize(swapUsage)
}

func (m *Memory) toReadableSize(size uint64) string {
	unitName := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	convertedSize := float64(size)

	i := 0
	for convertedSize >= 1024 && i < len(unitName)-1 {
		convertedSize /= 1024
		i++
	}

	formatted := fmt.Sprintf("%.2f %s", convertedSize, unitName[i])
	return formatted
}

func NewMemory() (*Memory, error) {
	err := syscall.Sysinfo(&syscallMemory)
	if err != nil {
		return nil, errorMemory
	}
	memory = Memory{}
	memory.initialize(syscallMemory)
	return &memory, nil
}
