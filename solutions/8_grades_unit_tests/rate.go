package main

type Rate struct {
	subject string
	mark    int
}

func NewRate(subject string, mark int) *Rate {
	return &Rate{
		subject: subject,
		mark:    mark,
	}
}
