// Package thread implements simple routines for forcing tasks to be executed
// on certain threads. Especially useful for code that must be called from the
// main thread.
package thread

import (
    "runtime"
)

// Thread is a handle for an OS thread.
type Thread struct {
    tasks chan func()
    done  chan struct{}
}

// NewThread creates a new Thread. The thread will not handles calls to Do
// until Run has been called.
func New() *Thread {
    return &Thread{
        make(chan func()),
        make(chan struct{}),
    }
}

// Run causes the current thread to execute all functions passed to Do. The
// thread might still be used randomly by other goroutines.
func (thread *Thread) Run() {
    runtime.LockOSThread()
    for task := range thread.tasks {
        task()
        thread.done <- struct{}{}
    }
}

// Do causes task to be executed on the Thread and waits for the task to
// finish.
func (thread *Thread) Do(task func()) {
    thread.tasks <- task
    <-thread.done
}

// Stop makes the Thread stop after the next task has completed. It does not
// wait for the queue to empty. To do this, you can call Do with thread.Stop.
func (thread *Thread) Stop() {
    close(thread.tasks)
}
