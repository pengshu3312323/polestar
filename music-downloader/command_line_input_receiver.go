package music_downloader

import (
	"bufio"
	"context"
	"fmt"
	"os"
)

type CommandLineInputReceiver struct {
	ch chan *InputCommand
}

func NewCommandLineInputReceiver() *CommandLineInputReceiver {
	return &CommandLineInputReceiver{
		ch: make(chan *InputCommand),
	}
}

func (r *CommandLineInputReceiver) Listen(ctx context.Context) chan *InputCommand {
	scanner := bufio.NewScanner(os.Stdin)
	go func() {
		defer func() {
			fmt.Println("input receiver quit")
		}()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				if scanner.Scan() {
					r.ch <- &InputCommand{
						Name: scanner.Text(),
					}
				}
			}
		}
	}()
	return r.ch
}
