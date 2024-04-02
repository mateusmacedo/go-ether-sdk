package model

type Post struct {
	id      int
	version int
	title   string
	content string
	author  *Author
}

type PostOption func(*Post) error

func NewPost(opts ...PostOption) *Post {
	p := &Post{}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

func (p *Post) ID() int {
	return p.id
}

func (p *Post) Version() int {
	return p.version
}

func WithPostID(id int) PostOption {
	return func(p *Post) error {
		p.id = id
		return nil
	}
}

func WithPostVersion(version int) PostOption {
	return func(p *Post) error {
		p.version = version
		return nil
	}
}

func WithPostTitle(title string) PostOption {
	return func(p *Post) error {
		p.title = title
		return nil
	}
}

func WithPostContent(content string) PostOption {
	return func(p *Post) error {
		p.content = content
		return nil
	}
}

func WithPostAuthor(author *Author) PostOption {
	return func(p *Post) error {
		p.author = author
		return nil
	}
}
