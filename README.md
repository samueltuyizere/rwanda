# Rwanda

This is a Go package which contains all administrative locations. I includes all provinces, districts, sectors, cells and villages.

## Usage Examples

Here are some examples of how to use this package:

```bash
go get github.com/samueltuyizere/rwanda
```

```go
package main
import (
    "github.com/samueltuyizere/rwanda"
    "fmt"
)

func main(){
    regions := rwanda.GetALlRegions()
    fmt.Print(regions)
}
```

## Upcoming
[] Tests

## Contributing

We welcome contributions from the community! If you're interested in contributing to this package, please follow these guidelines:

1. Fork the repository and clone it locally.
2. Create a new branch for your contribution.
3. Make your changes and ensure that the code passes all tests.
4. Commit your changes and push them to your forked repository.
5. Open a pull request, describing the changes you've made and why they should be merged.

Please ensure that your contributions adhere to our code of conduct and follow the established coding style.

Thank you for considering contributing to thi package!
