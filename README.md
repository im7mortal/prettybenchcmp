rename maybe to .benchmarkslog
5) расскрасить в разные цвета

[![Build Status](https://travis-ci.org/im7mortal/UTM.svg)](https://travis-ci.org/im7mortal/UTM)
[![Coverage Status](https://coveralls.io/repos/im7mortal/UTM/badge.svg?branch=master)](https://coveralls.io/r/im7mortal/UTM?branch=master)
[![GoDoc](https://godoc.org/github.com/im7mortal/prettybenchcmp?status.svg)](https://godoc.org/github.com/im7mortal/prettybenchcmp)

prettybenchcmp
===

Prettybenchcmp is cmd tool for automated comparision of results of benchmarks.
There are standart tool benchcmp. But it really akward to use for programmers 

How to use
-----

```
go get github.com/im7mortal/prettybenchcmp
// exec this command in a project directory
prettybenchcmp
//or with test and benchcmp flags
prettybenchcmp -short -benchtime 10s -count 2 -cpu 1,2,4 -changed -mag -best
```


First time when you use it. It will create .benchLog file which exist 
bench historu in the next format

```
PASS
BenchmarkUnquoteEasy	10000000	       182 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard	 1000000	      1117 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/benchcmp2	3.146s

yoshkarola 3235f14078c462b38c4b79912c1e34c868d34049

PASS
BenchmarkUnquoteEasy	10000000	       182 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard	 1000000	      1119 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/benchcmp2	3.141s
```

Where "yoshkarola" is separator.
Commit hash "3235f14078c462b38c4b79912c1e34c868d34049"

PASS
BenchmarkUnquoteEasy	10000000	       182 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard	 1000000	      1119 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/benchcmp2	3.141s

Is body of benchmark.

Supported flags from "go test"
-----
prettybenchcmp support couple of flags from ["go test"](https://golang.org/cmd/go/#hdr-Test_packages). It allow pass these flags

1. -benchtime
2. -count
3. -cpu
4. -short

Check ["go test" documentation](https://golang.org/cmd/go/#hdr-Description_of_testing_flags) for details.

Supported flags from benchcmp
-----
prettybenchcmp support all flags from [benchcmp](https://godoc.org/golang.org/x/tools/cmd/benchcmp).

1. -mag
2. -best
3. -changed

Check [benchcmp documentation](https://godoc.org/golang.org/x/tools/cmd/benchcmp) for details.

How it work 
-----

1) If file doesn't exist. 
It create file. Do first benchmark and write it to file.

2)  If file was inited. 
It clear file and write current benchmark

3) If file has history
It parse file.
If file exist and exist in git.
It check that file has hash of previous commit which exist .benchLog changes.
 If false
 it write to file separator and previous commit which exist .benchLog changes and current benchmark
 else
 it truncate file part which exist previous benchmark





Authors
-------

* Petr Lozhkin <im7mortal@gmail.com>

License
-------

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
