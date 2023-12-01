package main

import (
	"fmt"
	"os"

	"github.com/alexflint/go-arg"
	"github.com/xyproto/autoimport"
)

// Args struct to hold command line arguments
type Args struct {
	Filename            string `arg:"positional" help:"The filename to process"`
	Apply               bool   `arg:"-a,-w,--apply,--write" help:"Write changes back to the file"`
	Verbose             bool   `arg:"-v,--verbose" help:"Enable verbose output"`
	KeepExistingImports bool   `arg:"-k,--keep" help:"Remove existing imports" default:"true"`
	DeGlob              bool   `arg:"-d,--deglob" help:"Expand wildcard imports" default:"true"`
}

// Version returns the current program name and version
func (Args) Version() string {
	return "Grog 1.0.1"
}

func main() {
	var args Args
	arg.MustParse(&args)

	// Check if Filename is provided
	if args.Filename == "" {
		fmt.Fprintln(os.Stderr, "A filename is required.")
		os.Exit(1)
	}

	data, err := autoimport.Fix(args.Filename, !args.KeepExistingImports, args.DeGlob, args.Verbose)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not re-organize imports in %s: %s\n", args.Filename, err)
		os.Exit(1)
	}

	// Just output the contents with the fixed imports and exit the program if -a or --apply are not supplied flag
	if !args.Apply {
		fmt.Println(string(data))
		return
	}

	// Get the current file mode
	fileInfo, err := os.Stat(args.Filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get file mode for %s: %s\n", args.Filename, err)
		os.Exit(1)
	}

	// Write the data back to the file with the original file mode
	err = os.WriteFile(args.Filename, data, fileInfo.Mode())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not write to %s: %s\n", args.Filename, err)
		os.Exit(1)
	}
}
