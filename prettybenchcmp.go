package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"

	"golang.org/x/tools/benchmark/parse"
)

var (
	changedOnly = flag.Bool("changed", false, "show only benchmarks that have changed\n")
	magSort     = flag.Bool("mag", false, "sort benchmarks by magnitude of change\n")
	best        = flag.Bool("best", false, "compare best times from old and new\n")
	shortFlag   = flag.Bool("short", false, `Tell long-running tests to shorten their run time.
	It is off by default but set during all.bash so that installing
	the Go tree can run a sanity check but not spend time running
	exhaustive tests.`+"\n")
	benchTimeFlag = flag.String("benchtime", "", `Run enough iterations of each benchmark to take t, specified
	as a time.Duration (for example, -benchtime 1h30s).
	The default is 1 second (1s).`+"\n")
	countFlag = flag.String("count", "", `Run each test and benchmark n times (default 1).
	If -cpu is set, run n times for each GOMAXPROCS value.
	Examples are always run once.`+"\n")
	cpuFlag = flag.String("cpu", "", `Specify a list of GOMAXPROCS values for which the tests or
	benchmarks should be executed.  The default is the current value
	of GOMAXPROCS.`+"\n")
)

// SEPARATOR contain string separator
// i doubt that somebody will use my hometown's name in name of benchmark function
const SEPARATOR = "yoshkarola"

type benchmarkObject struct {
	currentHash           string
	file                  *os.File
	buffer                *bufio.ReadWriter
	fileSize              int64
	currentBenchmark      *bytes.Buffer
	lastBenchmark         *bytes.Buffer
	isItInitialization    bool
	fileDoesNotExistInGit bool
	wasNotBeforeCommit    bool
	truncate              int64
}

func (b *benchmarkObject) doHistoryExistInGit() {
	//http://stackoverflow.com/questions/2405305/git-how-to-tell-if-a-file-is-git-tracked-by-shell-exit-code
	cmd := exec.Command("git", "ls-files", ".benchHistory")

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {

	}
	if !strings.Contains(out.String(), ".benchHistory") {
		b.fileDoesNotExistInGit = true
	}
}

func (b *benchmarkObject) initFileSize() {
	po, _ := b.file.Stat()
	b.fileSize = po.Size()
	if b.fileSize == 0 {
		b.isItInitialization = true
	} else {
		b.isItInitialization = false
	}
}

func (b *benchmarkObject) fileExist() {
	if b.fileDoesNotExistInGit {
		if b.isItInitialization {
			os.Stdout.Write([]byte("History is inited. Created .benchHistory.\n"))
		}
		_ = b.file.Truncate(0)
		b.getCurrentBenchmark()
		b.buffer.Write([]byte(b.currentBenchmark.String()))
		b.buffer.Flush()
		b.file.Close()
		os.Exit(1)
	}
}
func (b *benchmarkObject) getLastBenchmark() {
	lines := []string{}
	tail := []string{}
	scan := bufio.NewScanner(b.buffer)
	scan.Split(bufio.ScanLines)
	for scan.Scan() {
		line := scan.Text()
		//if you run prettybenchcmp before. It had rewritten tail of benchlog
		if b.wasNotBeforeCommit {
			tail = append(tail, line)
			continue
		}
		if strings.Contains(line, SEPARATOR) {
			if strings.Contains(line, b.currentHash) {
				//it's current result. It mean that we have done already benchmarks. Get previous result
				b.wasNotBeforeCommit = true
				// "\n" from previous result
				// second one is for "\n" from end of previous-current result
				tail = append(tail, "\n"+"\n"+line)
				continue
			} else {
				// it's older result. just reset array
				lines = []string{}
				continue
			}
		}
		lines = append(lines, line)
	}
	if err := scan.Err(); err != nil {
		fatal(err.Error())
	}
	if b.wasNotBeforeCommit {
		b.truncate = int64(len(strings.Join(tail, "\n")))
	} else {
		b.truncate = 0
	}
	b.lastBenchmark = bytes.NewBufferString(strings.Join(lines, "\n"))
}

func (b *benchmarkObject) getCurrentBenchmark() {
	benchTimeValue := ""
	if len(*benchTimeFlag) > 0 {
		benchTimeValue = "-benchtime=" + *benchTimeFlag
	}
	countValue := ""
	if len(*countFlag) > 0 {
		println("Count Flag isn't impemeted rightly for now.")
		//countValue = "-count=" + *countFlag
	}
	cpuValue := ""
	if len(*cpuFlag) > 0 {
		cpuValue = "-cpu=" + *cpuFlag
	}
	shortValue := ""
	if *shortFlag {
		shortValue = "-short"
	}
	println("go", "test", "-bench=.", "-benchmem", shortValue, benchTimeValue, cpuValue, countValue)
	cmd := exec.Command("go", "test", "-bench=.", "-benchmem", shortValue, benchTimeValue, cpuValue, countValue)
	var out bytes.Buffer
	cmd.Stdout = &out
	var stdErr bytes.Buffer
	cmd.Stderr = &stdErr
	err := cmd.Run()
	if err != nil {
		println("prettybenchcmpError: Tests was failed")
		println(out.String())
		b.file.Close()
		os.Exit(1)
	}
	b.currentBenchmark = &out
}

func (b *benchmarkObject) writeBenchmarkToBenchLog() {
	b.buffer.Write([]byte("\n" + SEPARATOR + " " + b.currentHash))
	b.buffer.Write([]byte("\n\n" + b.currentBenchmark.String()))
	b.buffer.Flush()
}

var hash = make(chan string)

func main() {
	cmd := exec.Command("git", "status")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fatal("git isn't exist\n" + fmt.Sprint(err) + ": " + stderr.String())
	}
	flag.Parse()
	go getHash()
	file, err := os.OpenFile(".benchHistory", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	defer file.Close()
	if err != nil {
		fatal(err)
	}
	benchObject := benchmarkObject{
		file: file,
	}
	benchObject.buffer = NewBufioNewReadWriter(benchObject.file, benchObject.file)
	benchObject.initFileSize()
	benchObject.doHistoryExistInGit()
	benchObject.fileExist()
	benchObject.currentHash = <-hash
	benchObject.getLastBenchmark()
	benchObject.getCurrentBenchmark()
	err = file.Truncate(benchObject.fileSize - benchObject.truncate)
	if err != nil {

	}
	benchObject.writeBenchmarkToBenchLog()

	after := parseBenchmarkData(benchObject.currentBenchmark)
	before := parseBenchmarkData(benchObject.lastBenchmark)

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
func parseBenchmarkData(r io.Reader) parse.Set {
	bb, err := parse.ParseSet(r)
	if err != nil {
		fatal(err)
	}
	if *best {
		selectBest(bb)
	}
	return bb
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
		println("prettybenchcmpError: Couldn't get commit hash")
		println(out.String())
		os.Exit(1)
	}
	str := out.String()
	strA := strings.Split(str, "\n")
	hash <- strA[0]
	close(hash)
}

func NewBufioNewReadWriter(r io.Reader, w io.Writer) *bufio.ReadWriter {
	reader := bufio.NewReader(r)
	writer := bufio.NewWriter(w)
	return bufio.NewReadWriter(reader, writer)
}
