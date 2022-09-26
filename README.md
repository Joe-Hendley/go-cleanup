# go-cleanup

This displays the difference in order of operation between t.Cleanup & defer

run go test -v and observe the order of the logs

```
❯ go test -v
=== RUN   TestDefer
Start Test
Helper Called!
Something that needs deferring called!
Helper Called!
Defer Called!
Helper Cleanup Called!
Helper Cleanup Called!
--- PASS: TestDefer (0.00s)
=== RUN   TestCleanup
Start Test
Helper Called!
Something that needs cleaning up called!
Helper Called!
Helper Cleanup Called!
Cleanup Called!
Helper Cleanup Called!
--- PASS: TestCleanup (0.00s)
PASS
ok      github.com/Joe-Hendley/go-cleanup       0.001s
```
