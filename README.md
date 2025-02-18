- 100% coverage
- zero memory allocations
- fuzz tests

```bash
$ go test -bench=. -benchmem .
goos: darwin
goarch: arm64
pkg: github.com/ndx-technologies/mm-go/bitset
cpu: Apple M3 Max
BenchmarkBitSet256_IsEmpty-16                   996201979                1.081 ns/op           0 B/op          0 allocs/op
BenchmarkBitSet256_Set-16                       1000000000               1.153 ns/op           0 B/op          0 allocs/op
BenchmarkBitSet256_Get-16                       967555314                1.239 ns/op           0 B/op          0 allocs/op
BenchmarkBitSet256_Bits-16                      16789438                71.11 ns/op            0 B/op          0 allocs/op
BenchmarkBitSet256_Union-16                     790709817                1.503 ns/op           0 B/op          0 allocs/op
BenchmarkBitSet256_AppendBinary-16              85984010                11.67 ns/op          162 B/op          0 allocs/op
BenchmarkBitSet256_MarshalBinary-16             83669370                15.12 ns/op           32 B/op          1 allocs/op
BenchmarkBitSet256_UnmarshalBinary-16           252775549                4.768 ns/op           0 B/op          0 allocs/op
BenchmarkBitSet_IsEmpty-16                      967321329                1.239 ns/op           0 B/op          0 allocs/op
BenchmarkBitSet_Set/256-16                      897358120                1.336 ns/op           0 B/op          0 allocs/op
BenchmarkBitSet_Set/700-16                      1000000000               1.063 ns/op           0 B/op          0 allocs/op
BenchmarkBitSet_Get/256-16                      897290182                1.334 ns/op           0 B/op          0 allocs/op
BenchmarkBitSet_Get/700-16                      899116060                1.332 ns/op           0 B/op          0 allocs/op
BenchmarkBitSet_Bits/256-16                      6038888               198.3 ns/op             0 B/op          0 allocs/op
BenchmarkBitSet_Bits/700-16                      2214734               543.1 ns/op             0 B/op          0 allocs/op
BenchmarkBitSet_Union/256-16                    356285148                3.371 ns/op           0 B/op          0 allocs/op
BenchmarkBitSet_Union/700-16                    232129774                5.182 ns/op           0 B/op          0 allocs/op
BenchmarkBitSet_AppendBinary/256-16             123249978                9.574 ns/op         177 B/op          0 allocs/op
BenchmarkBitSet_AppendBinary/700-16             60263776                20.54 ns/op          488 B/op          0 allocs/op
BenchmarkBitSet_MarshalBinary/256-16            70763061                16.88 ns/op           32 B/op          1 allocs/op
BenchmarkBitSet_MarshalBinary/700-16            39793405                29.76 ns/op           96 B/op          1 allocs/op
BenchmarkBitSet_UnmarshalBinary/256-16          223735075                5.375 ns/op           0 B/op          0 allocs/op
BenchmarkBitSet_UnmarshalBinary/700-16          131237452                9.135 ns/op           0 B/op          0 allocs/op
PASS
ok      github.com/ndx-technologies/bitset      27.964s
```
