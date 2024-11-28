package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/codeation/impress/joint/bus"
	"github.com/codeation/impress/joint/drawrecv"
	"github.com/codeation/impress/joint/drawwait"
	"github.com/codeation/impress/joint/eventsend"
	"github.com/codeation/itlog/gtk4"
	"github.com/codeation/itlog/uiapi4"
)

func init() {
	runtime.LockOSThread()
}

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
	dom := uiapi4.New(e)
	s := drawwait.NewStreamCommand(dom, client.StreamFile())
	d := drawrecv.NewSyncCommand(dom, client.StreamPipe, client.SyncPipe)

	var streamIO *gtk4.WatchIO
	streamIO, err = gtk4.NewStreamIO(client.StreamFile(), func() {
		if err := s.StreamCommand(); err != nil {
			if errors.Is(err, drawwait.ErrPipeClosing) {
				streamIO.Done()
				return
			}
			log.Println(err)
		}
	})
	if err != nil {
		return fmt.Errorf("gtk.NewStreamIO: %w", err)
	}

	var requestIO *gtk4.WatchIO
	requestIO, err = gtk4.NewRequestIO(client.RequestFile(), func() {
		if err := d.SyncCommand(); err != nil {
			if errors.Is(err, drawrecv.ErrPipeClosing) {
				requestIO.Done()
				return
			}
			log.Println(err)
		}
	})
	if err != nil {
		return fmt.Errorf("gtk.NewRequestIO: %w", err)
	}

	dom.Run()

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
