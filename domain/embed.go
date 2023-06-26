package domain

import (
	"embed"
)

var (
	// Banner is the embeded filesystem of different sytles
	//go:embed *txt
	Banner embed.FS
)
