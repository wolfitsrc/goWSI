// WSI short for Wolf Software Installer
// This projects idea is to create a simple CLI tool for making a installer for a binary program.
// The first couple features to be implemented are:
// 1. Allow for the installer to add path's to the PATH environment variable.
// 2. Allow for the installer to hold static binary files that get copied to a specified location.

package main

import (
	"embed"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/sys/windows"

	"github.com/wolfitsrc/goWSI/components/gen"
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
		if checkAdmin() {
			fmt.Println(err)
		} else {
			becomeAdmin()
		}
	}
}

func becomeAdmin() {
	verb := "runas"
	exe, _ := os.Executable()
	cwd, _ := os.Getwd()
	args := strings.Join(os.Args[1:], " ")

	verbPtr, _ := syscall.UTF16PtrFromString(verb)
	exePtr, _ := syscall.UTF16PtrFromString(exe)
	cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
	argPtr, _ := syscall.UTF16PtrFromString(args)

	var showCmd int32 = 1 //SW_NORMAL

	err := windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd)
	if err != nil {
		fmt.Println(err)
		err := gen.InstallerExtract(Embedded, EmbFileName, ExportLoc)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func checkAdmin() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")

	return err == nil
}
