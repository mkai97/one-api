package common

import (
	"github.com/songquanpeng/one-api/common/env"
)

var UsingSQLite = true
var UsingPostgreSQL = false
var UsingMySQL = false

var SQLitePath = "db/one-api.db"
var SQLiteBusyTimeout = env.Int("SQLITE_BUSY_TIMEOUT", 3000)
