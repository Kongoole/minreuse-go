package service

import (
	"os"

	"github.com/gorilla/sessions"
)

type SessionManager struct {
	Store              *sessions.FilesystemStore
	DefaultSessionName string
}

var s *SessionManager

func NewSessionManager() *SessionManager {
	once.Do(func() {
		s = &SessionManager{
			Store:              sessions.NewFilesystemStore("storage"+string(os.PathSeparator)+"sessions", []byte("asdfasdf8a9s9fajsfjqwiejoiasdf")),
			DefaultSessionName: "minreuse",
		}
	})
	return s
}
