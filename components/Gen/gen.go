package gen

import (
	"embed"
)

func InstallerExtract(Embedded embed.FS, EmbeddedFN string) error {
	fileData, err := Embedded.ReadFile(Emb_FileName)
	if err != nil {
		return err
	}
	return err
}
