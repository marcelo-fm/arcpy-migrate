/*
Copyright © 2024 Marcelo Mesquita
*/
package cmd

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"

	"github.com/elewis787/boa"
	gen "github.com/marcelo-fm/arcpy-migrate/internal/arcpy-gen"
	"github.com/spf13/cobra"
)

var (
	err error
	// cfgFile   string
	filenames []string
	files     []fs.DirEntry
	dir       string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "arcpy-migrate",
	Short: "Parse SQL Files into arcpy commands",
	Long: `
  Arcpy Migrate will parse any SQL files passed into the
  Stdin  and will generate arcpy python commands to
  run and create tables through the ArcGIS environment.

  Usage:
    cat <some_file>.sql | arcpy-migrate

  Examples:
    create a migration script file:
      cat <some_file>.sql | arcpy-migrate > migration_$(date "+%Y%m%d%H%M").py
  `,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		// text, err := io.ReadAll(reader)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		res, err := gen.Generate(reader)

		cobra.CheckErr(err)
		fmt.Println(string(res))
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.SetUsageFunc(boa.UsageFunc)
	rootCmd.SetHelpFunc(boa.HelpFunc)
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.arcpy-migrate.yaml)")
}

func initConfig() {
	// } else {
	// 	home, err := os.UserHomeDir()
	// 	cobra.CheckErr(err)
	//
	// 	viper.AddConfigPath(home)
	// 	viper.SetConfigType("yaml")
	// 	viper.SetConfigName(".arcpy-migrate")
	// }

	// viper.AutomaticEnv()
	//
	// if err = viper.ReadInConfig(); err == nil {
	// 	fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	// }
}
