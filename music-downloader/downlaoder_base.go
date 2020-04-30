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
	Download(songID string) (*SongDownload, error)
	ReadyToStore(song *SongDownload)
}

// TODO
type SongDetail struct {
	Name    string
	Singer  string
	Quality int32
}

type DownloadFileHandler interface {
	Store(song *SongDownload)
	Consume(ctx context.Context, cancel context.CancelFunc) chan *SongDownload
}

type SongDownload struct {
	SongName string
	SongFile []byte
}

// todo
func (sd *SongDownload) isValid() bool {
	return true
}

// todo
func (sd *SongDownload) buildFilePath() string {
	return ""
}

type InputCommand struct {
	Name   string
	Action Action
}

func (ic *InputCommand) isValid() bool {
	return true
}

type inputReceiver interface {
	Listen(ctx context.Context, cancel context.CancelFunc) chan *InputCommand
}
