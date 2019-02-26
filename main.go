package main

import (
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo"

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
			ext := cmd.Flags().Lookup("ext").Value.String()
			e := echo.New()
			e.GET("/:file", func(c echo.Context) error {
				return c.File(dir + "/" + c.Param("file") + "." + ext)
			})
			log.Println("Listening on " + addr)
			e.Start(addr)
		},
	}
	cmd.Flags().StringP("dir", "r", ".", "Data Directory")
	cmd.Flags().StringP("ext", "e", "json", "File extension to serve")
	cmd.Flags().StringP("addr", "a", ":8080", "API listening address")
	viper.BindPFlag("dir", cmd.Flags().Lookup("dir"))
	viper.BindPFlag("ext", cmd.Flags().Lookup("ext"))
	viper.BindPFlag("addr", cmd.Flags().Lookup("addr"))
	return cmd
}
