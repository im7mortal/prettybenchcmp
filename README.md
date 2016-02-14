[![Build Status](https://travis-ci.org/im7mortal/prettybenchcmp.svg?branch=master)](https://travis-ci.org/im7mortal/UTM)
[![GoDoc](https://godoc.org/github.com/im7mortal/prettybenchcmp?status.svg)](https://godoc.org/github.com/im7mortal/prettybenchcmp)

prettybenchcmp
===

**prettybenchcmp** is cmd tool for storage and comparison of benchmarks results. 
There is [a standard tool benchcmp](https://godoc.org/golang.org/x/tools/cmd/benchcmp),
but I don't think that the standard programmer will use it because it so boring! Do you doubt? Ok, 
just check instruction for benchcmp!
**prettybenchcmp** simply automates routine and do it good!

How to use
-----

```
go get github.com/im7mortal/prettybenchcmp
// exec this command in a project directory
prettybenchcmp
//or with test and benchcmp flags
prettybenchcmp -short -benchtime 10s -count 2 -cpu 1,2,4 -changed -mag -best
```

First time when you use it. It will create *.benchHistory* file which exist 
bench history in the special format. Check [format of .benchHistory](https://github.com/im7mortal/prettybenchcmp#Format-of-.benchHistory) for details

It will return standard output of [benchcmp](https://godoc.org/golang.org/x/tools/cmd/benchcmp). Like:
```
benchmark                            old ns/op     new ns/op     delta
BenchmarkGetLastBenchmark3651-4      1715          1736          +1.22%
BenchmarkGetLastBenchmark24456-4     1703          1689          -0.82%

benchmark                            old allocs     new allocs     delta
BenchmarkGetLastBenchmark3651-4      3              3              +0.00%
BenchmarkGetLastBenchmark24456-4     3              3              +0.00%

benchmark                            old bytes     new bytes     delta
BenchmarkGetLastBenchmark3651-4      4224          4224          +0.00%
BenchmarkGetLastBenchmark24456-4     4224          4224          +0.00%
```

What exaactly different from benchcmp
-----


Supported flags from "go test"
-----
**prettybenchcmp** support couple of flags from ["go test"](https://golang.org/cmd/go/#hdr-Test_packages). It allow pass these flags

1. -benchtime
2. -count
3. -cpu
4. -short

Check ["go test" documentation](https://golang.org/cmd/go/#hdr-Description_of_testing_flags) for details.

Supported flags from benchcmp
-----
**prettybenchcmp** support all flags from [benchcmp](https://godoc.org/golang.org/x/tools/cmd/benchcmp).

1. -mag
2. -best
3. -changed

Check [benchcmp documentation](https://godoc.org/golang.org/x/tools/cmd/benchcmp) for details.

How it work 
-----

1. If *.benchHistory* doesn't exist. 
 * It create *.benchHistory*. Do first benchmark and write it to *.benchHistory*.

2. If *.benchHistory* exist but doesn't exist in git. 
 * It clean *.benchHistory* and write current benchmark

3. If *.benchHistory* has history
 * It parse *.benchHistory*.
 * It check that file has a hash of previous commit which contain *.benchHistory* changes.
 * If **true**
   * it truncate part of file which contain previous benchmark
 * else
   * it write to file: a separator, previous hash commit which contain *.benchHistory* changes, current benchmark

Format of .benchHistory
-----

```
yoshkarola 3235f14078c462b38c4b79912c1e34c868d34049

PASS
BenchmarkUnquoteEasy	10000000	       182 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard	 1000000	      1119 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/benchcmp2	3.141s
```

Where **```yoshkarola```** is separator.

**```3235f14078c462b38c4b79912c1e34c868d34049```** is commit hash where file *.benchHistory* was changed last time.

Further is a standard output of ```go test -bench=. -benchmem```
```
PASS
BenchmarkUnquoteEasy	10000000	       182 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard	 1000000	      1119 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/benchcmp2	3.141s
```

Authors
-------

* Petr Lozhkin <im7mortal@gmail.com>

License
-------

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
