# go-rate

### Benchmark
```bash
go test ./... -v -bench=. -benchmem -cpu=1,2,4,8
```
```text
BenchmarkMapMutex_Allow
BenchmarkMapMutex_Allow     	11476294	       102 ns/op	       0 B/op	       0 allocs/op
BenchmarkMapMutex_Allow-2   	 8419794	       120 ns/op	       0 B/op	       0 allocs/op
BenchmarkMapMutex_Allow-4   	 6760108	       173 ns/op	       0 B/op	       0 allocs/op
BenchmarkMapMutex_Allow-8   	 6015914	       195 ns/op	       0 B/op	       0 allocs/op
```