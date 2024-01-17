import(
	"sync"
	"testing"
)
func TestCounter(t *testing.T) {
	var counter int64
	var wg sync.WaitGroup
	for i := 0; i < 64; i++ {
		go func() {
			for i := 0; i < 10000; i++ {
				counter++
			}
			wg.Done()
		}()
	}
	wg.Wait()
	if counter != 64000000 {
		t.Errorf("counter should be 64000000,but got %d", counter)
	}
}