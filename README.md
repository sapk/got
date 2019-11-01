# GoT [WIP]

A map tile server implemented in go

## Quickstart

- Download last release of GoT.
- Download a .mbtiles file from https://openmaptiles.com/downloads/planet/
- Run: `./got my-file.mbtiles`
- Go to http://localhost:3000

## Build from sources

Use mage helper to run tasks:
```
go run -mod=vendor mage.go -l
go run -mod=vendor mage.go webapp:build
go run -mod=vendor mage.go generate
go run -mod=vendor mage.go build
 -- OR --
alias m="go run -mod=vendor mage.go"
m -l
```
