thread
======

Simple threading library for Go. `thread` allows you to ensure that certain
functions are called from certain OS threads.


Examples
========

The following example ensures that `foo` and `bar` are both called from the
same thread.

```go
t := thread.NewThread()
go t.Run()
t.Do(foo)
t.Do(bar)
```

By calling `Run` directly from the main thread, all tasks get executed on this
thread.

```go
t := thread.NewThread()
go func() {
  t.Do(foo)
  t.Do(bar)
}()
t.Run()
```

To stop the thread, use `Stop`. If you do not want to wait for the task to
complete, you can call `Do` from a goroutine:

```go
go t.Do(foo)
```
