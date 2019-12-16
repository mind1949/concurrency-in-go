package main

/*
// page: 119
loop:
	for {
		select {
		case <-done:
			break loop
		case mayVal, ok := <-myChan:
			if ok == false {
				return // or maybe break from for
			}
			// Do something with val
		}
	}

// page: 120
orDone := func(done, c <-chan interface{}) <-chan interface{} {
		valStream := make(chan interface{})
		go func() {
			defer close(valStream)
			select {
			case <-done:
				return
			case v, ok := <-c:
				if ok == false {
					return
				}
				select {
				case valStream <- v:
				case <-done:
				}
			}
		}()
		return valStream
	}
for val := range orDone(done, myChan) {

}
*/
