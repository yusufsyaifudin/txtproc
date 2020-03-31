# Text Processor

This is Go module for processing text. Currently only support word splitter.

## Usage

```
go get -v github.com/yusufsyaifudin/txtproc
```

Then use import module `ysf/txtproc` in import path.

Example:

```go
package main

import (
    "context"
    "fmt"
    "log"
    "ysf/txtproc"
)

func main() {
    text := "This is a words collection."
    words, err := txtproc.WordSeparator(context.Background(), text)

    if err != nil {
        log.Fatal(err.Error())
        return
     }

    for _, word := range words {
        fmt.Println(word.GetOriginal())
    }
}
```

it should print:

```
This

is

a

words

collection.
```

## Benchmark

Benchmark on Macbook Pro 16GB, Quad-Core Intel Core i5 2.4Ghz

```
goos: darwin
goarch: amd64
pkg: ysf/txtproc
BenchmarkWordSeparator_1Word-8           3089930               386 ns/op
BenchmarkWordSeparator_100Words-8          37720             31883 ns/op
BenchmarkWordSeparator_200Words-8          18387             65294 ns/op
PASS
ok      ysf/txtproc     5.381s
```