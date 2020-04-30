package music_downloader

import (
	"context"
	"log"
	"sync"
)

var fileHandlerOnce sync.Once
var fileHandler *LocalDirFileHandler

type LocalDirFileHandler struct {
	downloadCh chan *SongDownload
}

func GetLocalDirFileHandler() *LocalDirFileHandler {
	fileHandlerOnce.Do(func() {
		fileHandler = &LocalDirFileHandler{
			downloadCh: make(chan *SongDownload),
		}
	})
	return fileHandler
}

func NewLocalDirFileHandler(ch chan *SongDownload) *LocalDirFileHandler {
	return &LocalDirFileHandler{
		downloadCh: make(chan *SongDownload),
	}
}

func (h *LocalDirFileHandler) Store(song *SongDownload) {
	h.downloadCh <- song
}

func (h *LocalDirFileHandler) Consume(ctx context.Context, cancel context.CancelFunc) chan *SongDownload {
	go func() {
		defer func() {
			close(h.downloadCh)
			cancel()
			log.Println("file handler quit")
		}()
		for {
			select {
			case song := <-h.downloadCh:
				if song.isValid() && !h.exist(song.buildFilePath()) {
					err := h.writeOnDisk(song)
					if err != nil {
						// todo
						continue
					}
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return h.downloadCh
}

// todo
func (h *LocalDirFileHandler) writeOnDisk(song *SongDownload) error {
	log.Printf("writing %s on the disk", song.SongName)
	return nil
}

func (h *LocalDirFileHandler) exist(filePath string) bool {
	return false
}
