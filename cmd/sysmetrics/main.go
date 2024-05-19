package main

import (
	"log"
	"time"

	"metrics/internal/config"
	"metrics/internal/metrics"
)

func main() {
	cfg := config.LoadConfig()

	for {
		// Сбор метрик
		cpuUsage, err := metrics.GetCPUUsage()
		if err != nil {
			log.Printf("Ошибка получения метрик CPU: %v", err)
		} else {
			log.Printf("CPU Usage: %.2f%%", cpuUsage)
		}

		memoryUsage, err := metrics.GetMemoryUsage()
		if err != nil {
			log.Printf("Ошибка получения метрик памяти: %v", err)
		} else {
			log.Printf("Memory Usage: %.2f%%", memoryUsage)
		}

		diskUsage, err := metrics.GetDiskUsage()
		if err != nil {
			log.Printf("Ошибка получения метрик диска: %v", err)
		} else {
			log.Printf("Disk Usage: %.2f%%", diskUsage)
		}

		processMetrics, err := metrics.GetProcessMetrics(cfg.PID)
		if err != nil {
			log.Printf("Ошибка получения метрик процесса: %v", err)
		} else {
			log.Printf("Process CPU Usage: %.2f%%", processMetrics.CPU)
			log.Printf("Process Memory Usage: %.2f%%", processMetrics.Memory)
			log.Printf("Process IO Read: %d bytes", processMetrics.IORead)
			log.Printf("Process IO Write: %d bytes", processMetrics.IOWrite)
		}

		// Пауза перед следующим сбором метрик
		time.Sleep(10 * time.Second)
	}
}
