// +build dev

package swagger

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	_, d := filepath.Split(pwd)
	if d == "swagger" { //For go generate command
		Swagger = http.Dir("../../assets/swagger")
	}
}

// Swagger contains project swagger assets.
var Swagger http.FileSystem = http.Dir("assets/swagger")
