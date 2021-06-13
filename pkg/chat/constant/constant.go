package constant

import "time"

const (

	// Time allowed to read the next pong message from the peer.
	PongWait = 60 * time.Second
)

const (
	// mem
	DB_ENGINE_MEM   = 1
	DB_ENGINE_MYSQL = 2
)
