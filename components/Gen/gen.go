package gen

import (
	"embed"
	"fmt"
	"os"
)

func InstallerExtract(Embedded embed.FS, FileName string, ExportLoc string) error {
	fmt.Println("function is running")
	fileData, err := Embedded.ReadFile(FileName)
	if err != nil {
		return err
	}

	err = os.MkdirAll(ExportLoc, 0755)
	if err != nil {
		return err
	}

	info, err := os.Stat(ExportLoc)
	if err != nil {
		return err
	} else {
		fmt.Println(info.Name())
	}

	exportPath := ExportLoc + "/" + FileName
	err = os.WriteFile(exportPath, fileData, 0755)
	if err != nil {
		return err
	}
	fmt.Println("function has no errors!")
	return nil
}
