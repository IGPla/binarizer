/* Utils package */
package utils

import (
	"flag"
	"log"
	"os"
)

type Args struct {
	ScriptType *string
	ScriptFile *string
	BinaryFile *string
	GOOS       *string
	GOARCH     *string
	Encrypt    *bool
}

/*
Description: Return script call based on script type
*/
func (args *Args) ScriptCall() *string {
	var result string
	switch *args.ScriptType {
	case "python":
		result = "python -c"
	case "bash":
		result = "bash -c"
	default:
		log.Fatalf("Script type provided is not supported: %v", args.ScriptType)
	}
	return &result
}

/*
Description: Parse binarizer arguments
*/
func ParseArguments() Args {
	scriptType := flag.String("script-type", "python", "Script type. Available values are: {python, bash}")
	scriptFile := flag.String("script-file", "./myscript.sh", "Script file")
	binaryFile := flag.String("binary-file", "./mybinary", "Compiled binary")
	goos := flag.String("goos", "linux", "GO Target Operating System (GOOS env var)")
	goarch := flag.String("goarch", "amd64", "GO Target Architecture (GOARCH env var)")
	encrypt := flag.Bool("encrypt", false, "Generate an encrypted binary with random key")

	flag.Parse()
	var args Args = Args{scriptType, scriptFile, binaryFile, goos, goarch, encrypt}
	return args
}

/*
Description: Remove temporary files
*/
func RemoveTemporaryFiles(files ...*string) {
	for _, _file := range files {
		if err := os.Remove(*_file); err != nil {
			log.Printf("An error arised while removing %v: %v", _file, err)
		}
	}
}
