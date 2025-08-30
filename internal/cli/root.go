package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	verbose bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "proton",
	Short: "Generate documentation for Go libraries",
	Long: `Proton is a documentation generator for Go libraries that creates 
GitBook-compatible documentation from your Go source code, comments, and templates.

It supports:
- Single and multi-package libraries
- Automatic API documentation generation
- Configurable templates and output
- GitBook integration with .gitbook.yml generation
- GitHub Actions support`,
	Version: "1.0.0",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// It initializes the CLI application and runs the root command.
// Returns error if command execution fails.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	// Global flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is .proton/config.yml)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	// Bind flags to viper
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search for config in common locations
		viper.AddConfigPath(".proton")
		viper.AddConfigPath(".config")
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	// Read in environment variables that match
	viper.SetEnvPrefix("PROTON")
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil && verbose {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
