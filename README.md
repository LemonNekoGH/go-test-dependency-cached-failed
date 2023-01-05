# Failed because cache
Try to run test singly.
```bash
go test -run ^TestA$ gotestdependencycachedfailed && go test -run ^TestB$ gotestdependencycachedfailed
```
You will see both test passed.  
Try to run all test in one process.
```bash
$ go test .
```
You will see the test failed.  
Try to remove comment in `go-test-dependency-cached-failed_test.go`
```diff
- 15    // aCached = []int{}
+ 15    aCached = []int{}
```
And run again, test passed.