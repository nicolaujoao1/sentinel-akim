package scorer

import "github.com/nicolaujoao1/sentinel-akim/internal/models"

func CalculateHealth(cpuPercent, memPercent float64) models.HealthReport {
	score := 100

	if cpuPercent > 80 {
		score -= 30
	} else if cpuPercent > 60 {
		score -= 20
	} else if cpuPercent > 40 {
		score -= 10
	}

	if memPercent > 80 {
		score -= 30
	} else if memPercent > 60 {
		score -= 20
	} else if memPercent > 40 {
		score -= 10
	}

	status := "OK"
	if score < 70 {
		status = "WARNING"
	}
	if score < 40 {
		status = "CRITICAL"
	}

	return models.HealthReport{
		Score:  score,
		Status: status,
	}
}
