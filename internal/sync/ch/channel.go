package ch

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func ChannelTypes() {
	biDirectCh := make(chan string)
	sendOnlyCh := make(chan<- int)
	recieveOnlyCh := make(<-chan int)

	fmt.Printf("chan T   denotes a bidirectional channel type : make(chan string)   is  %T\n", biDirectCh)
	fmt.Printf("chan<- T denotes a send-only channel type     : make(chan<- int)    is  %T\n", sendOnlyCh)
	fmt.Printf("<-chan T denotes a receive-only channel type  : make(<-chan string) is  %T\n", recieveOnlyCh)
}

func ChannelOperations() {
	ch := make(chan string, 5)
	defer close(ch)
	fmt.Println("Close oper: close(ch).")
	fmt.Printf("Cap: %d, Len: %d\n", cap(ch), len(ch))

	ch <- "hello"
	fmt.Printf("\nSend oper: ch <- 'hello'.\n")
	fmt.Printf("Cap: %d, Len: %d\n", cap(ch), len(ch))

	s := <-ch
	fmt.Printf("\nRecieve oper: <-ch: %s\n", s)
	fmt.Printf("Cap: %d, Len: %d\n", cap(ch), len(ch))
}

func printMsg(ch chan string) {
	fmt.Println(<-ch)
}

func ChannelPrintMsg() {
	ch := make(chan string)
	defer close(ch)

	go printMsg(ch)
	ch <- "Hello from goroutine 'printMsg'"
}

func eval(c chan int) {
	for i := 1; i < 10; i++ {
		c <- i * i
	}

	close(c)
}

func ChannelForLoopEvaluation() {
	ch := make(chan int)

	go eval(ch)

	// we can use for-range loop for automatic checking channel for closed state
	for {
		val, ok := <-ch
		if !ok {
			log.Println("the channel is closed and no more read operations can be performed")
			break
		}
		log.Println(val)
	}
}

func ChannelTalking() {
	log.Println("init c channel")
	c := make(chan int) // an unbuffered channel
	go func(ch chan<- int, x int) {
		time.Sleep(time.Second)
		// Send the value and block until the result is received.
		log.Printf("send %d to c channel", x*x)
		ch <- x * x
		log.Printf("after send a value to c channel")
	}(c, 4)

	log.Println("init done channel")
	done := make(chan struct{})
	go func(ch <-chan int) {
		// Block until value is received.
		n := <-ch
		log.Printf("recieve %d from c channel", n)
		time.Sleep(time.Second)
		log.Printf("send empty struct to done channel")
		done <- struct{}{}
		log.Printf("after send to done channel")
	}(c)
	// Block here until a value is received by
	// the channel "done".
	log.Printf("block until a value is received by the channel done")
	<-done
	fmt.Println("end")
}

func ChannelBuffered() {
	log.Println("init c channel")
	c := make(chan int, 2)
	log.Println("send 1 to c channel")
	c <- 1
	log.Println("send 2 to c channel")
	c <- 2
	log.Println("close c channel")
	close(c)

	printCh := func(len, cap int, n string) {
		log.Printf("%s channel: [ len: %d, cap: %d ]", n, len, cap)
	}
	printRecieve := func(v int, b bool, ch string) {
		log.Printf("recieve from %s channel: [ %v, %v ]", ch, v, b)
	}
	printCh(len(c), cap(c), "c") // 2 2
	x, ok := <-c
	printRecieve(x, ok, "c")     // 1 true
	printCh(len(c), cap(c), "c") // 1 2

	x, ok = <-c
	printRecieve(x, ok, "c")     // 2 true
	printCh(len(c), cap(c), "c") // 0 2

	x, ok = <-c
	printRecieve(x, ok, "c")     // 0 false
	printCh(len(c), cap(c), "c") // 0 2
}

func ChannelNeverEndingAction() {
	ball := make(chan string)
	kickBall := func(player string) {
		for {
			log.Println(<-ball, "kicked the ball.")
			time.Sleep(time.Second)
			ball <- player
		}
	}
	go kickBall("John")
	go kickBall("Alice")
	go kickBall("Bob")
	go kickBall("Emily")
	ball <- "referee"
	// var c chan bool // nil
	// <-c             // blocking for ever
}

func longTimeRequestReturnReceiveOnly() <-chan int32 {
	c := make(chan int32)

	go func() {
		// Simulate a workload.
		time.Sleep(time.Second * 3)
		c <- rand.Int31n(100)
	}()

	return c
}

func ChannelReturnReceiveOnlyResult() {
	rand.Seed(time.Now().UnixNano())

	sum := func(a, b int32) int32 {
		return a + b
	}

	a, b := longTimeRequestReturnReceiveOnly(), longTimeRequestReturnReceiveOnly()
	fmt.Println(sum(<-a, <-b))
}

func longTimeRequestReceiveSendOnly(c chan<- int32) {
	// Simulate a workload.
	time.Sleep(time.Second * 3)
	c <- rand.Int31n(100)
}

func ChannelPassSendOnly() {
	rand.Seed(time.Now().UnixNano())

	sum := func(a, b int32) int32 {
		return a + b
	}

	a, b := make(chan int32), make(chan int32)
	go longTimeRequestReceiveSendOnly(a)
	go longTimeRequestReceiveSendOnly(b)
	fmt.Println(sum(<-a, <-b))
}

// First response wins
// A piece of data can be received from several sources to avoid high latencies.
// For a lot of factors, the response durations of these sources may vary much.
// Even for a specified source, its response durations are also not constant.
// To make the response duration as short as possible, we can send a request to
// every source in a separated goroutine. Only the first response will be used,
// other slower ones will be discarded.

// If there are N sources, the capacity of the communication channel must be at least N-1,
// to avoid the goroutines corresponding the discarded responses being blocked for ever.

func source(c chan<- int32) {
	a, b := rand.Int31n(5)+1, rand.Int31n(3)+1
	// sleep 1/2/3s
	time.Sleep(time.Duration(b) * time.Second)
	c <- a
}

func ChannelFirstResponseWins() {
	rand.Seed(time.Now().UnixNano())

	startTime := time.Now()
	fmt.Println("started at: ", startTime.Format(time.RFC3339))
	c := make(chan int32, 1)
	for i := 0; i < cap(c); i++ {
		go source(c)
	}

	// Only the first response will be used.
	select {
	case res := <-c:
		fmt.Println("elapsed: ", time.Since(startTime))
		fmt.Println(res)
	}
}

func generateNumbers(total int, ch chan<- int32) {
	for i := 1; i < total; i++ {
		v := rand.Int31()
		fmt.Printf("sending %d to channel\n", v)
		ch <- v
	}
}

func printNumbers(idx int, ch <-chan int32, wg *sync.WaitGroup) {
	defer wg.Done()

	for v := range ch {
		fmt.Printf("%d: read %d from channel\n", idx, v)
	}
}

func ChannelGenerateAndPrintNumbers() {
	rand.Seed(time.Now().UnixMicro())

	var wg sync.WaitGroup

	ch := make(chan int32)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go printNumbers(i, ch, &wg)
	}

	generateNumbers(5, ch)
	close(ch)

	fmt.Println("waiting for goroutines to finish...")
	wg.Wait()
	fmt.Println("done...")
}
