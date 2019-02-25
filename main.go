package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	serverCmd := serverCmd()
	if err := serverCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func serverCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Simple Static Server",
		Run: func(cmd *cobra.Command, args []string) {
			flags := cmd.Flags()
			fs := http.FileServer(http.Dir(flags.Lookup("data").Value.String()))
			http.Handle("/", fs)
			http.ListenAndServe(flags.Lookup("addr").Value.String(), nil)
		},
	}
	cmd.Flags().StringP("data", "r", ".", "Data folder")
	cmd.Flags().StringP("addr", "a", ":8080", "API listening address")
	viper.BindPFlag("data", cmd.Flags().Lookup("data"))
	viper.BindPFlag("addr", cmd.Flags().Lookup("addr"))
	return cmd
}
