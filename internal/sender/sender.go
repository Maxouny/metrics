package sender

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type Metrics struct {
	CPUUsage       float64 `json:"cpu_usage"`
	MemoryUsage    float64 `json:"memory_usage"`
	DiskUsage      float64 `json:"disk_usage"`
	ProcessCPU     float64 `json:"process_cpu_usage"`
	ProcessMemory  float32 `json:"process_memory_usage"`
	ProcessIORead  uint64  `json:"process_io_read"`
	ProcessIOWrite uint64  `json:"process_io_write"`
}

func SendMetrics(metrics Metrics, url string) {
	jsonData, err := json.Marshal(metrics)
	if err != nil {
		log.Printf("Ошибка кодирования метрик: %v", err)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Ошибка отправки метрик: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Сервер вернул ошибку: %v", resp.Status)
	}
}
