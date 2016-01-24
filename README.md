# goof

About
-----

This is my rather pointless entry for the 2016 [Gopher Gala](http://www.gophergala.com).

Instructions
------------

+ Make sure you have [Go](http://golang.org/dl) and [Git](https://git-scm.com),
  although if you're reading this you probably do.

+ Setup and [test your Go installation](https://golang.org/doc/install#testing)

+ Install the Go dependencies:

```shell
go get -u github.com/fatih/color
```

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

+ Usage:

```shell
goof -list   # List git repositories found under your $HOME/%HOME% directory
```

```shell
goof -update # Update git repositories found under your $HOME/%HOME% directory
```

+ License

```shell

        DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
                    Version 2, December 2004

 Copyright (C) 2016 Julien Castelain <julien@users.noreply.github.com>

 Everyone is permitted to copy and distribute verbatim or modified
 copies of this license document, and changing it is allowed as long
 as the name is changed.

            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
   TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION

  0. You just DO WHAT THE FUCK YOU WANT TO.

```

