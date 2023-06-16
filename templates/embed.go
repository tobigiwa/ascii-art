package pages

import "embed"

var (
	//go:embed base.html
	BaseHTML embed.FS

	//go:embed ascii-art.html
	AsciiArtHTML embed.FS
)
