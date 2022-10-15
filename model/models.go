package model

import (
	_ "github.com/lib/pq"
)

var tables = []interface{}{
	new(User),
}
