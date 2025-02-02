package main

import "github.com/qaZar1/checkerNewGO/cron/internal/parser"

func main() {
	site := parser.NewSite("https://go.dev/doc/devel/release", "http://localhost:8001/api")
	site.ParseReleases()
}
