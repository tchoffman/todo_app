/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var dataFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo_app",
	Short: "A Command Line Tool to Manage TODO Items",
	Long:  `Exploring creation of Go projects and Cobra-CLI projects in particular.\nThis is just an exercise to explore some of the basics of Go.`,
	Run: func(cmd *cobra.Command, args []string) {
		t, _ := cmd.Flags().GetBool("toggle")
		fmt.Println("Value of Toggle: ", t)
		m, _ := cmd.Flags().GetString("message")
		fmt.Println("Message: ", m)
		status, _ := cmd.PersistentFlags().GetString("status")
		fmt.Println("Status: ", status)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	home, err := homedir.Dir()
	if err != nil {
		log.Println("Unable to detect home directory. Please set data file using --datafile.")
	}
	rootCmd.PersistentFlags().StringVar(
		&dataFile,
		"datafile",
		home+string(os.PathSeparator)+".tridos.json",
		"data file to store Todos")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringP("message", "m", "", "Message for the TODO item")
}
