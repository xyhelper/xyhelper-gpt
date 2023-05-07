package main

import (
	_ "xyhelper-gpt/internal/packed"

	_ "xyhelper-gpt/backend-api"

	_ "xyhelper-gpt/v1"

	"github.com/gogf/gf/v2/os/gctx"

	"xyhelper-gpt/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
