package main

import (
	"github.com/alikhajeh1/goa_app1/client"
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"log"
	"os"
)

type (
	// ShowNumbersCommand is the command line data structure for the show action of numbers
	ShowNumbersCommand struct {
		// Prime numbers less than this number are calculated
		LessThan int
	}
)

// Run makes the HTTP request corresponding to the ShowNumbersCommand command.
func (cmd *ShowNumbersCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/numbers"
	}
	logger := goa.NewStdLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowNumbers(ctx, path, cmd.LessThan)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowNumbersCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var tmp2 int
	cc.Flags().IntVar(&cmd.LessThan, "lessThan", tmp2, `Prime numbers less than this number are calculated`)
}
