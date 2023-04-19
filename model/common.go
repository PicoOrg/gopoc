package model

type Poc struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Source      string `json:"source"`
	Finger      string `json:"finger"`
	Url         string `json:"url"`
	Host        string `json:"host"`
	Port        int    `json:"port"`
}

type Option func(*Poc)

func NewPoc(option ...Option) *Poc {
	p := &Poc{}
	for _, opt := range option {
		opt(p)
	}
	return p
}

func WithTitle(title string) Option {
	return func(p *Poc) {
		p.Title = title
	}
}

func WithDescription(description string) Option {
	return func(p *Poc) {
		p.Description = description
	}
}

func WithAuthor(author string) Option {
	return func(p *Poc) {
		p.Author = author
	}
}

func WithSource(source string) Option {
	return func(p *Poc) {
		p.Source = source
	}
}

func WithFinger(finger string) Option {
	return func(p *Poc) {
		p.Finger = finger
	}
}

func WithUrl(url string) Option {
	return func(p *Poc) {
		p.Url = url
	}
}

func WithHostPort(host string, port int) Option {
	return func(p *Poc) {
		p.Host = host
		p.Port = port
	}
}
