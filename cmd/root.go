package main

import (
	"errors"
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "miaosha",
	Short: "miaosha",
	Long:  `简洁高效的秒杀系统设计 https://miaosha.sopans.com`,
	Args:  args,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func args(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("至少需要一个参数!")
	}
	return nil
}
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
	}
}
func init() {
	rootCmd.AddCommand(serverCmd)
}
