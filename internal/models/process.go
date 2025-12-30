package models

type ProcessInfo struct {
	PID        int32
	Name       string
	CPUPercent float64
	MemoryMB   float64
}
