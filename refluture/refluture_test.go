// Copyright (c) 2021 James Bowes. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package refluture_test

import (
	"fmt"
	"time"

	"github.com/jbowes/future/refluture"
)

func blocking(name string) (string, error) {
	time.Sleep(10 * time.Millisecond)

	return "hello, " + name, nil
}

func Example() {
	fut := refluture.New(func() (string, error) { return blocking("you") })

	var out string
	err := fut.Await(&out)

	fmt.Println(out, err)

	// Output:
	// hello, you <nil>
}
