package finder

import (
	"github.com/iveahugeship/gofind/utils/file/depth"
)

type MinDepthFilter struct {
	n int
}

func (f MinDepthFilter) Match(root string, path string) (bool, error) {
	d := depth.NewDepth(root)
	return d.UpperOf(path, f.n), nil
}

func ByMinDepth(n int) Option {
	return func(f *Finder) {
		f.filters = append(f.filters, MinDepthFilter{
			n: n,
		})
	}
}

type MaxDepthFilter struct {
	n int
}

func (f MaxDepthFilter) Match(root string, path string) (bool, error) {
	d := depth.NewDepth(root)
	return d.LowerOf(path, f.n), nil
}

func ByMaxDepth(n int) Option {
	return func(f *Finder) {
		f.filters = append(f.filters, MaxDepthFilter{
			n: n,
		})
	}
}
