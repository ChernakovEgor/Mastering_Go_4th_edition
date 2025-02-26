/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// wordsCmd represents the words command
var wordsCmd = &cobra.Command{
	Use:   "words",
	Short: "count words in file(s)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("words called")
	},
}

func init() {
	rootCmd.AddCommand(wordsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// wordsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// wordsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
