package metrics

import (
	"github.com/shirou/gopsutil/v3/process"
)

type ProcessMetrics struct {
	CPU     float64
	Memory  float32
	IORead  uint64
	IOWrite uint64
}

func GetProcessMetrics(pid int32) (*ProcessMetrics, error) {
	p, err := process.NewProcess(pid)
	if err != nil {
		return nil, err
	}

	cpuPercent, err := p.CPUPercent()
	if err != nil {
		return nil, err
	}

	memPercent, err := p.MemoryPercent()
	if err != nil {
		return nil, err
	}

	ioCounters, err := p.IOCounters()
	if err != nil {
		return nil, err
	}

	return &ProcessMetrics{
		CPU:     cpuPercent,
		Memory:  memPercent,
		IORead:  ioCounters.ReadBytes,
		IOWrite: ioCounters.WriteBytes,
	}, nil
}
