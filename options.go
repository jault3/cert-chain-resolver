package main

import (
	"io"
	"os"

	flags "github.com/jessevdk/go-flags"
)

type FlagOptions struct {
	OutputFilename         string `short:"o" long:"output" value-name:"OUTPUT_FILE" description:"Output filename (default: stdout)"`
	OutputDerFormat        bool   `short:"d" long:"der" description:"Output DER format"`
	OutputIntermediateOnly bool   `short:"i" long:"intermediate-only" description:"Output intermediate certificates only"`
	Args                   struct {
		InputFilename string `positional-arg-name:"INPUT_FILE" description:"Input filename (default: stdin)"`
	} `positional-args:"yes"`
}

type Options struct {
	InputReader            io.Reader
	OutputWriter           io.Writer
	OutputDerFormat        bool
	OutputIntermediateOnly bool
}

func GetOptions() (*Options, error) {
	var (
		flagOptions FlagOptions
		options     Options
	)

	flagsParser := flags.NewParser(&flagOptions, flags.HelpFlag|flags.PassDoubleDash)
	if _, err := flagsParser.Parse(); err != nil {
		return nil, err
	}

	if flagOptions.Args.InputFilename != "" {
		reader, err := os.Open(flagOptions.Args.InputFilename)
		if err != nil {
			return nil, err
		}

		options.InputReader = reader
	} else {
		options.InputReader = os.Stdin
	}

	if flagOptions.OutputFilename != "" {
		writer, err := os.Create(flagOptions.OutputFilename)
		if err != nil {
			return nil, err
		}

		options.OutputWriter = writer
	} else {
		options.OutputWriter = os.Stdout
	}

	options.OutputDerFormat = flagOptions.OutputDerFormat
	options.OutputIntermediateOnly = flagOptions.OutputIntermediateOnly

	return &options, nil
}
