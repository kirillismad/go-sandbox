package main

import (
	"fmt"
	"log"
	"sandbox/petproj"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, World!")
		},
	}

	rootCmd.AddCommand(petproj.RunServer)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
