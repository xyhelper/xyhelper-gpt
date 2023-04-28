package main

import (
	_ "pandora-go/internal/packed"

	_ "pandora-go/backend-api"

	"github.com/gogf/gf/v2/os/gctx"

	"pandora-go/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
