package cmd

import (
	"fmt"
	"os"

	"github.com/misbahulard/jenkins-tracer/config"
	"github.com/misbahulard/jenkins-tracer/tracer"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "jenkins-tracer",
	Short: "Jenkins tracer is used to record all the jenkins job variables.",
	Long: `Jenkins tracer is used to record all the jenkins job variables. It's like record the build duration, build
variables, repository metadata, etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Jenkins Tracer v%s\n", Version)
		initialize()
		tracer.Run()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func initialize() {
	fmt.Println("Initialize tracer..")
	config.New()
	config.ConfigureLogger()
	config.ConfigureElasticsearch()
}
