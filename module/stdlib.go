package module

import (
	"net/http"

	"github.com/rakyll/statik/fs"
	_ "github.com/superloach/defunct/stdlib/statik"
)

var Stdlib = func() http.FileSystem {
	stdlib, err := fs.New()
	if err != nil {
		panic(err)
	}

	return stdlib
}()
