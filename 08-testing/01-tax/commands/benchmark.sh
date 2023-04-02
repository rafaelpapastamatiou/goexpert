# RUN BENCHMARK ALONG TESTS
go test -bench .

# RUN ONLY BENCHMARK
go test -bench=. -run=^#

# RUN ONLY BENCHMARK (10 runs per benchmark)
go test -bench=. -run=^# -count=10

# RUN ONLY BENCHMARK (run benchmarks for 3s)
go test -bench=. -run=^# -benchtime=3s

# RUN ONLY BENCHMARK (show memory alloc)
go test -bench=. -run=^# -benchmem