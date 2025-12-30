package collector

import (
	"sort"

	"github.com/nicolaujoao1/sentinel-akim/internal/models"
	"github.com/shirou/gopsutil/v3/process"
)

func CollectProcesses() ([]models.ProcessInfo, error) {
	procs, err := process.Processes()
	if err != nil {
		return nil, err
	}

	var result []models.ProcessInfo

	for _, p := range procs {
		name, err := p.Name()
		if err != nil {
			continue
		}

		cpu, _ := p.CPUPercent()
		mem, _ := p.MemoryInfo()

		if mem == nil {
			continue
		}

		result = append(result, models.ProcessInfo{
			PID:        p.Pid,
			Name:       name,
			CPUPercent: cpu,
			MemoryMB:   float64(mem.RSS) / 1024 / 1024,
		})
	}

	return result, nil
}

func SortByMemory(processes []models.ProcessInfo) {
	sort.Slice(processes, func(i, j int) bool {
		return processes[i].MemoryMB > processes[j].MemoryMB
	})
}

func SortByCPU(processes []models.ProcessInfo) {
	sort.Slice(processes, func(i, j int) bool {
		return processes[i].CPUPercent > processes[j].CPUPercent
	})
}
