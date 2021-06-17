package model

type Quote struct {
	Id     int
	Author string
	Text   string
}

func NewQuote(author, text string) Quote {
	return Quote{
		Author: author,
		Text:   text,
	}
}

func (q *Quote) SetId(id int) {
	q.Id = id
}
