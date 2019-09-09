package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/xafardero/gitpr/github"
)

var t string
var b string

func init() {
	createCmd.Flags().StringVarP(&t, "title", "t", "", "This is the title of the PR")
	createCmd.Flags().StringVarP(&b, "body", "b", "", "This is the body of the PR")

	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a PR in github",
	Long:  `Creates a PR in github's`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("Command required the head parameter")
		}

		if len(args) == 1 {
			return errors.New("Command required the base parameter")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("calling the github API")
		cpd := github.CreatePRDTO{
			Title: t,
			Head:  args[0],
			Base:  args[1],
			Body:  b,
			Owner: "owner",
			Repo:  "repository",
		}
		github.CreatePR(cpd)
	},
}
