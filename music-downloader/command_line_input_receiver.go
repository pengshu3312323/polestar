package music_downloader

import (
	"bufio"
	"context"
	"log"
	"os"
)

type CommandLineInputReceiver struct {
	ch           chan *InputCommand
	inputHandler func(originalInput string) *InputCommand
}

// TODO
func simpleInputHandler(originalInput string) *InputCommand {
	return &InputCommand{}
}

func NewCommandLineInputReceiver() *CommandLineInputReceiver {
	return &CommandLineInputReceiver{
		ch:           make(chan *InputCommand),
		inputHandler: simpleInputHandler,
	}
}

func (r *CommandLineInputReceiver) Listen(ctx context.Context, cancel context.CancelFunc) chan *InputCommand {
	scanner := bufio.NewScanner(os.Stdin)
	go func() {
		defer func() {
			close(r.ch)
			cancel()
			log.Println("input receiver quit")
		}()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				if scanner.Scan() {
					input := r.inputHandler(scanner.Text())
					if input.isValid() {
						r.ch <- input
					}
				}
			}
		}
	}()
	return r.ch
}
