# goof

### About

This is my rather pointless entry for the 2016 [Gopher Gala](http://www.gophergala.com).

### Instructions

+ Make sure you have [Go](http://golang.org/dl) and [Git](https://git-scm.com),
  although if you're reading this you probably do.

+ Setup and [test your Go installation](https://golang.org/doc/install#testing)

+ Install the Go dependencies:

```shell
go get -u github.com/fatih/color
```

 *This is an awesome go package! You should check it out.*

+ Clone this repo (until this repo is private):

```shell
# Make sure to put it in the "right" place

mkdir -p $GOPATH/src/github.com/gophergala2016

cd $GOPATH/src/github.com/gophergala2016

git clone git@github.com:gophergala2016/goof.git
```

+ Install it with:

```shell
go install github.com/gophergala2016/goof
```

### Usage

```shell
goof -list   # List git repositories found under your $HOME/%HOME% directory
```

```shell
goof -update # Update git repositories found under your $HOME/%HOME% directory
```

### License

The MIT License (MIT)

Copyright (c) 2016 Julien Castelain

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
