// +build dev

package webapp

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
	if d == "webapp" { //For go generate command
		WebApp = http.Dir("../../assets/webapp/dist")
	}
}

// WebApp contains project WebApp assets.
var WebApp http.FileSystem = http.Dir("assets/webapp/dist")
