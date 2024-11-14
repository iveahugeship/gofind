package main

import (
	"fmt"

	finder "github.com/iveahugeship/gofind"
	"github.com/iveahugeship/gofind/utils"
)

func main() {
	f := finder.NewFinder(
		"/Users/iveahugeship/Projects/pulsepoint/kubernetes/monitoring/configs/prometheus-server/alerts.d",
		// finder.ByName("*.yaml"),
		finder.BySinceDate("2024-11-13"),
		finder.ByType(utils.DirType),
		finder.ByMinDepth(1),
		finder.ByMaxDepth(2),
	)
	hits, _ := f.Find()
	// fmt.Println(len(hits))
	fmt.Println(hits)
}
