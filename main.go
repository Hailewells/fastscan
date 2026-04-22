package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

func main() {
	target := flag.String("host", "127.0.0.1", "Target host")
	startPort := flag.Int("start", 1, "Start port")
	endPort := flag.Int("end", 1024, "End port")
	threads := flag.Int("threads", 50, "Number of concurrent workers")
	timeout := flag.Int("timeout", 500, "Timeout in ms")
	delay := flag.Int("delay", 10, "Delay between requests in ms")

	flag.Parse()

	scanner := Scanner{
		Timeout: time.Duration(*timeout) * time.Millisecond,
		Rate:    time.Duration(*delay) * time.Millisecond,
	}

	ports := make(chan int, *threads)
	var wg sync.WaitGroup

	// 🔥 Worker pool
	for i := 0; i < *threads; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for port := range ports {
				// Debug (optional)
				// fmt.Printf("[Worker %d] scanning port %d\n", workerID, port)

				if scanner.ScanPort(*target, port) {
					fmt.Printf("[+] Open port: %d\n", port)
				}

				time.Sleep(scanner.Rate) // rate limiting
			}
		}(i)
	}

	// 🔥 FIX: feed ports in a goroutine (prevents blocking)
	go func() {
		for port := *startPort; port <= *endPort; port++ {
			ports <- port
		}
		close(ports)
	}()

	wg.Wait()
	fmt.Println("Scan complete")
}
