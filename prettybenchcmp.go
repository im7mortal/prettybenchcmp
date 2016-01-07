// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"text/tabwriter"

	"golang.org/x/tools/benchmark/parse"
	"os/exec"
	"bytes"
	"strings"
	"io"
)

var (
	changedOnly = flag.Bool("changed", false, "show only benchmarks that have changed")
	magSort     = flag.Bool("mag", false, "sort benchmarks by magnitude of change")
	best        = flag.Bool("best", false, "compare best times from old and new")
)

// SEPARATOR contain string separator
// i doubt that somebody will use my hometown's name in name of benchmark function
const SEPARATOR = "yoshkarola"

var global string

var hash = make(chan string)

func main() {
	cmd := exec.Command("git", "status")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fatal("git isn't exist\n" + fmt.Sprint(err) + ": " + stderr.String())
	}
	var currentResultChan chan *bytes.Buffer
	go func() {
		currentResultChan <- getCurrentResult()
	}()
	go getHash()
	file, err := os.OpenFile(".benchHistory", os.O_RDWR | os.O_APPEND | os.O_CREATE, 0777)
	defer file.Close()
	if err != nil {
		fatal(err)
	}
	if !doHistoryExistInGit() {
		_ = file.Truncate(0)
	}
	po, _ := file.Stat()
	fileSize := po.Size()
	if fileSize == 0 {
		currentResult := <-currentResultChan
		result := currentResult.Bytes()
		os.Stdout.Write([]byte("History is inited. Created .benchHistory."))
		os.Stdout.Write(result)
		_,_ = file.Write(result)
		return
	}




	currentHash := <-hash
	lastResult := getLastBenchmark(file, currentHash, fileSize)
	var yu *bytes.Buffer
	yu = bytes.NewBufferString(lastResult)
	after := parseBenchmarkData(getCurrentResult())
	before := parseBenchmarkData(yu)
	file.Write([]byte("\n" + SEPARATOR + " " + currentHash))
	file.Write([]byte("\n\n"+ global))

	cmps, warnings := Correlate(before, after)
	for _, warn := range warnings {
		fmt.Fprintln(os.Stderr, warn)
	}

	if len(cmps) == 0 {
		fatal("benchcmp: no repeated benchmarks")
	}

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 0, 5, ' ', 0)
	defer w.Flush()

	var header bool // Has the header has been displayed yet for a given block?

	if *magSort {
		sort.Sort(ByDeltaNsPerOp(cmps))
	} else {
		sort.Sort(ByParseOrder(cmps))
	}
	for _, cmp := range cmps {
		if !cmp.Measured(parse.NsPerOp) {
			continue
		}
		if delta := cmp.DeltaNsPerOp(); !*changedOnly || delta.Changed() {
			if !header {
				fmt.Fprint(w, "benchmark\told ns/op\tnew ns/op\tdelta\n")
				header = true
			}
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", cmp.Name(), formatNs(cmp.Before.NsPerOp), formatNs(cmp.After.NsPerOp), delta.Percent())
		}
	}

	header = false
	if *magSort {
		sort.Sort(ByDeltaMBPerS(cmps))
	}
	for _, cmp := range cmps {
		if !cmp.Measured(parse.MBPerS) {
			continue
		}
		if delta := cmp.DeltaMBPerS(); !*changedOnly || delta.Changed() {
			if !header {
				fmt.Fprint(w, "\nbenchmark\told MB/s\tnew MB/s\tspeedup\n")
				header = true
			}
			fmt.Fprintf(w, "%s\t%.2f\t%.2f\t%s\n", cmp.Name(), cmp.Before.MBPerS, cmp.After.MBPerS, delta.Multiple())
		}
	}

	header = false
	if *magSort {
		sort.Sort(ByDeltaAllocsPerOp(cmps))
	}
	for _, cmp := range cmps {
		if !cmp.Measured(parse.AllocsPerOp) {
			continue
		}
		if delta := cmp.DeltaAllocsPerOp(); !*changedOnly || delta.Changed() {
			if !header {
				fmt.Fprint(w, "\nbenchmark\told allocs\tnew allocs\tdelta\n")
				header = true
			}
			fmt.Fprintf(w, "%s\t%d\t%d\t%s\n", cmp.Name(), cmp.Before.AllocsPerOp, cmp.After.AllocsPerOp, delta.Percent())
		}
	}

	header = false
	if *magSort {
		sort.Sort(ByDeltaAllocedBytesPerOp(cmps))
	}
	for _, cmp := range cmps {
		if !cmp.Measured(parse.AllocedBytesPerOp) {
			continue
		}
		if delta := cmp.DeltaAllocedBytesPerOp(); !*changedOnly || delta.Changed() {
			if !header {
				fmt.Fprint(w, "\nbenchmark\told bytes\tnew bytes\tdelta\n")
				header = true
			}
			fmt.Fprintf(w, "%s\t%d\t%d\t%s\n", cmp.Name(), cmp.Before.AllocedBytesPerOp, cmp.After.AllocedBytesPerOp, cmp.DeltaAllocedBytesPerOp().Percent())
		}
	}
}

func fatal(msg interface{}) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

func selectBest(bs parse.Set) {
	for name, bb := range bs {
		if len(bb) < 2 {
			continue
		}
		ord := bb[0].Ord
		best := bb[0]
		for _, b := range bb {
			if b.NsPerOp < best.NsPerOp {
				b.Ord = ord
				best = b
			}
		}
		bs[name] = []*parse.Benchmark{best}
	}
}

// formatNs formats ns measurements to expose a useful amount of
// precision. It mirrors the ns precision logic of testing.B.
func formatNs(ns float64) string {
	prec := 0
	switch {
	case ns < 10:
		prec = 2
	case ns < 100:
		prec = 1
	}
	return strconv.FormatFloat(ns, 'f', prec, 64)
}
func parseBenchmarkData(r io.Reader) (parse.Set) {
	bb, err := parse.ParseSet(r)
	if err != nil {
		fatal(err)
	}
	if *best {
		selectBest(bb)
	}
	return bb
}

func getCurrentResult() *bytes.Buffer {
	cmd := exec.Command("go", "test", "-bench=.", "-benchmem")
	var out bytes.Buffer
	cmd.Stdout = &out
	var stdErr bytes.Buffer
	cmd.Stderr = &stdErr
	err := cmd.Run()
	if err != nil {
		fatal(fmt.Sprint(err) + ": " + stdErr.String())
	}
	global = out.String()
	return &out
}


func doHistoryExistInGit() bool {
	//http://stackoverflow.com/questions/2405305/git-how-to-tell-if-a-file-is-git-tracked-by-shell-exit-code
	cmd := exec.Command("git", "ls-files", ".benchHistory")

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {

	}
	return strings.Contains(out.String(), ".benchHistory")
}





func getHash() {
	cmd := exec.Command("git", "log", "-1", "--pretty=tformat:%H", "-p", ".benchHistory")

	var out bytes.Buffer
	cmd.Stdout = &out
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String()) //todo err from console
	}
	str := out.String()
	strA := strings.Split(str, "\n")
	hash <- strA[0]
	close(hash)
}

func getLastBenchmark(file *os.File, currentHash string, fileSize int64) string {
	results := make([]string, 1)
	var lastPart string

	for ;; {
		scan := make([]byte, 4096)
		count, err := file.Read(scan)
		if err == io.EOF {
			break
		}
		var str string
		var stringSlice []string
		str = string(scan)
		str = lastPart + str
		lenLastPart := len(lastPart)
		lastPart = ""
		isThereSeparator := strings.Contains(str, SEPARATOR)
		if count == 4096 {
			if isThereSeparator {
				stringSlice = strings.Split(str, SEPARATOR)
				lastPart = stringSlice[len(stringSlice) - 1]
				stringSlice = stringSlice[:len(stringSlice) - 1]
			} else {
				lastPart = str
				continue
			}
		} else {
			if isThereSeparator {
				temp := []byte(str)
				stringSlice = strings.Split(string(temp[: lenLastPart + count]), SEPARATOR)
			} else {
				results = append(results, str)
				continue
			}
		}
		for _, str := range stringSlice {
			results = append(results, str)
		}
	}
	lastResult := results[len(results) - 1]
	wasNotCommited := strings.Contains(lastResult, currentHash)
	if wasNotCommited {
		_ = file.Truncate(fileSize - int64(len(SEPARATOR + " " + lastResult)))
		lastResult = results[len(results) - 2]
	}
	return lastResult
}