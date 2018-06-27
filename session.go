package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
)

type Manager struct {
	cookieName string
	lock       sync.Mutex
	provider   Provider
	maxLife    int64
}

// SessionInit initialize and returns a new session
// SessionRead returns the session for 'sid' or creates a new one
// SessionDestroy deletes the session for 'sid'
// SessionGC deletes expired sessions based on maxLife
type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) (Session, error)
	SessionGC(maxLifeTime int64)
}

// Only four session operations: get, set, delete, and get sid
type Session interface {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Delete(key interface{}) error
	SessionID() string
}

// create a map of provider names to Providers
var provides = make(map[string]Provider)

func NewManager(provideName, cookieName string, maxLife int64) (*Manager, error) {
	provider, ok := provides[provideName]

	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", provideName)
	}

	return &Manager{
		provider:   provider,
		cookieName: cookieName,
		maxLife:    maxLife,
	}, nil
}

// Register makes a session Provider available by the given name
// If a Register is called twice with the same name or if the driver is nil,
// it panics.
func Register(name string, provider Provider) {
	if provider == nil {
		panic("session: Register provider is nil")
	}

	if _, dup := provides[name]; dup {
		panic("session: Register called twice for provider " + name)
	}

	provides[name] = provider
}

// sessionId makes a random 32-character alphanumeric string to be used
// as a unique sid
func (manager *Manager) sessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}

	return base64.URLEncoding.EncodeToString(b)
}

func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()

	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		sid := manager.sessionId()
		session, _ = manager.provider.SessionInit(sid)
		cookie := http.Cookie{
			Name:     manager.cookieName,
			Value:    url.QueryEscape(sid),
			Path:     "/",
			HttpOnly: true,
			MaxAge:   int(manager.maxLife),
		}
		http.SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.SessionRead(sid)
	}

	return
}
