package music_downloader

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
