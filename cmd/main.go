package main

import (
	"fmt"

	finder "github.com/iveahugeship/gofind"
	"github.com/iveahugeship/gofind/utils/file/ftype"
)

func main() {
	f := finder.NewFinder(
		"/Users/iveahugeship/Projects/pulsepoint/kubernetes/monitoring/configs/prometheus-server/alerts.d",
		// finder.BySinceMtime("2024-11-16"),
		finder.ByType(ftype.DirType),
		finder.ByMinDepth(1),
		finder.ByMaxDepth(3),
		finder.WithPrintExec(),
		// finder.ByRegex(regexp.MustCompile(".*asd/")),
		// TODO:
		// perm filter
		// uid filter
		// gid filter
		// size filter
		// fs filter?
	)
	hits, _ := f.Find()
	fmt.Println(len(hits))
}
