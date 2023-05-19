package conf

import (
	"embed"
)

// configure your file name below
var EmbeddedFN string = "../../filename.txt"

//go:embed ../filename.txt
var Embedded embed.FS
