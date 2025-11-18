package impresslink

import (
	"image"
	"sync"

	"github.com/codeation/impress/clipboard"
	"github.com/codeation/impress/driver"
	"github.com/codeation/impress/event"
	"github.com/codeation/impress/joint/eventchan"
	"github.com/codeation/impress/joint/iface"
	"github.com/codeation/impress/joint/lazy"
	gtk "github.com/codeation/itlog/gtk4"
)

type Application struct {
	mainFn    func(d driver.Driver)
	app       *gtk.Application
	top       *gtk.TopWindow
	layout    *gtk.Layout
	menu      *gtk.Menu
	commands  chan func()
	events    <-chan event.Eventer
	callbacks iface.CallbackSet
	wg        sync.WaitGroup
}

func Exec(mainFn func(d driver.Driver)) {
	a := New(mainFn)
	defer a.wg.Wait()
}

func New(mainFn func(d driver.Driver)) *Application {
	eventChan := eventchan.New()
	a := &Application{
		mainFn:    mainFn,
		app:       gtk.NewApplication(),
		commands:  make(chan func(), 64),
		events:    eventChan.Chan(),
		callbacks: eventChan,
	}
	gtk.SignalActivate(a.onActivate)
	gtk.SignalShutdown(a.onShutdown)
	gtk.SignalDrawCallback(onDraw)
	gtk.SignalMenuItemActivateCallback(a.onItemActivate)
	gtk.SignalDelete(a.onDelete)
	gtk.SignalSizeAllocate(a.onSizeAllocate)
	gtk.SignalKeyPress(a.onKeyPress)
	gtk.SignalButtonPress(a.onButtonPress)
	gtk.SignalButtonRelease(a.onButtonPress)
	gtk.SignalMotionNotify(a.onMotionNotify)
	gtk.SignalScroll(a.onScroll)
	a.app.AppSignalConnect()
	a.wg.Go(a.app.Run)
	return a
}

func (a *Application) Init() {}

func (a *Application) onActivate() {
	a.menu = a.app.NewMenu()
	a.top = a.app.NewTopWindow()
	a.top.TopSignalConnect()
	gtk.SignalIdle(a.onIdle)
	a.wg.Go(func() { a.mainFn(lazy.New(a)) })
}

func (a *Application) onIdle() {
	for {
		select {
		case f := <-a.commands:
			f()
		default:
			return
		}
	}
}

func (a *Application) onShutdown() {
	a.app.AppSignalDisconnect()

}
func (a *Application) Done() {
	a.commands <- func() {
		a.top.TopSignalDisconnect()
		a.top.Destroy()
		a.app.Quit()
	}
}

func (a *Application) Title(title string) {
	a.commands <- func() {
		a.app.SetName(title)
	}
}

func (a *Application) Size(rect image.Rectangle) {
	a.commands <- func() {
		a.top.Size(rect.Min.X, rect.Min.Y, rect.Dx(), rect.Dy())
	}
}

func (a *Application) ClipboardGet(typeID int) {
	gtk.RequestClipboardText(a.top, a.onClipboardText)
}

func (a *Application) ClipboardPut(c clipboard.Clipboarder) {
	gtk.SetClipboardText(a.top, string(c.Data()))
}

func (a *Application) Chan() <-chan event.Eventer { return a.events }

func (a *Application) Sync() {}
