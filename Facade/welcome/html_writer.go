package welcome

import "fmt"

// HTMLWriter わざと複雑にしてある
type HTMLWriter interface {
	Title(title string)
	Paragraph(msg string)
	Link(href string, caption string)
	MailTo(mailAddr string, userName string)
	Close()
}

type htmlWriter struct{}

func newHTMLWriter() HTMLWriter {
	return htmlWriter{}
}

func (htmlWriter) Title(title string) {
	fmt.Printf("<!DOCTYPE html>")
	fmt.Printf("<html>")
	fmt.Printf("<head>")
	fmt.Printf("<title>" + title + "</title>")
	fmt.Printf("</head>")
	fmt.Printf("<body>")
	fmt.Printf("\n")
	fmt.Printf("<h1>" + title + "</h1>")
	fmt.Printf("\n")
}

func (htmlWriter) Paragraph(msg string) {
	fmt.Printf("<p>" + msg + "</p>")
	fmt.Printf("\n")
}

func (h htmlWriter) Link(href string, caption string) {
	h.Paragraph("<a href=\"" + href + "\">" + caption + "</a>")
}

func (h htmlWriter) MailTo(mailAddr string, userName string) {
	h.Link(fmt.Sprintf("mailto: %s", mailAddr), userName)
}

func (w htmlWriter) Close() {
	fmt.Printf("</body>")
	fmt.Printf("</html>")
	fmt.Printf("\n")
}
