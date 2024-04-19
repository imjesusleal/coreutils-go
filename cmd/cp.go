package cmd

import (
	"fmt"
//	"io/fs"
//	"os"
	"github.com/spf13/cobra"
)

var cpCmd = &cobra.Command{
    Use: "cp",
    Short: "copy a file to another directory",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("hi copy")
    },
}

