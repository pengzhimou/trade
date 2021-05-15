package main

import (
	"trade/trade"
)

func main() {

	// for {
	// 	trade.RunAllExamplesMarket()
	// 	time.Sleep(time.Millisecond * 500)
	// }
	trade.RunAllExamplesMarket()

}
type Bank struct {
	sync.RWMutex
	balance map[string]float64
}

func (b *Bank) In(account string, value float64) {
	b.Lock()
	defer b.Unlock()

	v, ok := b.balance[account]
	if !ok {
		b.balance[account] = 0.0
	}

	b.balance[account] += v
}

func (b *Bank) Out(account string, value float64) error {
	b.Lock()
	defer b.Unlock()

	v, ok := b.balance[account]
	if !ok || v < value {
		return errors.New("account not enough balance")
	}

	b.balance[account] -= value
	return nil
}

func (b *Bank) Query(account string) float64 {
	b.RLock()
	defer b.RUnlock()

	v, ok := b.balance[account]
	if !ok {
		return 0.0
	}

	return v
}


package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"

	"k8s.io/klog"
)

func DecodeX(iobody io.Reader, inst interface{}) error {
	dcder := json.NewDecoder(iobody)
	err := dcder.Decode(inst)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func JsonTempHD(lk sync.Mutex, fileloc string, i interface{}) interface{} {

	// once := sync.Once{}
	// once.Do(func() {
	// 	fh, err := os.OpenFile(fileloc, 0644, os.FileMode(os.O_RDONLY))

	// 	if err != nil {
	// 		klog.Error("Openfile Error: Openfile Failed: ", err)
	// 	}
	// 	defer fh.Close()

	// })
	lk.Lock()
	fh, err := os.OpenFile(fileloc, 0644, os.FileMode(os.O_RDONLY))

	if err != nil {
		klog.Error("Openfile Error: Openfile Failed: ", err)
	}
	defer fh.Close()
	lk.Unlock()

	return i
}

func main() {
	lk := sync.Mutex{}
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			aaa := JsonTempHD(lk, "aaa.json", 11)
			fmt.Println(aaa)
			wg.Done()
		}()
	}
	wg.Wait()
}

