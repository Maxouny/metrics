package metrics

import "github.com/shirou/gopsutil/cpu"

func GetCPUUsage() (float64, error) {
	cpuPercent, err := cpu.Percent(0, false)
	if err != nil {
		return 0, err
	}
	return cpuPercent[0], nil
}
