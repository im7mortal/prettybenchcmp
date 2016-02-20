package main

import (
	"bytes"
	"reflect"
	"testing"

	"golang.org/x/tools/benchmark/parse"
)

const emptyString = ""
const sameCommitHash = "3235f14078c462b38c4b79912c1e34c868d34049"
const otherCommitHash = "67f90910d610546ce1a4be8f971409178c63de5a"

const firstRecord = `PASS
BenchmarkUnquoteEasy	10000000	       182 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard	 1000000	      1117 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/benchcmp2	3.146s` + "\n"

const separator = "\nyoshkarola " + sameCommitHash + "\n\n"

const secondRecord = `PASS
BenchmarkUnquoteEasy	10000000	       190 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard	 1000000	      1140 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/benchcmp2	3.146s` + "\n"

var oneRecord = bytes.NewBufferString(firstRecord)

var coupleOfRecords = firstRecord + separator + secondRecord

// i added "\n" in the end because "go test -bench=." return result with "\n" in the end
var currentResult = `PASS
BenchmarkUnquoteEasy	10000000	       185 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnquoteHard	 2000000	       949 ns/op	     192 B/op	       2 allocs/op
ok  	github.com/im7mortal/prettybenchcmp	4.927s` + "\n"

var expectedWhenNew = coupleOfRecords + "\nyoshkarola " + otherCommitHash + "\n\n" + currentResult
var expectedWhenRepeat = firstRecord + separator + currentResult

func TestParseBenchHistoryFirstTime(t *testing.T) {
	testInstance := benchmarkObject{}
	testInstance.currentHash = otherCommitHash
	testInstance.buffer = NewBufioNewReadWriter(bytes.NewBufferString(coupleOfRecords), bytes.NewBuffer([]byte{}))
	testInstance.getLastBenchmark()
	if testInstance.lastBenchmark.String()+"\n" != "\n"+secondRecord {
		t.Errorf("We got \n======================\n" + testInstance.lastBenchmark.String() + "\n======================\n" +
			"but expected \n======================\n" + "\n" + secondRecord + "\n======================\n")
	}

	// 0 it length of string which have to be truncate from file
	if testInstance.truncate != 0 {
		t.Errorf("In this case we expected testInstance.truncate = 0. As we have new benchmark\n")
	}

	// simulation of os.file.Truncate()
	truncateResult := coupleOfRecords[:len(coupleOfRecords)-int(testInstance.truncate)]
	currentBuffer := bytes.NewBufferString(truncateResult)

	testInstance.buffer.Writer.Reset(currentBuffer)

	testInstance.currentBenchmark = bytes.NewBufferString(currentResult)

	testInstance.writeBenchmarkToBenchHistory()

	result := currentBuffer.String()
	if result != expectedWhenNew {
		t.Errorf("We got \n======================\n" + result + "\n======================\n" +
			"but expected \n======================\n" + expectedWhenNew + "\n======================\n")
	}
}

func TestParseBenchHistorySecondTime(t *testing.T) {
	testInstance := benchmarkObject{}
	testInstance.currentHash = sameCommitHash
	testInstance.buffer = NewBufioNewReadWriter(bytes.NewBufferString(coupleOfRecords), bytes.NewBuffer([]byte{}))
	testInstance.getLastBenchmark()
	if testInstance.lastBenchmark.String() != firstRecord {
		t.Errorf("We got \n======================\n" + testInstance.lastBenchmark.String() + "\n======================\n" +
			"but expected \n======================\n" + firstRecord + "\n======================\n")
	}
	// 260 it length of string which have to be truncate from file
	println(testInstance.truncate)
	if testInstance.truncate != 262 {
		t.Errorf("In this case we expected testInstance.truncate = 260. As we already had bencmarcs before commit\n")
	}

	// simulation of os.file.Truncate()
	truncateResult := coupleOfRecords[:len(coupleOfRecords)-int(testInstance.truncate)]
	currentBuffer := bytes.NewBufferString(truncateResult)

	testInstance.buffer.Writer.Reset(currentBuffer)

	testInstance.currentBenchmark = bytes.NewBufferString(currentResult)

	testInstance.writeBenchmarkToBenchHistory()

	result := currentBuffer.String()
	if result != expectedWhenRepeat {
		t.Errorf("We got \n======================\n" + result + "\n======================\n" +
			"but expected \n======================\n" + expectedWhenRepeat + "\n======================\n")
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
