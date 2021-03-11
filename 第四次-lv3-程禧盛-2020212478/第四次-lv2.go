
package main

import (
	"fmt"
	"sync"
	"time"
)

/**
*ticker只要定义完成，从此刻开始计时，不需要任何其他的操作，每隔固定时间都会触发。
*timer定时器，是到固定时间后会执行一次
*如果timer定时器要每隔间隔的时间执行，实现ticker的效果，使用 func (t *Timer) Reset(d Duration) bool
 */
func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	//NewTicker 返回一个新的 Ticker，该 Ticker 包含一个通道字段，并会每隔时间段 d 就向该通道发送当时的时间。它会调
	//整时间间隔或者丢弃 tick 信息以适应反应慢的接收者。如果d <= 0会触发panic。关闭该 Ticker 可
	//以释放相关资源。
	ticker1 := time.NewTicker(1 * time.Hour)
	go func(t *time.Ticker) {
		for {
			<-t.C
			fmt.Println("芜湖！起飞！")
		}
	}(ticker1)
	go func() {
		now:=time.Now()
		t:=time.Date(now.Year(),now.Month(),now.Day(),2,0,0,0,time.Local)
		if t.Before(now){
			t=t.Add(24*time.Hour)
		}
		timer:=time.NewTimer(t.Sub(now))
		for  {
			<-timer.C
			fmt.Println("没有困难的工作，只有勇敢的打工人！")
			timer.Reset(24*time.Hour)
		}
	}()
	go func() {
		now:=time.Now()
		t:=time.Date(now.Year(),now.Month(),now.Day(),8,0,0,0,time.Local)
		if t.Before(now){
			t=t.Add(24*time.Hour)
		}
		timer:=time.NewTimer(t.Sub(now))
		for  {
			<-timer.C
			fmt.Println("早安，打工人！")
			timer.Reset(24*time.Hour)
		}
	}()
	wg.Wait()
}