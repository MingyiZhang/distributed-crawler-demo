package engine

type ConcurrentEngine struct {
  Scheduler        Scheduler
  WorkerCount      int
  ItemChan         chan Item
  RequestProcessor Processor
}

type Processor func(Request) (ParseResult, error)

type Scheduler interface {
  ReadyNotifier
  Submit(Request)
  WorkerChan() chan Request
  Run()
}

type ReadyNotifier interface {
  WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

  out := make(chan ParseResult)
  e.Scheduler.Run()

  for i := 0; i < e.WorkerCount; i++ {
    e.createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
  }

  for _, r := range seeds {
    if isDuplicate(r.Url) {
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
      if isDuplicate(request.Url) {
        continue
      }
      e.Scheduler.Submit(request)
    }
  }
}

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

// TODO: Use separate service to check duplicated request
var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
  if visitedUrls[url] {
    return true
  }
  visitedUrls[url] = true
  return false
}
