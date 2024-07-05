No need to create json file. Generate package will take care of it.

To test: 
    yerkebulanserikov@MacBook-Air-Yerkebulan mechtago % go run cmd/main.go                   
    Enter inputfile name to generate: 
    test1
    Enter item count:
    1000000
    Enter number of concurrent workers: 
    4
    2024/07/05 17:30:30 Finished generating successfully
    2024/07/05 17:30:31 groundTruth:  2320
    2024/07/05 17:30:31 calculated value:  2320

BenchMarks: 
    yerkebulanserikov@MacBook-Air-Yerkebulan mechtago % cd app  
    
    yerkebulanserikov@MacBook-Air-Yerkebulan app % go test -bench=BenchmarkCalculate -v -benchmem 
    goos: darwin
    goarch: arm64
    pkg: github.com/quietguido/mechtago/app
    BenchmarkCalculate
    2024/07/05 17:33:46 Finished generating successfully
    2024/07/05 17:33:49 Finished generating successfully
    BenchmarkCalculate-8           2         673824562 ns/op           23896 B/op         40 allocs/op
    PASS
    ok      github.com/quietguido/mechtago/app      8.672s

Which means 0.6738s to calculate 1.000.000 object. 
Not bad, but now great either, taking into account that https://github.com/minio/simdjson-go?ref=blog.min.io can do it 10 times better. LOL 
