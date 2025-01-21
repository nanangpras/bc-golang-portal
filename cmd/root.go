package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var rootCmd = &cobra.Command{
	Use: "core-api",
	Short: "this api for news portal",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Run(startCmd,nil)
	},
}

func Execute()  {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is .env)")

	rootCmd.Flags().BoolP("toggle", "t", false , "Help message for toggle")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}else{
		viper.SetConfigFile(`.env`)
	}

	viper.AutomaticEnv()

	// if err := viper.ReadInConfig(); err != nil {
	// 	fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	// }

	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading config file:", err)
		os.Exit(1)
	} else {
		fmt.Fprintln(os.Stdout, "Using config file:", viper.ConfigFileUsed())
		fmt.Println("Config file loaded successfully!")
	}
	
}