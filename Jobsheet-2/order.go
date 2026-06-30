package main

type Audit struct{ CreatedBy string }

func (a Audit) Who() string { return a.CreatedBy }