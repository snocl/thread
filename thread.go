// Package thread implements simple routines for forcing tasks to be executed
// on certain threads. It is useful for functions that must be run on the same
// thread, or code that must be run on the main thread (e.g. GUI code under
// OS X).
package thread

import (
    "runtime"
)

// Thread is a handle for executing code on an OS thread.
type Thread struct {
    tasks chan func()
    done  chan struct{}
}

// New creates a new Thread. The thread will not process tasks until Run is
// called.
func New() *Thread {
    return &Thread{
        // Buffering tasks should reduce the required goroutine switches for
        // single-goroutine-threaded programs.
        make(chan func(), 1),
        make(chan struct{}),
    }
}

// Run causes the current thread to execute all functions passed to Do. The
// thread might still be used randomly by other goroutines. This calls
// runtime.LockOSThread but does not unlock it again.
func (thread *Thread) Run() {
    runtime.LockOSThread()
    for task := range thread.tasks {
        task()
        thread.done <- struct{}{}
    }
}

// Do causes the task to be executed on the Thread and waits for the task to
// finish. Do panics, if 
func (thread *Thread) Do(task func()) {
    thread.tasks <- task
    <-thread.done
}

// Stop makes the Thread stop after the next task has completed. It does not
// wait for the queue to empty. To do this, you can call Do with thread.Stop.
func (thread *Thread) Stop() {
    close(thread.tasks)
}
