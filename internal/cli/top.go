package cli

import (
	"fmt"

	"github.com/nicolaujoao1/sentinel-akim/internal/collector"
	"github.com/spf13/cobra"
)

var topCmd = &cobra.Command{
	Use:   "top",
	Short: "Show top processes",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use --cpu or --mem")
	},
}

var topCPU bool
var topMem bool

var topRunCmd = &cobra.Command{
	Use:    "run",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {

		processes, err := collector.CollectProcesses()
		if err != nil {
			fmt.Println("Error collecting processes:", err)
			return
		}

		if topMem {
			collector.SortByMemory(processes)
			fmt.Println("TOP PROCESSES (Memory)")
		} else if topCPU {
			collector.SortByCPU(processes)
			fmt.Println("TOP PROCESSES (CPU)")
		} else {
			fmt.Println("Use --cpu or --mem")
			return
		}

		fmt.Printf("%-8s %-20s %-10s\n", "PID", "NAME", "VALUE")
		for i, p := range processes {
			if i >= 5 {
				break
			}

			if topMem {
				fmt.Printf("%-8d %-20s %.2f MB\n", p.PID, p.Name, p.MemoryMB)
			} else {
				fmt.Printf("%-8d %-20s %.2f %%\n", p.PID, p.Name, p.CPUPercent)
			}
		}
	},
}

func init() {
	topCmd.Flags().BoolVar(&topCPU, "cpu", false, "Sort by CPU usage")
	topCmd.Flags().BoolVar(&topMem, "mem", false, "Sort by memory usage")

	topCmd.Run = topRunCmd.Run
	rootCmd.AddCommand(topCmd)
}
