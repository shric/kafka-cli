package cmd

import (
	"github.com/shric/kafka-cli/pkg/kafka-cli"
	"github.com/spf13/cobra"
)

func init() {
	consumerCmd.PersistentFlags().String("group", "my-consumer-group", "consumer group")
	consumerCmd.PersistentFlags().String("topic", "my-topic", "topic")
	RootCmd.AddCommand(consumerCmd)
}

const (
	flagGroup = "group"
	flagTopic = "topic"
)

var consumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "Display information about a consumer/topic",
	RunE: func(cmd *cobra.Command, args[] string) error {
		var err error
		flags := cmd.Flags()
		c := kafka_cli.ConsumerCmd{}
		c.Group, err = flags.GetString(flagGroup)
		if err != nil {
			return err
		}
		c.Topic, err = flags.GetString(flagTopic)
		if err != nil {
			return err
		}
		return c.Run(cmd.OutOrStdout())
	},
}

