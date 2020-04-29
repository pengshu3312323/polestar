package music_downloader

import "context"

type Action int

const (
	ACTION_NA       = Action(0)
	ACTION_SEARCH   = Action(1)
	ACTION_DOWNLOAD = Action(2)
)

type MusicDownloader interface {
	Name() string
	Search(songName string) []*SongDetail
	Download(songID string)
}

// TODO
type SongDetail struct {
	Name    string
	Singer  string
	Quality int32
}

type DownloadFileHandler interface {
	Store([]byte) error
}

type SongDownload struct {
	SongName string
	SongFile []byte
}

type InputCommand struct {
	Name   string
	Action Action
}

type inputReceiver interface {
	Listen(ctx context.Context) chan *InputCommand
}
