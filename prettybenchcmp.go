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
//	"io"
	"strings"
	"io"
)

var (
	changedOnly = flag.Bool("changed", false, "show only benchmarks that have changed")
	magSort     = flag.Bool("mag", false, "sort benchmarks by magnitude of change")
)

const usageFooter = `
Each input file should be from:
	go test -run=NONE -bench=. > [old,new].txt

Benchcmp compares old and new for each benchmark.

If -test.benchmem=true is added to the "go test" command
benchcmp will also compare memory allocations.
`
var global string

var hash = make(chan string)

func main() {
	file, err := os.OpenFile(".benchHistory", os.O_RDWR | os.O_APPEND | os.O_CREATE, 0777) // todo if not exist
	if err != nil {
		fatal(err)
	}
	if !doHistoryExistInGit() {
		_ = file.Truncate(0)
	}
	po, _ := file.Stat()
	if po.Size() == 0 {
		result := getCurrentResult().Bytes()
		os.Stdout.Write([]byte("History is inited. Created .benchHistory."))
		os.Stdout.Write(result)
		_,_ = file.Write(result)
		file.Close()
		return
	}

	go getHash2()


	scan := make([]byte, 4096)
	_, _ = file.Read(scan)
	str := string(scan)


	currentHash := <-hash
	bool2 := strings.Contains(str, "separator")
	var strA []string
	if bool2 {
		strA = strings.Split(str, "separator")
	} else {
		strA = strings.Split(str, "PASS")
	}


	bool1 := strings.Contains(str, currentHash)

	var yu *bytes.Buffer
	if bool1 {
		yu = bytes.NewBufferString(strA[len(strA) - 2])
	} else {
		yu = bytes.NewBufferString(strA[len(strA) - 1])
	}



	after := parsePipe()
	before := parseFile(yu)
	defer file.Close()
	cmps, warnings := Correlate(before, after)


	if bool1 {
		_ = file.Truncate(0.)
		 for i , l := range strA{
			if i == len(strA) - 2 {
				_,_ =  file.Write([]byte(l[:len(l) - 2]))
				break
			} else {
				_,_ =  file.Write([]byte(l + "separator"))
			}
		 }

	}





	file.Write([]byte(currentHash))
	file.Write([]byte("\n\n"+ global))








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

func parseFile(r io.Reader) (parse.Set) {

	bb, err := parse.ParseSet(r)
	if err != nil {
		fatal(err)
	}
	return bb
}

func parsePipe() parse.Set {
	r := getCurrentResult()
	bb, err := parse.ParseSet(r)
	if err != nil {
		fatal(err)
	}
	return bb
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

func getCurrentResult() *bytes.Buffer {
	cmd := exec.Command("go", "test", "-bench=.", "-benchmem")
	cmd.Env = []string{"GOPATH=/home/peter/gocode"}
	var out bytes.Buffer
	cmd.Stdout = &out
	var stderr bytes.Buffer
	cmd.Stdout = &out
//	_, _ = os.Stdout.Write([]byte("Start benchmarks\n"))
//	_, _ = os.Stdout.Write([]byte("************************************************************\n"))
//	cmd.Stdout = os.Stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
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





func getHash2() {
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
	hash <- "\n\nseparator " + strA[0]
	close(hash)
}