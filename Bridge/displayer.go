package main

/*
Displayerという名前は微妙だが、DisplayにするとCountDisplayに継承的な動きをする際に
`countDisplay does not implement CountDisplay (missing Display method)`
というコンパイルエラーが発生するため、このような命名にした。
*/

/*
Goで継承っぽいことをやる時はこのような書き方でよいのか？
そもそも、Goで継承っぽいことをやるべきなのか？
*/

type Displayer interface {
	open()
	print()
	close()
	Display()
}

type displayer struct {
	displayImpl DisplayerImpl
}

func NewDisplayer(displayImpl DisplayerImpl) Displayer {
	return displayer{displayImpl: displayImpl}
}

func (d displayer) open() {
	d.displayImpl.RawOpen()
}

func (d displayer) print() {
	d.displayImpl.RawPrint()
}

func (d displayer) close() {
	d.displayImpl.RawClose()
}

func (d displayer) Display() {
	d.open()
	d.print()
	d.close()
}

type CountDisplayer interface {
	Displayer
	MultiDisplay(times int)
}

type countDisplayer struct {
	Displayer
	displayImpl DisplayerImpl
}

func NewCountDisplayer(displayImpl DisplayerImpl) CountDisplayer {
	return countDisplayer{
		Displayer:   NewDisplayer(displayImpl),
		displayImpl: displayImpl,
	}
}

func (c countDisplayer) MultiDisplay(times int) {
	c.open()
	for i := 0; i < times; i++ {
		c.print()
	}
	c.close()
}
