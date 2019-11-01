// +build ignore

package main

import (
	"log"

	"github.com/sapk/got/public/swagger"
	"github.com/shurcooL/vfsgen"
)

func main() {
	err := vfsgen.Generate(swagger.Swagger, vfsgen.Options{
		PackageName:  "swagger",
		VariableName: "Swagger",
		BuildTags:    "!dev",
		//Filename:     "swagger.go",
	})
	if err != nil {
		log.Fatal("%v", err)
	}
}
