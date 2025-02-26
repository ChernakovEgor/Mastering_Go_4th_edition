/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// charsCmd represents the chars command
var charsCmd = &cobra.Command{
	Use:   "chars",
	Short: "count chars in file(s)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("chars called")
	},
}

func init() {
	rootCmd.AddCommand(charsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// charsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// charsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
