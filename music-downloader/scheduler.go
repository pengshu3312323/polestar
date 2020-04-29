package music_downloader

import (
	"context"
	"fmt"
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

func (scheduler *MusicDownloaderScheduler) Start(ctx context.Context) {
	cctx, cancel := context.WithCancel(ctx)
	defer cancel()
	inputChan := scheduler.InputReceiver.Listen(cctx)
	for input := range inputChan {
		fmt.Println("developing", input.Name)
	}

	// for song := range scheduler.ch {
	//
	// }
}

func (scheduler *MusicDownloaderScheduler) RegisterMusicDownloader(downloader MusicDownloader) {
	scheduler.RegisteredDownloader[downloader.Name()] = downloader
}
