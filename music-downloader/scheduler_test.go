package music_downloader

import (
	"context"
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
	scheduler := NewMusicDownloaderScheduler(NewCommandLineInputReceiver())
	scheduler.RegisterMusicDownloader(newDummyDownloader())
	scheduler.Start(context.Background())

	for _, downloader := range scheduler.RegisteredDownloader {
		fmt.Println(downloader.Name())
		downloader.Search("")
		downloader.Download("test song")
	}
}

func TestPanicHandler(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			switch r.(type) {
			case string:
				fmt.Println("string: ", r.(string))
			case error:
				fmt.Println("error: ", r.(error).Error())
			}
		}
	}()
	testList := make([]string, 0)
	fmt.Println("start")
	// panic("panic coming!!")
	// panic(errors.New("panic coming"))
	fmt.Println(testList[0])
}
