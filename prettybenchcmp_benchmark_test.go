package main

import (
	"bytes"
	"testing"
)

func BenchmarkGetLastBenchmark3651(b *testing.B) {
	b.ResetTimer()
	buffer := bytes.NewBufferString(history3651)
	benchObject := benchmarkObject{}
	benchObject.buffer = NewBufioNewReadWriter(buffer, buffer)
	benchObject.currentHash = "b960e287b86f7018b43ff81582912ccce89b3678"
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		benchObject.getLastBenchmark()
	}
}

func BenchmarkGetLastBenchmark24456(b *testing.B) {
	b.ResetTimer()
	buffer := bytes.NewBufferString(history24456)
	benchObject := benchmarkObject{}
	benchObject.buffer = NewBufioNewReadWriter(buffer, buffer)
	benchObject.currentHash = "b960e287b86f7018b43ff81582912ccce89b3678"
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		benchObject.getLastBenchmark()
	}
}


var history3651 = `PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s
`

var history24456 = `PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola 09d84a852e08a8fe267aacad58a2e44c678c75a9

PASS
BenchmarkUnquoteEasy-4	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 2000000	       943 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.914s

yoshkarola b1156bff84efe3d293c533ffc5989e27ed4e9a68

PASS
BenchmarkUnquoteEasy-4	10000000	       210 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 1000000	      1270 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	3.600s

yoshkarola 36f0cf34919f6c9eef9562a8d41079482d3cc67f

PASS
BenchmarkUnquoteEasy-4	10000000	       192 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 1000000	      1163 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	4.431s

yoshkarola 7ba0f93e4bc43f7f2bc2afad1bc16785e861b238

PASS
BenchmarkUnquoteEasy-4	10000000	       216 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 1000000	      1377 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	3.765s

yoshkarola 4512d2385a1ca74cd7c7810a2d4780eaf80de1df

PASS
BenchmarkUnquoteEasy-4	10000000	       213 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 1000000	      1277 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	3.645s

yoshkarola fb5182078f0d3afbf3333e6bbcf6b7355d86a057

PASS
BenchmarkUnquoteEasy-4	10000000	       241 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 1000000	      1321 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	3.981s

yoshkarola 13e734f22348f71fc7f2c92d63df3dea6b34cb4d

PASS
BenchmarkUnquoteEasy-4	10000000	       212 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 1000000	      1255 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	3.604s

yoshkarola 676b76774e427d543427f71e9199b37708227675

PASS
BenchmarkUnquoteEasy-4	10000000	       195 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard-4	 1000000	      1111 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp/testDirectory	3.284s
`
