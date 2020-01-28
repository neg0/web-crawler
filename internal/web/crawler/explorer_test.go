package crawler

import (
	"testing"
	"time"

	"cuvva/internal/web/page"
)

func TestNewCrawlerClient(t *testing.T) {
	sut := NewExplorerClient()
	respCh := make(chan []*page.Page)

	go func() {
		resp, err := sut.Handle("https://cuvva.com", 10)
		if err != nil {
			t.Log(err)
		}
		respCh <- resp
		close(respCh)
	}()

	for {
		select {
		case res := <-respCh:
			t.Log(res)
			return
		case <-time.After(time.Second * 300):
			println("timeout!")
			return
		}
	}
}
