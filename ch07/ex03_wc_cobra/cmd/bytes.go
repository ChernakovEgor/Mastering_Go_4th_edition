/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// bytesCmd represents the bytes command
var bytesCmd = &cobra.Command{
	Use:   "bytes",
	Short: "count bytes in file(s)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bytes called")
	},
}

func init() {
	rootCmd.AddCommand(bytesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bytesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bytesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
