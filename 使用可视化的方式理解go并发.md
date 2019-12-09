# 19年12月09号用可视化的方式理解go并发

## 背景
昨天我在尝试理解or-channel部分的or函数时遇到了如下问题, 当我尝试理解如下代码时:
```go
	var or func(
		channels ...<-chan interface{},
	) <-chan interface{}
	or = func(
		channels ...<-chan interface{},
	) <-chan interface{} {
		switch len(channels) {
		case 0:
			return nil
		case 1:
			return channels[0]
		}
		orDone := make(chan interface{})
		go func() {
			defer close(orDone)

			switch len(channels) {
			case 2:
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
			default:
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				case <-or(append(channels[3:], orDone)...):
				}
			}
		}()
		return orDone
	}
```
遇到一些理解上的困难感. 虽然我能够用文字描述大致理解整个or函数的实现思路, 但是总觉得自己思考的不够流畅, 
感觉不够连贯, 无法清除的将实现的思路在脑海里呈现出来.
但是当我无意间看到这篇文章[visualizing concurrency in go](https://divan.dev/posts/go_concurrency_visualize/), 学习了如何通过可视化的方式理解go的并发, 并尝试用文章中给出的可视化方式描述
代码中的思路后, 哪些不连贯, 不流畅的困难感消失了.
下图是我手绘的大致的可视化表示方式:
![](https://gitlab.com/mind1949.roadmap/technique.it/backend/golang/uploads/a4131adef0d5f1807eebf00184ed1729/image.png)

## 对可视化方式描述go并发的方式的描述
### 解决的问题
帮助描述go并发代码, 从而帮助人更好的思考

### 背后的方法论
图形化表示一个概念的方式比用文字描述概念的方式占用的工作记忆空间更小, 因为一旦在脑海里构建好了图形, 那么就只会占据一个chunk

### 可视化方式描述go并发的思路
* 从上往下表示时间的递增;
* 用一条垂直线表示一个goroutine;
* 用两个goroutine顶层相连的虚线表示两个goroutine之间的fork关系;
* 用两个goroutine中间相连的有向线段表示两个goroutine中通过channel完成的数据沟通;
* 用红色表示阻塞, 用绿色表示正在占用cpu;
