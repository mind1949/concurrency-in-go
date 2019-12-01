## 使用sync.Con修改前
```go
// it would consume all cycles of one core
for conditionTrue() == false {}

// 低效的, 并且不得不去考虑要sleep多长时间
for conditionTrue() == false {
    time.Sleep(1*time.Millisecond)
}
```
## 使用sync.Con修改后
```go
c := sync.NewCond(&sync.Mutex{})
c.L.Lock()
for conditionTrue() == false {
	c.Wait()
}
c.L.Unlock()
```