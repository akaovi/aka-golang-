package util

import (
	"fmt"
)

type Bar struct {
	percent int64  
	cur     int64  
	total   int64  
	rate    string 
	graph   string 
}

func (b *Bar) NewOption(start, total int64) {
	b.cur = start
	b.total = total
	if b.graph == "" {
		b.graph = "█"
	}
	b.percent = b.GetPercent()
}

func (b *Bar) GetPercent() int64 {
	return int64(float32(b.cur) / float32(b.total) * 100)
}

func (b *Bar) Play(cur int64) {
	b.cur = cur
	last := b.percent
	b.percent = b.GetPercent()
	if b.percent != last && b.percent%2 == 0 {
		b.rate += b.graph
	}
	fmt.Printf("\r|%s] 已下载%d%%    当前：%d/%d", b.rate, b.percent, b.cur, b.total)
}
