package main

import (
	_ "GF_Shop/internal/packed"

	_ "GF_Shop/internal/logic"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"GF_Shop/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
