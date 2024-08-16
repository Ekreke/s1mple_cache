package cache

import (
	"context"
	"sync"
	"time"
)

type cleaner struct {
	ctx      context.Context
	interval time.Duration
	done     chan struct{}
	once     sync.Once
}

func newCleaner(ctx context.Context, interval time.Duration) *cleaner {
	return &cleaner{
		ctx:      ctx,
		interval: interval,
		done:     make(chan struct{}),
	}
}

func (c *cleaner) stop() {
	// if it need stop , just close the done channel
	c.once.Do(func() { close(c.done) })
}

// run cleaner
func (c *cleaner) run(cleanup func(ctx context.Context)) {
	go func() {
		// new ticker
		ticker := time.NewTicker(c.interval)
		defer ticker.Stop()
		// if clean up get a done signal or attach the timestamp , cleaner will clean up once
		for {
			select {
			case <-ticker.C:
				cleanup(c.ctx)
				// if context get the done signal , stop the cleaner
			case <-c.ctx.Done():
				c.stop()
			case <-c.done:
				cleanup(c.ctx)
				return
			}
		}
	}()
}
