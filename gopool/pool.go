package gopool

type Go struct {
	ch chan bool
}

func New(size int) *Go {
	return &Go{ch: make(chan bool, size)}
}

func (g *Go) start() {
	g.ch <- true
}

func (g *Go) done() {
	<-g.ch
}

func (g *Go) Go(f func()) {
	go func() {
		g.start()
		defer g.done()
		f()
	}()
}

func (g *Go) Do(f func()) {
	g.start()
	defer g.done()
	f()
}
