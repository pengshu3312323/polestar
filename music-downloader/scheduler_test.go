package music_downloader

import (
	"fmt"
	"testing"
)

type dummyDownloader struct {
}

func newDummyDownloader() *dummyDownloader {
	return &dummyDownloader{}
}

func (downloader *dummyDownloader) Name() string {
	return "dummy-downloader"
}

func (downloader *dummyDownloader) Search(songName string) []*SongDetail {
	fmt.Println("searching...")
	return []*SongDetail{
		{
			Name:   "a terrible song",
			Singer: "stupid singer",
		},
		{
			Name:   "another terrible song",
			Singer: "another stupid singer",
		},
	}
}

func (downloader *dummyDownloader) Download(songID string) {
	fmt.Printf("%s: downloading", songID)
}

func TestMusicDownloaderScheduler_RegisterMusicDownloader(t *testing.T) {
	scheduler := NewMusicDownloaderScheduler()
	scheduler.RegisterMusicDownloader(newDummyDownloader())
	for _, downloader := range scheduler.RegisteredDownloader {
		fmt.Println(downloader.Name())
		downloader.Search("")
		downloader.Download("test song")
	}
}
