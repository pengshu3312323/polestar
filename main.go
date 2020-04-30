package main

import (
	"context"

	md "polestar/music-downloader"
	"polestar/music-downloader/baidu"
)

func main() {
	cctx, cancel := context.WithCancel(context.Background())
	fileHandler := md.GetLocalDirFileHandler()
	downloadCh := fileHandler.Consume(cctx, cancel)
	scheduler := md.NewMusicDownloaderScheduler(md.NewCommandLineInputReceiver())
	scheduler.RegisterMusicDownloader(baidu.NewBaiduMusicDownloader(downloadCh))
	scheduler.Start(cctx, cancel)
}
