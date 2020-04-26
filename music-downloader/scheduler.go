package music_downloader

import "context"

type MusicDownloaderScheduler struct {
	RegisteredDownloader map[string]MusicDownloader
}

func NewMusicDownloaderScheduler() *MusicDownloaderScheduler {
	return &MusicDownloaderScheduler{
		RegisteredDownloader: make(map[string]MusicDownloader),
	}
}

func (scheduler *MusicDownloaderScheduler) Start(ctx context.Context) {
}

func (scheduler *MusicDownloaderScheduler) RegisterMusicDownloader(downloader MusicDownloader) {
	scheduler.RegisteredDownloader[downloader.Name()] = downloader
}
