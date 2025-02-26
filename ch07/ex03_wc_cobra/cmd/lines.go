/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// linesCmd represents the lines command
var linesCmd = &cobra.Command{
	Use:   "lines",
	Short: "count lines in file(s)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("lines called")
	},
}

func init() {
	rootCmd.AddCommand(linesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// linesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// linesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
