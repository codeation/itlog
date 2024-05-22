package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/codeation/impress/joint/bus"
	"github.com/codeation/impress/joint/drawrecv"
	"github.com/codeation/impress/joint/eventsend"
	"github.com/codeation/itlog/gtk"
	"github.com/codeation/itlog/uiapi"
)

func run() error {
	if len(os.Args) < 2 {
		return errors.New("bus suffix not found")
	}
	client, err := bus.NewClient(os.Args[1])
	if err != nil {
		return fmt.Errorf("bus.NewClient: %w", err)
	}
	if err := client.Connect(); err != nil {
		return fmt.Errorf("client.Connect: %w", err)
	}
	defer client.Close()

	e := eventsend.New(client.EventPipe)
	dom := uiapi.New(e)
	d := drawrecv.NewDrawCommand(dom, client.StreamPipe, client.SyncPipe)

	var streamIO *gtk.WatchIO
	streamIO, err = gtk.NewStreamIO(client.StreamFile(), func() {
		if err := d.StreamCommand(); err != nil {
			if errors.Is(err, drawrecv.ErrPipeClosing) {
				streamIO.Done()
				return
			}
			log.Println(err)
		}
	})
	if err != nil {
		return err
	}

	var requestIO *gtk.WatchIO
	requestIO, err = gtk.NewRequestIO(client.RequestFile(), func() {
		if err := d.SyncCommand(); err != nil {
			if errors.Is(err, drawrecv.ErrPipeClosing) {
				requestIO.Done()
				return
			}
			log.Println(err)
		}
	})
	if err != nil {
		return err
	}

	dom.Run()

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
