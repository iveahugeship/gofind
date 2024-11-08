package finder

type Filter interface {
	Match(root string, path string) (bool, error)
}
