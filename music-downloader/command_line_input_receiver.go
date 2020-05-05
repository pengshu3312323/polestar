package music_downloader

import (
	"bufio"
	"context"
	"log"
	"os"
	"regexp"
)

var Actions = map[string]Action{
	"-s": ACTION_SEARCH,
	"-d": ACTION_DOWNLOAD,
}

const inputReStr = `^(-[a-zA-Z]) +(.+)$`

type CommandLineInputReceiver struct {
	ch           chan *InputCommand
	inputHandler func(originalInput string) *InputCommand
}

// TODO
func simpleInputHandler(originalInput string) *InputCommand {
	re := regexp.MustCompile(inputReStr)
	match := re.FindStringSubmatch(originalInput)
	if len(match) == 3 {
		if _, ok := Actions[match[1]]; ok {
			return &InputCommand{
				Name:   match[2],
				Action: Actions[match[1]],
			}
		}
		log.Println("Err Action, please use: [-s, -d]")
	}
	log.Println("Err Input, please input '<Action> <Name>'")
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
