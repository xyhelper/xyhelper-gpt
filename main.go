package main

import (
	_ "xyhelper-gpt/internal/packed"

	_ "xyhelper-gpt/backend-api"

	"github.com/gogf/gf/v2/os/gctx"

	"xyhelper-gpt/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
