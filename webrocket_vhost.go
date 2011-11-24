// Copyright 2011 Chris Kowalik (nu7hatch). All rights reserved.
// Use of this source code is governed by a BSD-style license that
// can be found in the LICENSE file.
//
// Package webrocket implements advanced WebSocket server with custom
// protocols support. 
package webrocket

import (
	"net/http"
	"io"
	"os"
	"log"
	"errors"
	"websocket"
)

// Vhost is an namespaced, standalone handler for websocket
// connections. Each vhost has it's own users and permission
// management setting, independent channels, etc.
type Vhost struct {
	Log         *log.Logger
	path        string
	isRunning   bool
	handler     websocket.Handler
	users       map[string]*User
	connections map[string]*conn
	channels    map[string]*channel
	codec       websocket.Codec
	frontAPI    websocketAPI
}

// Returns new vhost configured to handle websocket connections.
func NewVhost(path string) *Vhost {
	v := &Vhost{path: path, isRunning: true}
	v.handler = websocket.Handler(func(ws *websocket.Conn) { v.handle(ws) })
	v.users = make(map[string]*User)
	v.connections = make(map[string]*conn)
	v.channels = make(map[string]*channel)
	v.codec = websocket.JSON
	v.Log = log.New(os.Stderr, "", log.LstdFlags)
	return v
}

// Prepares new connection to enter in the event loop.
func (v *Vhost) handle(ws *websocket.Conn) {
	c := wrapConn(ws, v)
	v.connections[c.token] = c
	v.eventLoop(c)
	v.cleanup(c)
}

// cleanup removes all subscprionts and other relations
// between closed connection and the system.
func (v *Vhost) cleanup(c *conn) {
	c.unsubscribeAll()
	delete(v.connections, c.token)
}

// eventLoop maintains main loop for handled connection.
func (v *Vhost) eventLoop(c *conn) {
	for {
		if !v.IsRunning() {
			return
		}
		var recv map[string]interface{}
		err := v.codec.Receive(c.Conn, &recv)
		if err != nil {
			if err == io.EOF {
				return
			}
			// TODO: show INVALID_PAYLOAD error
			continue
		}
		message, err := NewMessage(recv)
		if err != nil {
			// TODO: 
		}
		keepgoing, _ := v.frontAPI.Dispatch(c, message)
		if !keepgoing {
			return
		}
	}
}

// Stop closes all connection handled by this vhost and stops
// its eventLoop.
func (v *Vhost) Stop() {
	v.isRunning = false
}

// Is this vhost running?
func (v *Vhost) IsRunning() bool {
	return v.isRunning
}

// Returns list of active connections.
func (v *Vhost) Connections() map[string]*conn {
	return v.connections
}

// AddUser configures new user account within this vhost.
func (v *Vhost) AddUser(name, secret string, permission int) error {
	if len(name) == 0 {
		return errors.New("User name can't be blank")
	}
	_, ok := v.users[name]
	if ok {
		return errors.New("User already exists")
	}
	if permission == 0 {
		return errors.New("Invalid permissions")
	}
	v.users[name] = NewUser(name, secret, permission)
	v.Log.Printf("vhost[%s]: ADD_USER name='%s' permission=%d", v.path, name, permission)
	return nil
}

// DeleteUser deletes user account with given name.
func (v *Vhost) DeleteUser(name string) error {
	_, ok := v.users[name]
	if !ok {
		return errors.New("User doesn't exist")
	}
	delete(v.users, name)
	v.Log.Printf("vhost[%s]: DELETE_USER name='%s'", v.path, name)
	return nil
}

// SetUserPermissions configures user access.
func (v *Vhost) SetUserPermissions(name string, permission int) error {
	user, ok := v.users[name]
	if !ok {
		return errors.New("User doesn't exist")
	}
	if permission == 0 {
		return errors.New("Invalid permissions")
	}
	user.Permission = permission
	v.Log.Printf("vhost[%s]: SET_USER_PERMISSION name='%s' permission=%d", v.path, name, permission)
	return nil
}

// Returns list of configured user accounts.
func (v *Vhost) Users() map[string]*User {
	return v.users
}

// OpenChannel creates new channel ready to subscribe.
func (v *Vhost) CreateChannel(name string) *channel {
	channel := newChannel(v, name)
	v.channels[name] = channel
	v.Log.Printf("vhost[%s]: CREATE_CHANNEL name='%s'", v.path, name)
	return channel
}

// Returns specified channel.
func (v *Vhost) GetChannel(name string) (*channel, bool) {
	channel, ok := v.channels[name]
	return channel, ok
}

// Returns specified channel. If channel doesn't exist, then will be
// created automatically.
func (v *Vhost) GetOrCreateChannel(name string) *channel {
	channel, ok := v.channels[name]
	if !ok {
		return v.CreateChannel(name)
	}
	return channel
}

// Returns list of used channels.
func (v *Vhost) Channels() map[string]*channel {
	return v.channels
}

// Returns specified user.
func (v *Vhost) GetUser(name string) (*User, bool) {
	user, ok := v.users[name]
	return user, ok
}

// ServeHTTP extends standard websocket.Handler implementation
// of http.Handler interface.
func (v *Vhost) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if v.isRunning {
		v.handler.ServeHTTP(w, req)
	}
}