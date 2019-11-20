# GoT

A map tile server implemented in go

## Quickstart

- Download last release of GoT: https://github.com/sapk/got/releases.
- Generate or Download a .mbtiles file like from https://openmaptiles.com/downloads/planet/
- Run: `./got my-file.mbtiles`
- Go to http://localhost:3000

## Build from sources

Use mage helper to run tasks:
```
go run -mod=vendor mage.go -l
go run -mod=vendor mage.go generate build
 -- OR --
alias m="go run -mod=vendor mage.go"
m generate build
 -- For more --
m -l
```
