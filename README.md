# Benchmarks

# Groups and Pools

I set up an experiment to compare standard library groups (`sync.WaitGroup`, `errgroup.Group`) with github.com/sourcegraph/conc groups and pools. In this experiment, **the conc groups/pools appear to be a bit slower than standard library offerings**. Afterwards, we can consider whether the slowdown is worth the extra functionality.

```
➜ cd groups_and_pools
➜ go test -bench=.
goos: darwin
goarch: arm64
pkg: github.com/denis-engcom/benchmarks/groups_and_pools
BenchmarkStdWaitGroup1000-10       	    4644	    277505 ns/op
BenchmarkStdErrgroup1000-10        	    3876	    304500 ns/op
BenchmarkConcWaitGroup1000-10      	    4557	    322034 ns/op
BenchmarkConcErrorPool1000-10      	    3648	    357631 ns/op
BenchmarkConcContextPool1000-10    	    3210	    362683 ns/op
PASS
ok  	github.com/denis-engcom/benchmarks/groups_and_pools	9.389s
```
