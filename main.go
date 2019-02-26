package main

import (
	"fmt"
	"log"
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
			addr := cmd.Flags().Lookup("addr").Value.String()
			dir := cmd.Flags().Lookup("dir").Value.String()
			fs := http.FileServer(http.Dir(dir))
			http.Handle("/", fs)
			log.Println("Listening on " + addr)
			http.ListenAndServe(addr, nil)
		},
	}
	cmd.Flags().StringP("dir", "r", ".", "Data Directory")
	cmd.Flags().StringP("addr", "a", ":8080", "API listening address")
	viper.BindPFlag("dir", cmd.Flags().Lookup("dir"))
	viper.BindPFlag("addr", cmd.Flags().Lookup("addr"))
	return cmd
}
