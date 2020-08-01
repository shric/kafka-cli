package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)


var rootCmd = &cobra.Command{
	Use:   "kafka",
	Short: "kafka-cli is a command line interface to Kafka",
	Long: `A pure Go kafka CLI using segmentio/kafka-go. See https://github.com/shric/kafka-go`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

