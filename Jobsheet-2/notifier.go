package main

import "fmt"

type Notifier interface {
	Send(msg string) error
}

type EmailNotifier struct { From string }

func (e EmailNotifier) Send(msg string) error {
	fmt.Printf("[email dari %s] %s\n", e.From, msg)
	return nil
}

func Notify(n Notifier, msg string) error {
	return n.Send(msg)
}