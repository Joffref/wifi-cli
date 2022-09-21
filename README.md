# WIFI-CLI

WIFI CLI is a command line interface to select WIFI networks.
There's two main features:
 - AP mode: Select the best wifi channel for your AP.
 - Terminal mode: Select the best wifi network.

# Install the CLI

```bash
go get github.com/Joffref/wifi-cli
```

> Note: You need to have `go` installed on your machine.
>See [golang.org](https://golang.org/doc/install) for more information.
> 
> You also need to have `git` installed on your machine.
> See [git-scm.com](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git) for more information.

# Build the CLI from source

```bash
git clone https://github.com/Joffref/wifi-cli
cd wifi-cli
make build
```

# Usage

## AP mode

```bash
Usage:
  wifi-cli ap [flags]

Flags:
  -a, --AccessPointNumberWeight int   AccessPointNumberWeight (default 1)
  -r, --CoverageWeight int            CoverageWeight (default 1)
  -s, --SignalStrengthWeight int      SignalStrengthWeight (default 1)
  -h, --help                          help for ap
  -i, --interface string              wifi interface (default "wlan0")
```

## Terminal mode

```bash
 Usage:
  wifi-cli terminal [flags]

Flags:
  -h, --help               help for terminal
  -i, --interface string   wifi interface (default "wlan0")
```

# Using the library
As a library, you can use the `ap` and `terminal` packages to get the best wifi channel or the best wifi network.
Plus, we define an interface `SelectionMiddleware` to define your own selection algorithm or use the included one (`coverage`, `empty`, `number`, `signal`).
```go
package ap

//...

// SelectionMiddleware is a middleware that selects the best channel given a criteria.
type SelectionMiddleware interface {
	// Select ranks the channels based on the criteria.
	Select(ScoredChannels map[int]int, UsedChannels map[int]*Channel) (map[int]int, error)
	// Criteria returns the criteria of the middleware.
	Criteria() string
	// Name returns the name of the criteria.
	Name() string
	// SetWeight sets the weight of the criteria.
	// The weight is an int between 0 and 100.
	SetWeight(int)
}
```

Then you can use the `ap` package to get the best channel.
```go
package main

import (
    "log"
    "github.com/Joffref/wifi-cli/ap"
)

func main() {
	chain := ap.SelectionChain{
		&ap.UnoccupiedChannel{
			Weight: ap.Infinite,
		},
		&ap.AccessPointNumber{
			Weight: 1,
		},
		&ap.Signal{
			Weight: 20,
		},
		&ap.Coverage{
			Weight: 100,
		},
	}
	chanel, err := BestChanel(ifname, chain)
	if err != nil {
		log.Errorf("Error while finding best channel: %v", err)
	}
	log.Infof("Best channel is %v", chanel)
}
func BestChanel(ifname string, chain ap.SelectionChain) (int, error) {
	return ap.FindBestChannel(ifname, chain)
}
```

# To go further

## Roadmap
- [x] Add a `terminal` mode to select the best wifi network.
- [x] Add a `coverage` criteria to select the best channel.
- [x] Add a `signal` criteria to select the best channel.
- [x] Add a `number` criteria to select the best channel.
- [x] Add a `empty` criteria to select the best channel.
- [x] Add a `weight` to each criteria.
- [x] Add a `SelectionChain` to select the best channel.
- [x] Add a `SelectionMiddleware` to define your own selection algorithm.
- [ ] Enhance the scoring algorithm or provide a way to define your own scoring algorithm.
- [ ] Add verbosity level.
- [ ] Add a `--json` flag to output the result in json format.
- [ ] Add a `--csv` flag to output the result in csv format.
- [ ] Add a `--yaml` flag to output the result in yaml format.

# License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details