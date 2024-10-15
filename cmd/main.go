package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/svanhalla/base-rest-server/internal"
)

var logFormat string

var (
	name      = "base-rest-server"
	version   = "dev"
	buildDate = "undefined"
	user      = "unknown"
)

// rootCmd representerar basen för alla kommandon
var rootCmd = &cobra.Command{
	Use:   "base-rest-server",
	Short: "A base-rest-server CLI",
	Long:  `This is a CLI application to manage and start the server.`,
}

// startCmd representerar start-kommandot
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Starting %s...\n", name)
		server := internal.NewServer(cmd.Flag("log-format").Value.String())
		log.Fatal(server.Start(":8080"))
	},
}

// versionCmd representerar version-kommandot
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Name: %s\nVersion: %s\nBuilt by: %s\nBuild date: %s\n", name, version, user, buildDate)
	},
}

func init() {
	// Lägg till start- och version-kommandon till rootCmd
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.PersistentFlags().StringVar(&logFormat, "log-format", "text", "Set log format (text|json)")
}

// Execute körs från main och exekverar rootCmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
