package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
    rootCmd = &cobra.Command{
    Use: "",
    Short: "cobra is a simple CLI to replicate commands.",
})

func Execute() error {
    return rootCmd.Execute()
}

func init() {
    var longListing bool
    var targetPath bool 

    rootCmd.PersistentFlags().StringP("author", "a", "Jesus Leal", "Created by Jesus Leal.")
    viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
    lsCmd.PersistentFlags().BoolVarP(&longListing, "long", "l", false, "print long format.")
    mkdirCmd.PersistentFlags().BoolVarP(&targetPath, "directory", "d", false, "create a directory on the target path")

    rootCmd.AddCommand(newCat.value)
    rootCmd.AddCommand(lsCmd)
    rootCmd.AddCommand(mkdirCmd)
}


