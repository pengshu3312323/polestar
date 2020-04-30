package baidu

import (
	"errors"
	"fmt"

	md "polestar/music-downloader"
)

type BaiduMusicDownloader struct {
	downloadCh chan *md.SongDownload
}

func NewBaiduMusicDownloader(ch chan *md.SongDownload) *BaiduMusicDownloader {
	return &BaiduMusicDownloader{
		downloadCh: ch,
	}
}

func (dl *BaiduMusicDownloader) Name() string {
	return "baidu-downloader"
}

// TODO
func (dl *BaiduMusicDownloader) Search(songName string) []*md.SongDetail {
	fmt.Printf("searching %s ...", songName)
	return nil
}

// TODO
func (dl *BaiduMusicDownloader) Download(songID string) (*md.SongDownload, error) {
	return nil, errors.New("not finished")
}

func (dl *BaiduMusicDownloader) ReadyToStore(song *md.SongDownload) {
	dl.downloadCh <- song
}
