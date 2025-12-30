package cli

import (
	"fmt"

	"github.com/nicolaujoao1/sentinel-akim/internal/analyzer"
	"github.com/nicolaujoao1/sentinel-akim/internal/collector"
	"github.com/nicolaujoao1/sentinel-akim/internal/scorer"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show system status",
	Run: func(cmd *cobra.Command, args []string) {

		metrics, err := collector.CollectSystemMetrics()
		if err != nil {
			fmt.Println("Error collecting system metrics:", err)
			return
		}

		health := scorer.CalculateHealth(metrics.CPUPercent, metrics.MemoryPercent)

		fmt.Printf("CPU Usage: %.2f%%\n", metrics.CPUPercent)
		fmt.Printf("Memory Usage: %.2f%%\n\n", metrics.MemoryPercent)

		fmt.Printf("Health Score: %d/100\n", health.Score)
		fmt.Printf("Status: %s\n", health.Status)

		diagnostics := analyzer.AnalyzeSystem(metrics.CPUPercent, metrics.MemoryPercent)

		fmt.Println("\nDiagnostics:")
		for _, d := range diagnostics {
			fmt.Printf("[%s] %s\n", d.Severity, d.Title)
			fmt.Printf("Cause: %s\n", d.Cause)

			if d.Recommendation != "" {
				fmt.Printf("Recommendation: %s\n", d.Recommendation)
			}
			fmt.Println()
		}

	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
