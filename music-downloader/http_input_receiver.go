package music_downloader

import (
	"context"
	"log"
	"net/http"
)

const address = "localhost:8080"

type HttpInputService struct {
	ch      chan *InputCommand
	service *http.ServeMux
}

func NewHttpInputService() *HttpInputService {
	return &HttpInputService{
		ch:      make(chan *InputCommand),
		service: http.NewServeMux(),
	}
}

func (s *HttpInputService) Listen(ctx context.Context, cancel context.CancelFunc) chan *InputCommand {
	go func() {
		defer func() {
			close(s.ch)
			cancel()
			log.Println("http command server quit")
		}()
		s.service.HandleFunc("/ping", s.Ping)
		s.service.HandleFunc("/search", s.search)
		s.service.HandleFunc("/download", s.download)
		if err := http.ListenAndServe(address, s.service); err != nil {
			log.Fatalf("start command server failed at %s\n, err: %s", address, err.Error())
		}
	}()
	return s.ch
}

func (s *HttpInputService) search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	song := query.Get("song")
	if len(song) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte("Give me a song")); err != nil {
			log.Printf("err: %v", err)
		}
		return
	}
	s.ch <- &InputCommand{
		Name:   song,
		Action: ACTION_SEARCH,
	}
	w.WriteHeader(http.StatusOK)
	// todo
	if _, err := w.Write([]byte("developing....")); err != nil {
		log.Printf("err: %v", err)
	}
	return
}

func (s *HttpInputService) download(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	songId := query.Get("songId")
	if len(songId) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte("Give me a songId")); err != nil {
			log.Printf("err: %v", err)
		}
		return
	}
	s.ch <- &InputCommand{
		Name:   songId,
		Action: ACTION_DOWNLOAD,
	}
	w.WriteHeader(http.StatusOK)
	// todo
	if _, err := w.Write([]byte("developing....")); err != nil {
		log.Printf("err: %v", err)
	}
	return
}

func (s *HttpInputService) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("pong")); err != nil {
		log.Println("command server is not ready: ", err)
	}
}
