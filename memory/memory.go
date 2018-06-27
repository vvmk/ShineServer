package memory

import (
	"container/list"
	"sync"
	"time"

	"github.com/vvmk/bounce/session"
)

var pder = &Provider{
	list: list.New(),
}

type SessionStore struct {
	sid          string
	timeAccessed time.Time
	value        map[interface{}]interface{}
}

func (st *SessionStore) Set(key, value interface{}) error {
	st.value[key] = value
	pder.SessionUpdate(st.sid)
	return nil
}

func (st *SessionStore) Get(key interface{}) interface{} {
	pder.SessionUpdate(st.sid)
	if v, ok := st.value[key]; ok {
		return v
	} else {
		return nil
	}
	return nil
}

func (st *SessionStore) Delete(key interface{}) error {
	delete(st.value, key)
	pder.SessionUpdate(st.sid)
	return nil
}

// SessionId returns the session id as a string
func (st *SessionStore) SessionId() string {
	return st.sid
}

// Provider is an in memory implementation of the Provider interface
type Provider struct {
	lock     sync.Mutex
	sessions map[string]*list.Element
	list     *list.List
}

// SessionInit creates a new session with the provided sid
func (pder *Provider) SessionInit(sid string) (session.Session, error) {
	pder.lock.Lock()
	defer pder.lock.Unlock()

	v := make(map[interface{}]interface{}, 0)

	newsess := &SessionStore{
		sid:          sid,
		timeAccessed: time.Now(),
		value:        v,
	}
	element := pder.list.PushBack(newsess)
	pder.sessions[sid] = element

	return newsess, nil
}

// SessionRead returns the session for sid or creates a new one
func (pder *Provider) SessionRead(sid string) (session.Session, error) {
	if element, ok := pder.sessions[sid]; ok {
		return element.Value.(*SessionStore), nil
	} else {
		sess, err := pder.SessionInit(sid)
		return sess, err
	}

	return nil, nil
}

// SessionDestroy deletes the session at sid if it exists
func (pder *Provider) SessionDestroy(sid string) error {
	if element, ok := pder.sessions[sid]; ok {
		delete(pder.sessions, sid)
		pder.list.Remove(element)
		return nil
	}

	return nil
}

// SessionGC is called by the sessionManager's GC() to destroy any
// expired sessions
func (pder *Provider) SessionGC(maxLife int64) {
	pder.lock.Lock()
	defer pder.lock.Unlock()

	for {
		element := pder.list.Back()
		if element == nil {
			break
		}
		if (element.Value.(*SessionStore).timeAccessed.Unix() + maxLife) < time.Now().Unix() {
			pder.list.Remove(element)
			delete(pder.sessions, element.Value.(*SessionStore).sid)
		} else {
			break
		}
	}
}

// SessionUpdate sets the sessions time accessed to now
func (pder *Provider) SessionUpdate(sid string) error {
	pder.lock.Lock()
	defer pder.lock.Unlock()

	if element, ok := pder.sessions[sid]; ok {
		element.Value.(*SessionStore).timeAccessed = time.Now()
		pder.list.MoveToFront(element)

		return nil
	}

	return nil
}

// init creates the session map and registers with the session manager
func init() {
	pder.sessions = make(map[string]*list.Element, 0)
	session.Register("memory", pder)
}
