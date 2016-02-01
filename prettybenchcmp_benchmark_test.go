package main

import (
	"testing"
	"os"
	"bufio"
)

func BenchmarkGetLastBenchmark3651(b *testing.B) {
	b.ResetTimer()
	file, err := os.OpenFile("testDirectory/.benchHistory3651", os.O_RDONLY, 0777)
	defer file.Close()
	if err != nil {
		b.Fatal(err)
	}
	benchObject := benchmarkObject{
		file: file,
	}
	benchObject.buffer = bufio.NewReader(benchObject.file)
	benchObject.currentHash = "b960e287b86f7018b43ff81582912ccce89b3678"
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		benchObject.getLastBenchmark()
	}
}

func BenchmarkGetLastBenchmark24456(b *testing.B) {
	b.ResetTimer()
	file, err := os.OpenFile("testDirectory/.benchHistory24456", os.O_RDONLY, 0777)
	defer file.Close()
	if err != nil {
		b.Fatal(err)
	}
	benchObject := benchmarkObject{
		file: file,
	}
	benchObject.buffer = bufio.NewReader(benchObject.file)
	benchObject.currentHash = "b960e287b86f7018b43ff81582912ccce89b3678"
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		benchObject.getLastBenchmark()
	}
}
