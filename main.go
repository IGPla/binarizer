/* This package will binarize a given script */
package main

import (
	"github.com/IGPla/binarizer/binary"
	"github.com/IGPla/binarizer/utils"
	"log"
)

func main() {
	args := utils.ParseArguments()
	sourceFile, encodedKey := binary.SourceMaker(args.ScriptFile, args.ScriptCall(), args.Encrypt)
	defer utils.RemoveTemporaryFiles(sourceFile)
	binary.CompileSource(sourceFile, args.BinaryFile, args.GOOS, args.GOARCH)
	log.Printf("Finished binarization of %v (GOOS=%v, GOARCH=%v). The new binary should be found in %v",
		*args.ScriptFile, *args.GOOS, *args.GOARCH, *args.BinaryFile)
	if *args.Encrypt {
		log.Printf("Generated key for encryption (save it and be sure to call your binary file with BIN_KEY env path filled with this key): %v", encodedKey)
	}
}
