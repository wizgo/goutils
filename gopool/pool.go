package gopool

type Go struct {
	ch chan bool
}

func New(size int) *Go {
	return &Go{ch: make(chan bool, size)}
}

func (g *Go) Acquire() {
	g.ch <- true
}

func (g *Go) Done() {
	<-g.ch
}

func (g *Go) Go(f func()) {
	go func() {
		g.Acquire()
		defer g.Done()
		f()
	}()
}

func (g *Go) Do(f func()) {
	g.Acquire()
	defer g.Done()
	f()
}
