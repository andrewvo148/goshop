package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"sync"
	"time"
)

type Message struct {
	str  string
	wait chan bool
}

type Ball struct{ hits int }

type Job struct {
	file    string
	pattern *regexp.Regexp
	result  chan Result
}

type Result struct {
	file       string
	lineNumber int
	text       string
}

func main() {
	rand.Seed(time.Now().UnixNano()) //
	results := make(chan int32, 2)
	go longTimeRequest(results)
	go longTimeRequest(results)

	fmt.Println(sumSquares(<-results, <-results))

}


func longTimeRequest(r chan<- int32) {
		time.Sleep(time.Second * 3)
		r <- rand.Int31n(100)
}
func sumSquares(a, b int32) int32 {
	return a*a + b*b
}

func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func worker(jobs chan Job) {
	for work := range jobs {
		f, err := os.Open(work.file)
		if err != nil {
			fmt.Println(err)
			continue
		}
		scn := bufio.NewScanner(f)
		lineNumber := 1
		for scn.Scan() {
			result := work.pattern.Find(scn.Bytes())
			if len(result) > 0 {
				//fmt.Printf("%s#%d: %s\n", work.file, lineNumber, string(result))
				work.result <- Result{
					file:       work.file,
					lineNumber: lineNumber,
					text:       string(result),
				}
			}
			lineNumber++
		}
		close(work.result)
	}
}

func workerEfficient(id int, jobs <-chan int, results chan<- int) {
	//
	var wg sync.WaitGroup
	for j := range jobs {
		wg.Add(1)
		// We start a gorountine to run the job
		go func(job int) {
			fmt.Println("worker", id, "started job", job)
			time.Sleep(time.Second)
			fmt.Println("worker", id, "finished job", job)
			results <- job * 2
			wg.Done()
		}(j)
	}

	wg.Wait()

}
func player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}
func boring(msg string, quit chan bool) <-chan string {
	c := make(chan string)
	//waitForIt := make(chan bool) // Shared between all messages.
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s: %d", msg, i):
			case <-quit:
				return
			}
		}
	}()
	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()

	return c
}
