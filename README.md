# Rwanda

This is a Go package which contains all administrative locations. I includes all provinces, districts, sectors, cells and villages.

## Usage Examples

Here are some examples of how to use this package:

1. Install The package
```bash
go get github.com/samueltuyizere/rwanda@latest
```
2. Import and use in your code

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
- #### Provincial Level
- [x] get all provinces as standalone entities
- [ ] get all districts in a province
- [ ] get all sectors in a province
- [ ] get all cells in a province
- [ ] get all villages in a province
- #### District Level
- [x] get all districts as standalone entities
- [ ] get all sectors in a district
- [ ] get all cells in a district
- [ ] get all villages in a district
- #### Sector Level
- [x] get all sectors as standalone entities
- [ ] get all cells in a sector
- [ ] get all villages in a sector
- #### Cell Level
- [x] get all cells as standalone entities
- [ ] get all villages in a cell
- #### Misc tasks
- [ ] Tests


## Contributing

We welcome contributions from the community! If you're interested in contributing to this package, please follow these guidelines:

1. Fork the repository and clone it locally.
2. Create a new branch for your contribution.
3. Make your changes and ensure that the code passes all tests.
4. Commit your changes and push them to your forked repository.
5. Open a pull request, describing the changes you've made and why they should be merged.

Please ensure that your contributions adhere to our code of conduct and follow the established coding style.

Thank you for considering contributing to this package!
