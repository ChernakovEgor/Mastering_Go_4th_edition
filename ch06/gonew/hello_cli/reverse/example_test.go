// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reverse_test

import (
	"fmt"

	"example.com/hello_cli/reverse"
)

func ExampleString() {
	fmt.Println(reverse.String("hello"))
	// Output: olleh
}
