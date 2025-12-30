package analyzer

import "github.com/nicolaujoao1/sentinel-akim/internal/models"

func AnalyzeSystem(cpuPercent, memPercent float64) []models.Diagnostic {
	var diagnostics []models.Diagnostic

	if cpuPercent > 80 {
		diagnostics = append(diagnostics, models.Diagnostic{
			Severity:       "WARNING",
			Title:          "High CPU usage",
			Cause:          "CPU usage above 80%",
			Recommendation: "Check running processes or scale system resources",
		})
	}

	if memPercent > 80 {
		diagnostics = append(diagnostics, models.Diagnostic{
			Severity:       "WARNING",
			Title:          "High memory usage",
			Cause:          "Memory usage above 80%",
			Recommendation: "Review running processes or increase available RAM",
		})
	}

	if len(diagnostics) == 0 {
		diagnostics = append(diagnostics, models.Diagnostic{
			Severity: "OK",
			Title:    "System operating normally",
			Cause:    "CPU and memory within acceptable limits",
		})
	}

	return diagnostics
}
