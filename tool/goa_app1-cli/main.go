package main

import (
	"fmt"
	"github.com/alikhajeh1/goa_app1/client"
	"github.com/alikhajeh1/goa_app1/tool/cli"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"time"
)

func main() {
	// Create command line parser
	app := &cobra.Command{
		Use:   "goa_app1-cli",
		Short: `CLI client for the goa_app1 service`,
	}

	// Create client struct
	c := client.New(newHTTPClient())

	// Register global flags
	app.PersistentFlags().StringVarP(&c.Scheme, "scheme", "s", "", "Set the requests scheme")
	app.PersistentFlags().StringVarP(&c.Host, "host", "H", "localhost:8080", "API hostname")
	app.PersistentFlags().DurationVarP(&c.Timeout, "timeout", "t", time.Duration(20)*time.Second, "Set the request timeout")
	app.PersistentFlags().BoolVar(&c.Dump, "dump", false, "Dump HTTP request and response.")

	// Initialize API client
	c.UserAgent = "goa_app1-cli/0"

	// Register API commands
	cli.RegisterCommands(app, c)

	// Execute!
	if err := app.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(-1)
	}
}

// newHTTPClient returns the HTTP client used by the API client to make requests to the service.
func newHTTPClient() *http.Client {
	// TBD: Change as needed (e.g. to use a different transport to control redirection policy or
	// disable cert validation or...)
	return http.DefaultClient
}
