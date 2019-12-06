## 无限for-select循环
### 作用:
> 一边等待停止, 一边执行
### 风格一:
```go
    for {
			select {
			case <-done:
				return
			default:
			}
			// Do non-preemptable work
		}
```
### 风格二:
```go
        	for {
        			select {
        			case <-done:
        				return
        			default:
        			// Do non-preemptable work
        			}
        		}
```