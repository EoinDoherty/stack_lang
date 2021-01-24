package main

import (
	"flag"
	"fmt"
	"stack_lang/compiler"
	"stack_lang/runtime"
)

func main() {

    compile := flag.Bool("c", false, "Compile an executable from a file")
    ifile := flag.String("infile", "", "Source code file to compile")
    interpret := flag.Bool("i", false, "Interpret a compiled executable")
    ofile := flag.String("outfile", "a.out", "Name of executable file; defaults to a.out")
    flag.Parse()

    if *compile {
        if len(*ifile) == 0 {
            fmt.Println("Please enter a filename with -ifile")
            return
        }

        err := compiler.CompileFile(*ifile, *ofile)

        if err != nil {
            fmt.Printf("Error: %v\n", err)
        }
    }

    runFile := *ifile

    if *interpret && *compile{
        runFile = *ofile
    }

    err := runtime.RunFile(runFile)

    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }
}
