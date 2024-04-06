package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"github.com/spf13/cobra"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

var lsCmd = &cobra.Command{
    Use: "ls",
    Short: "listing of directory", 
    Example: "cobra ls",
    Run: func(cmd *cobra.Command, args []string) {
        lsFlag,_ := cmd.Flags().GetBool("long")
        if lsFlag ==  false {
            switch len(args) {
            case 0:
                res := lsFunc()
                for _,v := range res {
                    fmt.Println(v)
                }
            default:
                res := lsFunc(args)
                for _,v := range res {
                    fmt.Println(v)
                }
            } 
        } else {
            res := longListing(args)
            for _,v := range res {
                fmt.Println(v)
            } 
        }
    },
}


func lsFunc(path ...[]string) (res []fs.DirEntry) {
    if path == nil || path[0][0] == "." { 
        cwd, err := os.Getwd()
        check(err)
        dir, err := os.ReadDir(cwd)
        check(err)
        res := []fs.DirEntry{}
        for _,v := range dir {
            res = append(res, v)
        }
        return res
    } else {
        c ,err := os.Stat(path[0][0])
        check(err) 
        if c.IsDir() {
            d,err := os.ReadDir(path[0][0])
            check(err)
            res := []fs.DirEntry{}
            for _, entry := range d {
                res = append(res, entry)
            }
            return res
        } else {return nil}
    }
}

func longListing(path ...[]string) []string {
    data := []string{}
    if path == nil || path[0][0] == "." {
        dir := lsFunc()
        tmp := transData(dir)
        for _,v := range tmp {
            data = append(data, v)
        }
    } else {
        dir := lsFunc(path[0])
        tmp := transData(dir)
        for _,v := range tmp {
            data = append(data, v)
        }
    }
    return data
}

func transData(dir []fs.DirEntry) []string {
    res := make([]string, 0)
    tmp := make([][]interface{},0)  
    for _,v := range dir {
        info,err := v.Info()
        check(err)
        c := make([]interface{}, 0)
        if v != nil {
            c = append(c, info.Mode())
            c = append(c, info.Name())
            c = append(c, info.Size())
            tmp = append(tmp, c)
        }
    } 
    for _,v := range tmp {
        s := fmt.Sprint(v)
        res = append(res, s)
    }
    return res
}
