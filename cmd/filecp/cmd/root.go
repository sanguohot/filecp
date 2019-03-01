package cmd

import (
	"github.com/sanguohot/filecp/etc"
	"github.com/sanguohot/filecp/pkg/common/log"
	"github.com/sanguohot/filecp/pkg/filecp"
	"github.com/spf13/cobra"
)

var (
	cfgFile string
	src string
	dst string
	csv string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "filecp",
	Short: "use to copy files.",
	Long: `a command tool to copy project's file to other diretory with the same relative path.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		cp := filecp.New(src, dst, csv)
		if err := cp.Copy(); err != nil {
			log.Logger.Fatal(err.Error())
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Logger.Fatal(err.Error())
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "","config file")
	rootCmd.PersistentFlags().StringVarP(&csv, "from-csv", "f", "", "the csv file that contain src and dst file")
	rootCmd.PersistentFlags().StringVarP(&src, "src", "s", "", "the source file to copy")
	rootCmd.PersistentFlags().StringVarP(&dst, "dst", "d", "", "the destination file path to copy")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	etc.InitConfig(cfgFile)
}
