package main

type PrintBanner struct {
	Banner
}

func NewPrintBanner(value string) PrintBanner {
	return PrintBanner{Banner{value: value}}
}

func (p PrintBanner) PrintWeak() {
	p.Banner.ShowWithParen()
}

func (p PrintBanner) PrintStrong() {
	p.Banner.ShowWithAster()
}
