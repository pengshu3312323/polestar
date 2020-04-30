package music_downloader

import (
	"context"
	"log"
)

type MusicDownloaderScheduler struct {
	ch                   chan *SongDetail
	RegisteredDownloader map[string]MusicDownloader
	InputReceiver        inputReceiver
}

func NewMusicDownloaderScheduler(receiver inputReceiver) *MusicDownloaderScheduler {
	return &MusicDownloaderScheduler{
		ch:                   make(chan *SongDetail),
		RegisteredDownloader: make(map[string]MusicDownloader),
		InputReceiver:        receiver,
	}
}

func (scheduler *MusicDownloaderScheduler) Start(ctx context.Context, cancel context.CancelFunc) {
	if len(scheduler.RegisteredDownloader) < 1 {
		log.Fatal("no downloader registered")
	}
	defer cancel()
	inputChan := scheduler.InputReceiver.Listen(ctx, cancel)
	for input := range inputChan {
		// todo
		log.Printf("input coming: %s", input.Name)
		for _, v := range scheduler.RegisteredDownloader {
			log.Printf("download %s on %s", input.Name, v.Name())
			res, err := v.Download(input.Name)
			if err != nil {
				log.Printf("%s download failed", input.Name)
				continue
			}
			go v.ReadyToStore(res)
		}
	}
}

func (scheduler *MusicDownloaderScheduler) RegisterMusicDownloader(downloader MusicDownloader) {
	scheduler.RegisteredDownloader[downloader.Name()] = downloader
}
