/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
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
		fmt.Fprintln(w, i.PrettyP()+"\t"+i.Text+"\t")
	}
	w.Flush()
}

func init() {
	rootCmd.AddCommand(listCmd)
}
