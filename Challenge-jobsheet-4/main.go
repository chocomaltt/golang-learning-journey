package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup){
	defer func (start time.Time) {
		fmt.Printf("Worker with ID %d process time: %s \n", id, time.Since(start))
	}(time.Now())
	defer wg.Done()
	
	for j := range jobs {
		fmt.Printf("Worker with ID %d start to work...\n", id)
		results <- j * j
	}
}

func main() {
	/* Task 1 */
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
	var wg sync.WaitGroup
	defer cancel()
	
	// heavyProcess()
	for i := 1; i <=5; i++{
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}
	
	for j := 1; j<=100; j++{
		jobs <- j
	}
	close(jobs)
	
	wg.Wait()
	close(results)
	/* End Task 1 */
	
	/* Task 2 */
	select{
	case <-time.After(3000 * time.Millisecond):
		for r := range results {
			fmt.Println("Result: ", r)
		}
	case <-ctx.Done():
		fmt.Println("Cancelled: ", ctx.Err())
	}
	
	/* End Task 2 */
	
	/* Task 3 */
	// Mutex approach
	var counterOne int
	var mu sync.Mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			counterOne++
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final counterOne result (mutex method): ", counterOne)
	
	// Channel approach
	var counterTwo int
	ch := make(chan int, 1000)
	for i := 0; i < 1000; i++{
		wg.Add(1)
		go func(){
			ch <- 1
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for v := range ch {
		counterTwo += v
	}

	fmt.Println("Final counterTwo result (channel method): ", counterTwo)
	/* End Task 3 */ 	
}