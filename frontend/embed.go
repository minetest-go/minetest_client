package frontend

import (
	"embed"
)

//go:embed index.html
var Webapp embed.FS
