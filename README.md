# gofind 🔎

`gofind` is a Go library that offers functionality akin to the Unix `find` command, allowing developers to search for files and directories based on various criteria within their Go applications.

## Features

- **Name Filtering**: Search for files or directories by name patterns.
- **Regular Expression Filtering**: Use regex patterns to match file or directory names.
- **Type Filtering**: Filter results by file type (e.g., regular files, directories).
- **Depth Control**: Limit the search to a specific directory depth.
- **Custom Predicates**: Define custom search criteria using predicates.

## Installation

To install the `gofind` library, ensure you have [Go](https://golang.org/dl/) installed, then run:

```bash
go get github.com/iveahugeship/gofind
```

## Usage

`gofind` is a really simple Go library. Check this out:

```golang
package main

import (
	"log"

	"github.com/iveahugeship/gofind/finder"
	"github.com/iveahugeship/gofind/ftype"
)

func main() {
	// Create a new Finder with various search criteria
	f := finder.NewFinder(
		"/Users/iveahugeship/Projects/personal/gofind",
        finder.BySinceMtime("2024-11-16"),	// Filters results to include items modified on or after November 16, 2024
		finder.ByType(ftype.DirType),       // Search for directories only
		finder.ByMinDepth(1),               // Start search at depth 1
		finder.ByMaxDepth(3),               // Search up to depth 3
		finder.ByPerm(0777),                // Match directories with 0777 permissions
		finder.WithPrintExec(),             // Print each matching result
	)

	// Execute the search
	hits, err := f.Find()
	if err != nil {
		log.Fatal(err)
	}
    fmt.Println(len(hits))
}
```

Other examples are available under [examples/](examples/) directory.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request with your changes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE.md) file for details.
