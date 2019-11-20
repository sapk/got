//+build mage

package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var gocmd = "go"

func init() {
	gocmd = mg.GoCmd()
}

// Generate go generate.
func Generate() error {
	return sh.RunV(gocmd, "generate", "./...")
}

// Build the binary.
func Build() error {
	/* Not needed since code is "vendored"
	if err := Generate(); err != nil {
		return err
	}
	*/
	return sh.RunV(gocmd, "build", "-mod=vendor", "-ldflags", "-s -w", ".")
}

// Run run with dev tags.
func Run() error {
	/* Not needed since code is "vendored"
	if err := Generate(); err != nil {
		return err
	}
	*/
	return sh.RunV(gocmd, "run", "-mod=vendor", "--tags", "dev", ".")
}

type Deps mg.Namespace

// Vendor store deps in vendor/ folder.
func (Deps) Vendor() error {
	return sh.RunV(gocmd, "mod", "vendor")
}

// Get get deps to mod store.
func (Deps) Get() error {
	return sh.RunV(gocmd, "get", "-v", "./...")
}

// Update update deps to last version.
func (Deps) Update() error {
	if err := sh.RunV(gocmd, "get", "-u", "-v", "./..."); err != nil {
		return err
	}
	return sh.RunV(gocmd, "mod", "tidy")
}
