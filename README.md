sprime [![GoDoc](https://godoc.org/github.com/otiai10/sprime?status.svg)](https://godoc.org/github.com/otiai10/sprime)
==========

- Find primary numbers
```sh
% sprime p 20
[2 3 5 7 11 13 17 19]
```
- Factorize numbers
```sh
% sprime f 329
[7 47]
```
- Reduce fractions
```sh
% sprime r 144/360
2/5
```

- all can be executed in go code
```go
fmt.Println(
        sprime.Factorize(329).List(),
        sprime.Factorize(329).Dict(),
)
// [7 47]
// map[7:1 47:1]
```

# install

![](http://cdn-ak.f.st-hatena.com/images/fotolife/o/otiai10/20141011/20141011203635.png)
