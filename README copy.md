- 100% coverage
- zero memory allocations
- fuzz tests

```bash
$ go test -bench=. -benchmem .
goos: darwin
goarch: arm64
pkg: github.com/ndx-technologies/mm-go/bitset
cpu: Apple M3 Max
BenchmarkBitSet256_IsEmpty-16                   983588416                1.082 ns/op           0 B/op          0 allocs/op
BenchmarkBitSet256_Set-16                       904824694                1.330 ns/op           0 B/op          0 allocs/op
BenchmarkBitSet256_Get-16                       965786023                1.243 ns/op           0 B/op          0 allocs/op
BenchmarkBitSet256_Bits-16                      16734543                71.16 ns/op            0 B/op          0 allocs/op
BenchmarkBitSet256_Union-16                     795979417                1.507 ns/op           0 B/op          0 allocs/op
BenchmarkBitSet256_AppendBinary-16              79466908                14.13 ns/op          176 B/op          0 allocs/op
BenchmarkBitSet256_MarshalBinary-16             27732892                43.51 ns/op           56 B/op          3 allocs/op
BenchmarkBitSet256_UnmarshalBinary-16           247623842                4.813 ns/op           0 B/op          0 allocs/op
BenchmarkBitSet_IsEmpty-16                      964625696                1.241 ns/op           0 B/op          0 allocs/op
BenchmarkBitSet_Set/256-16                      1000000000               1.065 ns/op           0 B/op          0 allocs/op
BenchmarkBitSet_Set/700-16                      901208745                1.331 ns/op           0 B/op          0 allocs/op
BenchmarkBitSet_Get/256-16                      901088911                1.333 ns/op           0 B/op          0 allocs/op
BenchmarkBitSet_Get/700-16                      899410047                1.332 ns/op           0 B/op          0 allocs/op
BenchmarkBitSet_Bits/256-16                      5888877               204.2 ns/op             0 B/op          0 allocs/op
BenchmarkBitSet_Bits/700-16                      2158359               554.5 ns/op             0 B/op          0 allocs/op
BenchmarkBitSet_Union/256-16                    329113932                3.605 ns/op           0 B/op          0 allocs/op
BenchmarkBitSet_Union/700-16                    228257520                5.259 ns/op           0 B/op          0 allocs/op
BenchmarkBitSet_AppendBinary/256-16             138147462                8.826 ns/op         198 B/op          0 allocs/op
BenchmarkBitSet_AppendBinary/700-16             57850143                19.98 ns/op          508 B/op          0 allocs/op
BenchmarkBitSet_MarshalBinary/256-16            25962153                45.88 ns/op           56 B/op          3 allocs/op
BenchmarkBitSet_MarshalBinary/700-16            11136928               108.3 ns/op           248 B/op          5 allocs/op
BenchmarkBitSet_UnmarshalBinary/256-16          233613890                5.129 ns/op           0 B/op          0 allocs/op
BenchmarkBitSet_UnmarshalBinary/700-16          134039142                9.003 ns/op           0 B/op          0 allocs/op
PASS
ok      github.com/ndx-technologies/mm-go/bitset        27.864s
```
