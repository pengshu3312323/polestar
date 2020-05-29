package music_downloader

import (
	"context"
	"fmt"
	"testing"
)

func TestHttpInputService_Listen(t *testing.T) {
	s := NewHttpInputService()
	cctx, cancel := context.WithCancel(context.Background())
	ch := s.Listen(cctx, cancel)
	for i := range ch {
		fmt.Println(i)
	}
}
