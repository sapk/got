// +build ignore

package main

import (
	"log"

	"github.com/sapk/got/public/webapp"
	"github.com/shurcooL/vfsgen"
)

func main() {
	err := vfsgen.Generate(webapp.WebApp, vfsgen.Options{
		PackageName:  "webapp",
		VariableName: "WebApp",
		BuildTags:    "!dev",
		//Filename:     "webapp.go",
	})
	if err != nil {
		log.Fatal("%v", err)
	}
}
