package main

type Storable interface {
	Save() error
	Load() error
}