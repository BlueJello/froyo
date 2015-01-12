package provider

import "time"

type SearchResult struct {
	Name string
	// Size in bytes
	Size     int64
	Files    int
	Added    time.Time
	Seeders  int
	Leachers int
}

type TorrentProvider interface {
	Search() ([]*SearchResult, error)
	Download(*SearchResult) error
	Login(username, password string) error
	LoginRequired() bool
}
