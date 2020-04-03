# Text Processor

[![Build Status](https://travis-ci.com/yusufsyaifudin/txtproc.svg?branch=master)](https://travis-ci.com/yusufsyaifudin/txtproc)
[![codecov](https://codecov.io/gh/yusufsyaifudin/txtproc/branch/master/graph/badge.svg)](https://codecov.io/gh/yusufsyaifudin/txtproc)
[![Go Report Card](https://goreportcard.com/badge/github.com/yusufsyaifudin/txtproc)](https://goreportcard.com/report/github.com/yusufsyaifudin/txtproc)

#### [View on Sourcegraph](https://sourcegraph.com/github.com/yusufsyaifudin/txtproc)

This is Go module for processing text. Currently only support word splitter.

## Usage

https://godoc.org/github.com/yusufsyaifudin/txtproc

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

## Work In Progress

- [x] Word Splitter (split by space, tab, new line)
- [x] Word Replacer (replace word with DIY Replacer function)
- [ ] Profanity Filter
