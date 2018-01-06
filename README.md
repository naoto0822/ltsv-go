# ltsv-go

[![Build Status](https://travis-ci.org/naoto0822/ltsv-go.svg?branch=master)](https://travis-ci.org/naoto0822/ltsv-go)
[![GoDoc](https://godoc.org/github.com/naoto0822/ltsv-go/ltsv?status.svg)](https://godoc.org/github.com/naoto0822/ltsv-go/ltsv)
[![Go Report Card](https://goreportcard.com/badge/github.com/naoto0822/ltsv-go)](https://goreportcard.com/report/github.com/naoto0822/ltsv-go)
[![License](https://img.shields.io/badge/license-MIT-green.svg?style=flat)](https://github.com/naoto0822/ltsv-go/blob/master/LICENSE.txt)

LTSV pkg for Go lang.

## Installation

```sh
$ go get github.com/naoto0822/ltsv-go/ltsv
```

## Usage

### Marshal

```go
type User struct {
  Name string   `ltsv:"name"`
  Age  int      `ltsv:"age"`
  Tags []string `ltsv:"tags"`
}

user := User{
  Name: "naoto0822",
  Age:  27,
  Tags: []string{"sports", "music", "enginner"},
}

log := ltsv.Marshal(user)
fmt.Println(log)
// name:naoto0822	age:27	tags:["sports","music","enginner"]
```

### Unmarshal

```go
type User struct {
  Name string   `ltsv:"name"`
  Age  int      `ltsv:"age"`
  Tags []string `ltsv:"tags"`
}

input := "name:naoto0822\tage:27\ttags:[\"sports\",\"music\",\"enginner\"]"
user := User{}
err := ltsv.Unmarshal(input, &user)
if err == nil {
  fmt.Println(user)
}
// {naoto0822 27 [sports music enginner]}
```

## TODO

- [ ] Unmarshal test
- [ ] handle not support type
- [ ] implement logger

## License

```
MIT License

Copyright (c) 2017 naoto yamaguchi

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
