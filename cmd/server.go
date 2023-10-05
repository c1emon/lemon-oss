/*
Copyright © 2023 clemon
*/
package cmd

import (
	"fmt"

	"github.com/c1emon/lemon_oss/internal/setting"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server called")
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	cfg := setting.GetCfg()
	serverCmd.PersistentFlags().IntVarP(&cfg.Http.Port, "port", "p", 8080, "server port")
	viper.BindPFlag("port", serverCmd.PersistentFlags().Lookup("port"))

	serverCmd.PersistentFlags().StringVar(&cfg.DB.Driver, "driver", "postgres", "db driver name")
	viper.BindPFlag("driver", serverCmd.PersistentFlags().Lookup("driver"))

	serverCmd.PersistentFlags().StringVar(&cfg.DB.Source, "source", "host=10.10.0.70 port=5432 user=postgres dbname=lemon_tree password=123456", "db source")
	viper.BindPFlag("source", serverCmd.PersistentFlags().Lookup("source"))
}
