/*
Copyright © 2023 clemon
*/
package cmd

import (
	"context"

	"github.com/c1emon/gcommon/gormx"
	"github.com/c1emon/gcommon/logx"
	"github.com/c1emon/lemon_oss/internal/server"
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
		defer func() {
			if err := gormx.DisConnect(); err != nil {
				logx.GetLogger().Warnf("unable close db: %s", err)
			}
		}()

		cfg := setting.GetCfg()

		gormx.Initialize(cfg.DB.Driver, cfg.DB.Source)
		s, _ := server.Initialize(cfg)
		go s.ListenToSystemSignals(context.Background())
		s.Run()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.PersistentFlags().IntP("port", "p", 8080, "server port")
	viper.BindPFlag("server.port", serverCmd.PersistentFlags().Lookup("port"))

	serverCmd.PersistentFlags().String("driver", "postgres", "db driver name")
	viper.BindPFlag("db.driver", serverCmd.PersistentFlags().Lookup("driver"))

	serverCmd.PersistentFlags().String("source", "host=localhost port=5432 user=postgres dbname=lemon_oss password=123456", "db source")
	viper.BindPFlag("db.source", serverCmd.PersistentFlags().Lookup("source"))
}
