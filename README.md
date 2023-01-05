# Failed because cache
Try to run test.
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