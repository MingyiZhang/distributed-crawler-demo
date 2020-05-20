package engine

import "log"

// ConcurrentEngine manages the crawler's core logic among components
type ConcurrentEngine struct {
	// Scheduler is the scheduler that managing requests and workers
	Scheduler Scheduler
	// WorkerCount is the total number of workers
	WorkerCount int
	// ItemChan is the channel of Item.
	// It gets items from parse results and pass to persist.
	ItemChan chan Item
	// RequestProcessor is a function of request which returns parse results.
	RequestProcessor Processor
	// DuplicateChecker is a function of string which check if a given url.
	DuplicateChecker Checker
}

// Processor is a function type that receives request and return parse result.
type Processor func(Request) (ParseResult, error)

// Checker is a function type that receives string and return bool.
type Checker func(string) (bool, error)

// Scheduler is an interface for worker, request scheduler
// A Scheduler with queue of workers and requests are implemented
type Scheduler interface {
	// ReadyNotifier is a notifier
	ReadyNotifier
	// Submit request to request channel
	Submit(Request)
	// Get channel of request
	WorkerChan() chan Request
	// Start Scheduler
	Run()
}

type ReadyNotifier interface {
	// WorkerReady pushes a request channel to worker channel in Scheduler
	WorkerReady(chan Request)
}

// ConcurrentEngine.Run starts the crawler engine from given seed request urls.
func (e *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		exists, err := e.DuplicateChecker(r.Url)
		if exists || err != nil {
			if err != nil {
				// TODO: need more sophisticated error handler
				log.Printf("Error: %v", err)
			}
			continue
		}
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			i := item
			go func() { e.ItemChan <- i }()
		}
		for _, request := range result.Requests {
			exists, err := e.DuplicateChecker(request.Url)
			if exists || err != nil {
				if err != nil {
					// TODO: need more sophisticated error handler
					log.Printf("Error: %v", err)
				}
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

// createWorker creates a worker goroutine to scrape web
func (e *ConcurrentEngine) createWorker(in chan Request, out chan ParseResult, r ReadyNotifier) {
	go func() {
		for {
			r.WorkerReady(in)
			request := <-in
			result, err := e.RequestProcessor(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
