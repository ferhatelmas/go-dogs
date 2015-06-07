## go-dogs

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/ferhatelmas/go-dogs)
[![Build Status](https://travis-ci.org/ferhatelmas/go-dogs.png?branch=master)](https://travis-ci.org/ferhatelmas/go-dogs)

> Get popular dog names.

### Install

```
go get github.com/ferhatelmas/go-dogs
```

### Usage

```go
import "github.com/ferhatelmas/go-dogs"

dogs.RandomFemale()
//=> Maggie

dogs.IsMale("Roxy")
//=> false

dogs.Random()
//=> Abby, dogs.Female
```

### License

MIT © [ferhat elmas](http://ferhatelmas.com)
