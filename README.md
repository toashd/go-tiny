Short Go URL Generator
---

Go implementation for generating Tiny URL- and bit.ly-like URLs.

A bit-shuffling approach is used to avoid generating consecutive, predictable URLs. However, the algorithm is deterministic and will guarantee that no collisions will occur.

The URL alphabet is fully customizable and may contain any number of characters. By default, digits and lower-case letters are used, with some removed to avoid confusion between characters like o, O and 0. The default alphabet is shuffled and has a prime number of characters to further improve the results of the algorithm.

The block size specifies how many bits will be shuffled. The lower BLOCK_SIZE bits are reversed. Any bits higher than BLOCK_SIZE will remain as is. BLOCK_SIZE of 0 will leave all bits unaffected and the algorithm will simply be converting your integer to a different base.

The intended use is that incrementing, consecutive integers will be used as keys to generate the short URLs. For example, when creating a new URL, the unique integer ID assigned by a database could be used to generate the URL by using this module. Or a simple counter may be used. As long as the same integer is not used twice, the same short URL will not be generated twice.

The module supports both encoding and decoding of URLs. The min_length parameter allows you to pad the URL if you want it to be a specific length.


Getting Started
--
1: Download the package

```
go get github.com/toashd/go-tiny/tiny
```

2: Import go-tiny/tiny to your Go project

```
import "github.com/toashd/go-tiny/tiny"
```

Example
--

```
package main

import (
	"fmt"
	"github.com/toashd/go-tiny/tiny"
)

func main() {
	// Url to decode (id)
	var url = 1337

	// Creates new TinyUrl
	t := tiny.NewTinyUrl()

	// Encodes url
	var enc = t.EncodeUrl(url, 5)

	// Prints ecoded url
	fmt.Printf("Encoded Url: %s\n", enc)

	// Prints decoded url
	fmt.Printf("Decoded Url: %v\n", t.DecodeUrl(enc))
}
```

This example can also be found in the example.go file.

License
--
go-tiny is available under the MIT license.

