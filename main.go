package main

import (
	"github.com/arymaulanamalik/sicepat_sample/cmd"
	// _ "github.com/arymaulanamalik/sicepat_sample/docs"

	_ "gocloud.dev/docstore/memdocstore"
	_ "gocloud.dev/docstore/mongodocstore"
)

// @title Authorization Service API
// @description Authorization Service API
// @BasePath /core/authorization-svc
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cmd.Run()
}
