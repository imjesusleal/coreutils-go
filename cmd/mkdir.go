package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)
var mkdirCmd = &cobra.Command {
    Use: "mkdir",
    Short: "create a directory",
    Example: "cobra mkdir [directory-name] [flags]",
    Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        targetPath,_ := cmd.Flags().GetBool("directory")
        if targetPath == false {
            mkdirFunc(args)        
        } else {
            mkdirFuncToTarget(args)
        }
    },

}

func mkdirFunc(argv ...[]string) {
    if len(argv) > 0  {
        var val string 
        for _,v := range argv {
            val = v[0] 
        }
        err := os.Mkdir(val, 0755)
        check(err)
        fmt.Println("directory created succesfully!")
    } 
} 

func mkdirFuncToTarget(argv ...[]string) {
    var name string
    var target string
    for _,v := range argv {
        name = v[0]
        target = v[1]
    }
    os.Chdir(target)
    err := os.Mkdir(name, 0755)
    check(err)
}
