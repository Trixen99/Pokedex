package main

type config struct {
	next     string
	previous string
}

func initconfig() config {
	var cfg config
	return cfg
}
