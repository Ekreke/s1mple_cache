package cache

import (
	"context"
	"sync/atomic"
	"testing"
	"time"
)

func Test_cleaner(t *testing.T) {
	// new context with cancel func
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	// new cleaner
	c := newCleaner(ctx, time.Millisecond)
	doneFlag := make(chan struct{})
	c.done = doneFlag
	num := int64(0)
	c.run(func(_ context.Context) {
		atomic.AddInt64(&num, 1)
	})
	c.done <- struct{}{}
	time.Sleep(5 * time.Millisecond)
	cancelFunc()
	select {
	case <-doneFlag:
		t.Log("done")
	case <-time.After(time.Second):
		t.Fatal("timeout")
	}
	if atomic.LoadInt64(&num) < 1 {
		t.Fatalf("failed to run cleanup function, num: %d", num)
	}

}
