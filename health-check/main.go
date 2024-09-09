package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/process"
)

type SystemHealthInPercent struct {
	SystemHealth  float64
	ServiceHealth float64
	OverallHealth float64
}

type Thresholds struct {
	Low            float64
	Medium         float64
	High           float64
	Critical       float64
	GoroutinesLow  int
	GoroutinesHigh int
}

func getSystemCPUUsage() (float64, error) {
	percentages, err := cpu.Percent(0, false)
	if err != nil {
		return 0, err
	}
	return percentages[0], nil
}

func getSystemMemoryUsage() float64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	totalMemory := float64(m.Sys) / (1024 * 1024)
	usedMemory := float64(m.Alloc) / (1024 * 1024)
	return (usedMemory / totalMemory) * 100
}

func getGoroutineCount() int {
	return runtime.NumGoroutine()
}

func getProcessCPUUsage(pid int32) (float64, error) {
	proc, err := process.NewProcess(pid)
	if err != nil {
		return 0, err
	}
	percent, err := proc.CPUPercent()
	if err != nil {
		return 0, err
	}
	return percent, nil
}

func getProcessMemoryUsage(pid int32) (float64, error) {
	proc, err := process.NewProcess(pid)
	if err != nil {
		return 0, err
	}
	memInfo, err := proc.MemoryInfo()
	if err != nil {
		return 0, err
	}
	totalMemory := float64(memInfo.RSS) / (1024 * 1024)
	return totalMemory, nil
}

func getHealthScore(usage float64, thresholds Thresholds) int {
	switch {
	case usage <= thresholds.Low:
		return 10
	case usage <= thresholds.Medium:
		return 7
	case usage <= thresholds.High:
		return 4
	default:
		return 1
	}
}

func getGoroutineHealthScore(goroutines int, thresholds Thresholds) int {
	switch {
	case goroutines <= thresholds.GoroutinesLow:
		return 10
	case goroutines <= thresholds.GoroutinesHigh:
		return 7
	default:
		return 4
	}
}

func calculateServiceHealth(pid int32, thresholds Thresholds) (float64, error) {
	cpuUsage, err := getProcessCPUUsage(pid)
	if err != nil {
		return 0, err
	}

	memoryUsage, err := getProcessMemoryUsage(pid)
	if err != nil {
		return 0, err
	}

	cpuScore := getHealthScore(cpuUsage, thresholds)
	memoryScore := getHealthScore(memoryUsage, thresholds)

	healthScore := float64(cpuScore+memoryScore) / 2
	healthScore = healthScore * 10

	return healthScore, nil
}

func calculateSystemHealth(pid int32, thresholds Thresholds) (float64, error) {
	systemCPUUsage, err := getSystemCPUUsage()
	if err != nil {
		return 0, err
	}

	systemMemoryUsage := getSystemMemoryUsage()

	serviceCPUUsage, err := getProcessCPUUsage(pid)
	if err != nil {
		return 0, err
	}

	serviceMemoryUsage, err := getProcessMemoryUsage(pid)
	if err != nil {
		return 0, err
	}

	adjustedCPUUsage := systemCPUUsage - serviceCPUUsage
	adjustedMemoryUsage := systemMemoryUsage - (serviceMemoryUsage * 100 / (1024 * 1024))

	if adjustedCPUUsage < 0 {
		adjustedCPUUsage = 0
	}
	if adjustedMemoryUsage < 0 {
		adjustedMemoryUsage = 0
	}

	goroutines := getGoroutineCount()

	cpuScore := getHealthScore(adjustedCPUUsage, thresholds)
	memoryScore := getHealthScore(adjustedMemoryUsage, thresholds)
	goroScore := getGoroutineHealthScore(goroutines, thresholds)

	healthScore := float64(cpuScore+memoryScore+goroScore) / 3
	healthScore = healthScore * 10

	return healthScore, nil
}

func main() {
	thresholds := Thresholds{
		Low:            20.0,
		Medium:         50.0,
		High:           80.0,
		Critical:       100.0,
		GoroutinesLow:  100,
		GoroutinesHigh: 500,
	}

	pid := int32(os.Getpid())
	for {
		systemHealth, err := calculateSystemHealth(pid, thresholds)
		if err != nil {
			fmt.Printf("Error calculating system health: %v\n", err)
			return
		}

		serviceHealth, err := calculateServiceHealth(pid, thresholds)
		if err != nil {
			fmt.Printf("Error calculating service health: %v\n", err)
			return
		}

		overallHealth := (systemHealth + serviceHealth) / 2

		log.Println("System Health:", systemHealth, "% out of 100%")
		log.Println("Service Health:", serviceHealth, "% out of 100%")
		log.Println("Overall Health Score:", overallHealth, "% out of 100%")

		time.Sleep(10 * time.Second)
	}
}
