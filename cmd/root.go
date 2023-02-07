package cmd

import (
	"os"

	"github.com/HarrisChu/compare_query/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	sampleFile  string
	sourceFile  string
	nebulaAddr  string
	nebulaSpace string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "compare_query",
	Short: "a tool to compare query results",
	Long:  ``,
}

var sampleCmd = &cobra.Command{
	Use:   "sample",
	Short: "sample source file",
	RunE: func(cmd *cobra.Command, args []string) error {
		c := pkg.NewController(sourceFile, sampleFile, 1000, 200, nebulaAddr, nebulaSpace)
		err := c.Sample()
		if err != nil {
			return err
		}
		return nil
	},
}

var compareCmd = &cobra.Command{
	Use:   "compare",
	Short: "compare different queries",
	RunE: func(cmd *cobra.Command, args []string) error {
		c := pkg.NewController(sourceFile, sampleFile, 1000, 200, nebulaAddr, nebulaSpace)
		err := c.Compare()
		if err != nil {
			return err
		}
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	sampleFlags := pflag.NewFlagSet("", pflag.ContinueOnError)
	sampleFlags.StringVar(&sourceFile, "source", "", "source file")
	sampleFlags.StringVar(&sampleFile, "sample", "", "sample file")
	sampleCmd.PersistentFlags().AddFlagSet(sampleFlags)

	compareFlags := pflag.NewFlagSet("", pflag.ContinueOnError)
	compareFlags.StringVar(&sampleFile, "sample", "", "sample file")
	compareFlags.StringVar(&nebulaAddr, "nebulagraph", "192.168.15.8:9669", "sample file")
	compareFlags.StringVar(&nebulaSpace, "space", "sf100", "sample file")
	compareCmd.PersistentFlags().AddFlagSet(compareFlags)

	rootCmd.AddCommand(sampleCmd)
	rootCmd.AddCommand(compareCmd)

}
