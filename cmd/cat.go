package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
)


type cat struct{
    value *cobra.Command    
}

var newCat = cat {
    value: &cobra.Command {
        Use: "cat", 
        Short: "concat some file",
        Args: cobra.MinimumNArgs(1),
        Run: func(cmd *cobra.Command, args []string) {
            if len(args) == 1 {
                funCat(args)
            }
            if len(args) == 3 {
                redCat(args)
            }
        },
    },
}


func funCat(args []string) {
    if len(args) == 1 {
        cont, err := os.ReadFile(args[0])
        if err != nil {
            panic(err)
        }
        fmt.Printf("%s", cont)
    }

}

func redCat(args []string) {
    fmt.Println(len(args))
    if len(args) == 3 && args[1] == ">" {
        temp := strings.Split(args[3], "/")
        fileIdx := len(temp) -1
        file := temp[fileIdx]
        tempSlice := temp[:fileIdx]
        dir := strings.Join(tempSlice, "/") 
        _, err := os.Stat(dir)
        if err == nil {return }
        if os.IsNotExist(err) {
            createFile(file, dir, args[0])
        }
    }
}

func createFile(file string, dir string, info string) {
    os.Chdir(dir)
    c, err := os.Create(file)
    if err != nil {
        panic(err)
    }
    res, err := c.WriteString(info)
    if err != nil {
        panic(err)
    }
    fmt.Println(res)
}


