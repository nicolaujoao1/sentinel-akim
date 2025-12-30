package cli

import (
	"fmt"

	"github.com/nicolaujoao1/sentinel-akim/internal/collector"
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

		fmt.Printf("CPU Usage: %.2f%%\n", metrics.CPUPercent)
		fmt.Printf("Memory Usage: %.2f%%\n", metrics.MemoryPercent)
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
