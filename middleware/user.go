package middleware

import (
	"buzzme/store"
)

// Store holds new store connection
var Store *store.Conn

// Init ...
func Init(st *store.Conn) {
	Store = st
}
