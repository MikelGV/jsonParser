package cli

import "flag"

type Cli struct {
    FileName string
}


func New() *Cli {

    return &Cli{
        FileName: determineFn(),
    }
}

func determineFn() string {
    if flag.NArg() > 0 {
        return flag.Arg(0)
    }

    return ""
}
