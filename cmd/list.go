/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/tchoffman/tri/todo"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List Todos",
	Long:  `List all Todo Items`,
	Run:   listRun,
}

func listRun(cmd *cobra.Command, args []string) {

	items, err := todo.ReadItems(dataFile)
	if err != nil {
		log.Printf("%v", err)
	}

	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	for _, i := range items {
		fmt.Fprintln(w, strconv.Itoa(i.Priority)+"\t"+i.Text+"\t")
	}
	w.Flush()
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
