/*

The prettybenchcmp command displays performance changes between benchmarks.

prettybenchcmp parses the output of two 'go test' benchmark, one from local history,
other one from current 'go test' output,
correlates the results per benchmark, and displays the deltas.

There are standard tool benchcm. Different is that prettybenchcmp don't use temp files
and keep story in local file.

prettybenchcmp will summarize and display the performance changes,
in a format like this:

	$ benchcmp old.txt new.txt
	benchmark           old ns/op     new ns/op     delta
	BenchmarkConcat     523           68.6          -86.88%

	benchmark           old allocs     new allocs     delta
	BenchmarkConcat     3              1              -66.67%

	benchmark           old bytes     new bytes     delta
	BenchmarkConcat     80            48            -40.00%

Also it will write current result to local file .benchHistory
*/
package main // import "github.com/im7mortal/prettybenchcmp"
