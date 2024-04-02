package model

type Comment struct {
	id      int
	version int
	post    string
	author  *Author
	content string
}

type CommentOption func(*Comment) error

func NewComment(opts ...CommentOption) *Comment {
	c := &Comment{}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *Comment) ID() int {
	return c.id
}

func (c *Comment) Version() int {
	return c.version
}

func (c *Comment) Post() string {
	return c.post
}

func (c *Comment) Author() *Author {
	return c.author
}

func (c *Comment) Content() string {
	return c.content
}

func WithCommentID(id int) CommentOption {
	return func(c *Comment) error {
		c.id = id
		return nil
	}
}

func WithCommentVersion(version int) CommentOption {
	return func(c *Comment) error {
		c.version = version
		return nil
	}
}

func WithCommentPost(post string) CommentOption {
	return func(c *Comment) error {
		c.post = post
		return nil
	}
}

func WithCommentAuthor(author *Author) CommentOption {
	return func(c *Comment) error {
		c.author = author
		return nil
	}
}

func WithCommentContent(content string) CommentOption {
	return func(c *Comment) error {
		c.content = content
		return nil
	}
}
