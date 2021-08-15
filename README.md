# future

An implementation of futures/promises for `go1.18+` type parameters (generics).

## Why would I want this?

Go's programming model of implicitly yielding to other `goroutines` leads to easy to read code, without extra syntax. Other languages might require additional keywords, callback functions, etc.

But `async`/`await` with Promises in JavaScript lets you choose if you want to have the current execution context block and yield to
other contexts while waiting for results, or do other work in the meantime. Doing this in Go requires running new `goroutines` and synchronizing with `channel`s or mutexes. This is powerful and flexible, but sometimes it's too much boilerplate.

Where `async`/`await` style languages let you easily opt-in to waiting, what if you could easily opt-out of waiting in Go?

## Example

```go
package main

import (
	"context"
	"fmt"
	"time"

    "github.com/jbowes/future"
)

func SomeAsyncThing(ctx context.Context) (int, error) {
	fmt.Println("doing some async thing")
	time.Sleep(2 * time.Second)
	return 2, nil
}

func main() {
	ctx := context.Background()

	// Create a new future to do SomeAsyncThing in the background
	fut := future.New(func() (int, error) { return SomeAsyncThing(ctx) })

	time.Sleep(1 * time.Second)
	fmt.Println("doing the main thing")

	// Wait for the results, if they're not done yet.
	x, err := fut.Await()
	fmt.Println("got result", x, err)
}
```

## Limitations

### An implementation for each supported argument count

There is no planned support for [variadic type paramters](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#omissions). To support wrapping funcs with 1, 2, 3, and 4 arguments, we need 4 different implementations. For callers, this means using `New` or `New3` etc. As `Await` is a method on the various `Future`s, its name doesn't have to change.

### Awkward syntax for `New`

For `New` to work and be type safe, you must pass a function to it, with a known set of inputs (that is, zero inputs):

```go
fut := future.New(func() (int, error) { return SomeAsyncThing(ctx, arg1, arg2) })
```

 A nicer syntax would be one supported by variadic type parameters:

```go
fut := future.New(SomeAsyncThing, ctx, arg1, arg2)
```

Supporting this with multiple implementations would require `(# of inputs) * (# of outputs)` implementations. For end users, calling
`future.New3_4` seems too unwieldy.


## What about reflection?

There are a handful of reflection-based future/promise Go libraries already available. For convenient comparison, one is included here under [`refluture`](./refluture). Note that other libraries include features like promise chaining and racing, and tend to expose APIs similar to JavaScript's promises (eg `resolve` and `reject`). I couldn't be bothered to add those and don't think they fit well in Go code, where you can fall back to using `goroutine`s, `channel`s, `select`, etc.

## Links

- [The type parameters proposal](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md)
- [chebyrash/promise](https://github.com/chebyrash/promise)
- [fanliao/go-promise](https://github.com/fanliao/go-promise)