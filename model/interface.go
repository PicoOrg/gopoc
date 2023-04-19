package model

type IPoc interface {
	Verify() bool
	Exploit() bool
}
