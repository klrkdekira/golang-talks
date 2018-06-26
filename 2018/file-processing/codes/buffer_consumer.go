import "sync"

func consume(in <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for item := range in {
		// just consume, do nothing
	}
}
