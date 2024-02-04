package main

type message struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	SentAt  string `json:"sent_at"`
	Sender  string `json:"sender"`
}

type user struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
