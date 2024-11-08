package main

import (
	"fmt"

	"github.com/iveahugeship/go-finder"
)

func main() {
	f := finder.NewFinder(
		"/Users/iveahugeship/Projects/pulsepoint/kubernetes/monitoring/configs/prometheus-server/alerts.d",
		finder.ByName("*.yaml"),
	)
	hits, _ := f.Find()
	fmt.Println(hits)
}
