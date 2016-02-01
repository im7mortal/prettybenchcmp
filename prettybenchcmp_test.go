package main

import (
	"reflect"
	"testing"
	"bytes"

	"golang.org/x/tools/benchmark/parse"
)

const emptyString = ""
const sameCommitHash = "3235f14078c462b38c4b79912c1e34c868d34049"
const otherCommitHash = "67f90910d610546ce1a4be8f971409178c63de5a"

const firstRecord = `PASS
BenchmarkUnquoteEasy	10000000	       182 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard	 1000000	      1117 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/benchcmp2	3.146s` + "\n"

const separator  = "\nyoshkarola " + sameCommitHash + "\n"

const secondRecord = "\n" + `PASS
BenchmarkUnquoteEasy	10000000	       182 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard	 1000000	      1117 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/benchcmp2	3.146s`


var oneRecord = bytes.NewBufferString(firstRecord)

var coupleOfRecords = firstRecord + separator + secondRecord

var currentResult = bytes.NewBufferString(`PASS
BenchmarkUnquoteEasy	10000000	       185 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard	 2000000	       949 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp	4.927s`)

func TestParseBenchHistoryFirstTime(t *testing.T) {
	testInstance := benchmarkObject{}
	testInstance.currentHash = otherCommitHash
	testInstance.buffer = NewBufioNewReadWriter(bytes.NewBufferString(coupleOfRecords), bytes.NewBuffer([]byte{}))
	testInstance.getLastBenchmark()
	if testInstance.lastBenchmark.String() != secondRecord {
		t.Errorf("TestParseBenchHistoryFirstTime")
	}
}

func TestParseBenchHistorySecondTime(t *testing.T) {
	testInstance := benchmarkObject{}
	testInstance.currentHash = sameCommitHash
	testInstance.buffer = NewBufioNewReadWriter(bytes.NewBufferString(coupleOfRecords), bytes.NewBuffer([]byte{}))
	testInstance.getLastBenchmark()
	if testInstance.lastBenchmark.String() != firstRecord {
		t.Errorf("TestParseBenchHistorySecondTime")
	}
}


func TestSelectBest(t *testing.T) {
	have := parse.Set{
		"Benchmark1": []*parse.Benchmark{
			{
				Name: "Benchmark1",
				N:    10, NsPerOp: 100, Measured: parse.NsPerOp,
				Ord: 0,
			},
			{
				Name: "Benchmark1",
				N:    10, NsPerOp: 50, Measured: parse.NsPerOp,
				Ord: 3,
			},
		},
		"Benchmark2": []*parse.Benchmark{
			{
				Name: "Benchmark2",
				N:    10, NsPerOp: 60, Measured: parse.NsPerOp,
				Ord: 1,
			},
			{
				Name: "Benchmark2",
				N:    10, NsPerOp: 500, Measured: parse.NsPerOp,
				Ord: 2,
			},
		},
	}

	want := parse.Set{
		"Benchmark1": []*parse.Benchmark{
			{
				Name: "Benchmark1",
				N:    10, NsPerOp: 50, Measured: parse.NsPerOp,
				Ord: 0,
			},
		},
		"Benchmark2": []*parse.Benchmark{
			{
				Name: "Benchmark2",
				N:    10, NsPerOp: 60, Measured: parse.NsPerOp,
				Ord: 1,
			},
		},
	}

	selectBest(have)
	if !reflect.DeepEqual(want, have) {
		t.Errorf("filtered bench set incorrectly, want %v have %v", want, have)
	}
}
