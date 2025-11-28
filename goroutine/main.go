package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

)

type CheckResult struct {
	url    string
	status string
}

func main() {


	urls := []string{
		"https://www.google.com",
		"https://www.yahoo.com",
		"https://www.bing.com",
		"https://www.duckduckgo.com",
		"https://www.yahoo.com",
		"https://www.bing.com",
		"https://www.duckduckgo.com",
		"https://www.yahoo.com",
		"https://www.bing.com",
	}

	var workersWG sync.WaitGroup
	var displayWG sync.WaitGroup

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	checkResult := make(chan CheckResult, len(urls))
	displayWG.Add(1)
	go displayResult(&displayWG, checkResult)

	for _, url := range urls {
		workersWG.Add(1)
		go checkUrl(ctx, url, &workersWG, checkResult)

	}

	// fmt.Printf("Before worker wait\n")
	// 	workersWG.Wait()
	// 	fmt.Printf("After worker wait\n")
	// 	close(checkResult)
	// 	fmt.Println("Before display wait123")
	// 	displayWG.Wait()
	// 	fmt.Printf("After display wait\n")

	go func() {
		workersWG.Wait()
		close(checkResult)

	}()
	fmt.Println("Before display wait123")
	displayWG.Wait()

}

func checkUrl(ctx context.Context, url string, wg *sync.WaitGroup, checkResult chan CheckResult) {
	defer wg.Done()
	//resp, err := http.Get(url)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		checkResult <- CheckResult{url: url, status: "DOWN"}
		return
	}
	resp, err1 := http.DefaultClient.Do(req)
	//time.Sleep(2 * time.Second)
	if err1 != nil {
		checkResult <- CheckResult{url: url, status: "DOWN"}
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode == 200 {

		checkResult <- CheckResult{url: url, status: "UP"}
		fmt.Println("result added to channel1")
	} else {
		checkResult <- CheckResult{url: url, status: "DOWN"}
		fmt.Println("result added to channel2")

	}
}

func displayResult(wg *sync.WaitGroup, checkResult <-chan CheckResult) {
	defer wg.Done()
	fmt.Println("Before display wait")
	//time.Sleep(2 * time.Second)
	for result := range checkResult {
		fmt.Println("inside for loop")
		if result.status == "DOWN" {
			fmt.Printf("❌ %s is DOWN\n", result.url)
		} else {
			fmt.Printf("✅ %s is UP\n", result.url)
		}

	}

	fmt.Println("\nAll checks completed!")

}
