package main

import (
	"fmt"
	"os"

	"github.com/alexflint/go-arg"
	"github.com/xyproto/autoimport"
)

// Args struct to hold command line arguments
type Args struct {
	Filename              string `arg:"positional,required"`
	Apply                 bool   `arg:"-a,--apply" help:"Write changes back to the file"`
	Verbose               bool   `arg:"-v,--verbose" help:"Enable verbose output"`
	RemoveExistingImports bool   `arg:"-r,--remove" help:"Remove existing imports" default:"true"`
	DeGlob                bool   `arg:"-d,--deglob" help:"Expand wildcard imports" default:"true"`
}

func main() {
	var args Args
	arg.MustParse(&args)

	data, err := autoimport.Fix(args.Filename, args.RemoveExistingImports, args.DeGlob, args.Verbose)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not re-organize imports in %s: %s\n", args.Filename, err)
		os.Exit(1)
	}

	if args.Apply {
		// Get the current file mode
		fileInfo, err := os.Stat(args.Filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not get file mode for %s: %s\n", args.Filename, err)
			os.Exit(1)
		}

		// Write the data back to the file with the original file mode
		err = os.WriteFile(args.Filename, data, fileInfo.Mode())
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not write to %s: %s\n", args.Filename, err)
			os.Exit(1)
		}
	} else {
		fmt.Println(string(data))
	}
}
