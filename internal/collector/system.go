package collector

import (
	"time"

	"github.com/nicolaujoao1/sentinel-akim/internal/models"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

func CollectSystemMetrics() (*models.SystemMetrics, error) {
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return nil, err
	}

	memStat, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	return &models.SystemMetrics{
		CPUPercent:    cpuPercent[0],
		MemoryPercent: memStat.UsedPercent,
	}, nil
}
