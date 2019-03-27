package main

import (
	"github.com/lnnujxxy/protoc-gen-star"
	"github.com/lnnujxxy/protoc-gen-star/lang/go"
	"github.com/lnnujxxy/protoc-gen-validate/module"
)

func main() {
	pgs.
		Init(pgs.DebugEnv("DEBUG_PGV")).
		RegisterModule(module.Validator()).
		RegisterPostProcessor(pgsgo.GoFmt()).
		Render()
}
