package music_downloader

import (
	"context"
	"testing"
	"time"
)

func TestLocalDirFileHandler_Consume(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	h := GetLocalDirFileHandler()
	h.Consume(ctx, cancel)
	h.Store(&SongDownload{
		SongName: "Hurt",
		SongFile: make([]byte, 0),
	})
	go func() {
		h.Store(&SongDownload{
			SongName: "One",
			SongFile: make([]byte, 0),
		})
	}()
	time.Sleep(5 * time.Second)
	h.Store(&SongDownload{
		SongName: "So Far From The Clyde",
		SongFile: make([]byte, 0),
	})
	<-ctx.Done()
}
