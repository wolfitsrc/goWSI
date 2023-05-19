// WSI short for Wolf Software Installer
// This projects idea is to create a simple CLI tool for making a installer for a binary program.
// The first couple features to be implemented are:
// 1. Allow for the installer to add path's to the PATH environment variable.
// 2. Allow for the installer to hold static binary files that get copied to a specified location.

package main

import (
	"embed"
	"fmt"

	"github.com/wolfit-src/goWSI/components/gen"
)

// ///////////////////////////////////////////
// configure your file name below
const (
	EmbFileName = "filename.txt"
	ExportLoc   = "c:/program files/wolfit/bin"
)

//go:embed filename.txt
var Embedded embed.FS

// Next run the following command: go run main.go -o name-installer.exe
/////////////////////////////////////////////

func main() {
	err := gen.InstallerExtract(Embedded, EmbFileName, ExportLoc)
	if err != nil {
		fmt.Println(err)
	}
}
