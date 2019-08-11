package LeakyBucket

import (
	"time"
)

//漏斗

type  LeakyBucket struct {
	Size uint64
	Fill float64
	LeakInterval time.Duration //多长时间流出一个
	LastUpdate time.Time
	Now func() time.Time //更新现在的时间
}

func NewLeakyBucket(size uint64, leakInterval time.Duration)  *LeakyBucket{
	return &LeakyBucket{
		Size:size,
		Fill:0,
		LeakInterval:leakInterval,
		Now:time.Now,
		LastUpdate:time.Now(),
	}
}

func (b *LeakyBucket) updateFill() {
	now := b.Now()
	if b.Fill > 0 {
		elapsed := now.Sub(b.LastUpdate)
		leakOut := float64(elapsed) /float64(b.LeakInterval)
		b.Fill -= leakOut

		if b.Fill < 0 {
			b.Fill = 0
		}
	}

	b.LastUpdate = now
}

func (b *LeakyBucket) Pour(amount uint16) bool  {
	b.updateFill()

	var newfill float64 = b.Fill + float64(amount)
	if newfill > float64(b.Size) {
		return false
	}

	b.Fill = newfill

	return true
}


