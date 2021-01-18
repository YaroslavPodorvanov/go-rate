# go-rate

### Benchmark
```bash
go test ./... -v -bench=. -benchmem -cpu=1,2,4,8
```
```text
BenchmarkMapMutex_Allow       	 9700345	       104 ns/op	       0 B/op	       0 allocs/op
BenchmarkMapMutex_Allow-2     	 8311795	       135 ns/op	       0 B/op	       0 allocs/op
BenchmarkMapMutex_Allow-4     	 6601027	       176 ns/op	       0 B/op	       0 allocs/op
BenchmarkMapMutex_Allow-8     	 5966455	       187 ns/op	       0 B/op	       0 allocs/op

BenchmarkArrayMutex_Allow     	20913630	        56.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkArrayMutex_Allow-2   	18118332	        65.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkArrayMutex_Allow-4   	14067346	        85.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkArrayMutex_Allow-8   	12035473	        97.6 ns/op	       0 B/op	       0 allocs/op
```
```bash
go test ./... -v -bench=. -benchmem -benchtime=10s -cpu=1,2,4,8
```
```text
BenchmarkMapMutex_Allow       	97256818	       107 ns/op	       0 B/op	       0 allocs/op
BenchmarkMapMutex_Allow-2     	83677227	       125 ns/op	       0 B/op	       0 allocs/op
BenchmarkMapMutex_Allow-4     	67290454	       159 ns/op	       0 B/op	       0 allocs/op
BenchmarkMapMutex_Allow-8     	59825700	       185 ns/op	       0 B/op	       0 allocs/op

BenchmarkArrayMutex_Allow     	207221521	        57.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkArrayMutex_Allow-2   	181120644	        70.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkArrayMutex_Allow-4   	154642262	        79.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkArrayMutex_Allow-8   	123674211	       100 ns/op	       0 B/op	       0 allocs/op
```