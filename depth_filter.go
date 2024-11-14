package finder

import (
	"github.com/iveahugeship/gofind/utils"
)

type MinDepthFilter struct {
	n int
}

func (f MinDepthFilter) Match(root string, path string) (bool, error) {
	return f.n <= utils.CountDepth(root, path), nil
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
	return f.n >= utils.CountDepth(root, path), nil
}

func ByMaxDepth(n int) Option {
	return func(f *Finder) {
		f.filters = append(f.filters, MaxDepthFilter{
			n: n,
		})
	}
}
