# Murmur3 for Go

<a href="https://godoc.org/github.com/tidwall/murmur3"><img src="https://img.shields.io/badge/api-reference-blue.svg?style=flat-square" alt="GoDoc"></a>

A port of the [Murmur3](https://github.com/PeterScott/murmur3) hash function. 

Installing
----------

```
go get -u github.com/tidwall/murmur3
```

Example
-------

```go
println(murmur3.Sum32("hello"))
// output: 613153351
```

Contact
-------
Josh Baker [@tidwall](http://twitter.com/tidwall)

License
-------
This is free and unencumbered software released into the public domain.
