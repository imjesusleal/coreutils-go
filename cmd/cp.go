package cmd

import (
	"fmt"
    "io"
    "os"
	"github.com/spf13/cobra"
)

var cpCmd = &cobra.Command{
    Use: "cp",
    Short: "copy a file to another directory",
    Example: "cp filename new/directory/filename",
    Args: cobra.MinimumNArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
        cpFunc(args) 
    },
}

func cpFunc(args ...[]string) (int64, error) {
    var dst string
    var src string
    for _,v := range args {
        dst = v[1]
        src = v[0]
    }
    srcFileStat, err := os.Stat(src)
    check(err)
    if !srcFileStat.Mode().IsRegular() {
        return 0, fmt.Errorf("%s is not a regular file", src)
    }
    source, err := os.Open(src)
    check(err)
    defer source.Close()

    destination, err := os.Create(dst)
    check(err)
    defer destination.Close()


    nBytes, err := io.Copy(destination, source)
    if err != nil {
        fmt.Println(err)
    }

    fmt.Printf("file %s copied to %s", src, dst)
    return nBytes, err
}
