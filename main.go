package main

import (
	"context"

	md "polestar/music-downloader"
)

func main() {
	scheduler := md.NewMusicDownloaderScheduler(md.NewCommandLineInputReceiver())
	scheduler.Start(context.Background())
}
