package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)
var (
	cfgFile 	string
	userLicense	string


	rootCmd = &cobra.Command{
		Use:	"tcr",
		Short:	"GoTCR is a Golang implementation of TEST && COMMIT || REVERT",
		Long:	"GoTCR is a Golang implementation of the TEST && COMMIT || REVERT workflow advocated by Kent Beck",
}

func Execute() error  {
	return rootCmd.Execute()
}

func init() {
	cobra.onInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/gotcr.yaml)")
	rootCmd.PersistentFlags().StringP("author", "a", "E.T. Anderson", "author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP($userLicense, "license", "l", "", "name of license for the project")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	viper.BindPFLag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "E.T. Anderson <anderson.eric.t@gmail.com>")
	viper.SetDefault("license", "TBD")

	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(initCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".gotcr")
	}
	
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.configFileUsed())
	}

}

