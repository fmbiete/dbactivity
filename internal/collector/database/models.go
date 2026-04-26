package database

import "time"

type Stats struct {
	ConnOthers                   uint64
	ConnActive                   uint64
	ConnIdle                     uint64
	ConnIdleInTransaction        uint64
	ConnIdleInTransactionAborted uint64
	ConnFastpath                 uint64
	ConnDisabled                 uint64
	ConnBlocked                  uint64
	ConnWaiting                  uint64
	Commits                      uint64
	Rollbacks                    uint64
	ActiveTime                   float64
	IdleInTransactionTime        float64
	TempFiles                    uint64
	TempBytes                    uint64
	Deadlocks                    uint64
	BlkReadTime                  uint64
	BlkWriteTime                 uint64
	TupReturned                  uint64
	TupFetched                   uint64
	TupInserted                  uint64
	TupUpdated                   uint64
	TupDeleted                   uint64
	BlksRead                     uint64
	BlksHit                      uint64
	LastTime                     time.Time
	LastCommits                  uint64
	LastRollbacks                uint64
	LastActiveTime               float64
	LastIdleInTransactionTime    float64
	LastTempFiles                uint64
	LastTempBytes                uint64
	LastDeadlocks                uint64
	LastBlkReadTime              uint64
	LastBlkWriteTime             uint64
	LastTupReturned              uint64
	LastTupFetched               uint64
	LastTupInserted              uint64
	LastTupUpdated               uint64
	LastTupDeleted               uint64
	LastBlksRead                 uint64
	LastBlksHit                  uint64
}
