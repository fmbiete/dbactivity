# DB Activity

Terminal UI monitor for database activity - Linux only

## Build
```
go build cmd/main.go
```

## OS Metrics
### CPU Usage
- Contention: number of runnable processes vs CPU count
- Number Processes Running
- Number Processes Blocked
- Total %: as CPU count-percentage
- IO Wait %: as CPU count-percentage
- Steal %: as CPU count-percentage
- System %: as CPU count-percentage
- User %: as CPU count-percentage
- Idle %: as CPU count-percentage
### Memory Usage
- Total GB: total memory usable after kernel reservations
- Used GB: memory marked as used
- Available GB: memory marked as available
- Free GB: memory marked as not referenced
- Cache GB: memory used for buffer-cache purposes
- Inactive (file) GB: memory used for file-cache purposes that can be reclaimed
- Dirty GB: modified memory that needs to be written to some device
- Swap Used GB: swap space used
- Swap Free GB: swap space not used
### IO Usage
- Read MB/s
- Write MB/S
- Total MB/s
### Network Usage
- Read MB/s
- Write MB/s
- Total MB/s

## Database Metrics
### Database Activity
Active Time in ms
Idle in Transaction Time in ms
Deadlocks per second
Transactions per second
Tuples inserted per second
Tuples updated per second
Tuples deleted per second
Tuples reads per second
Tuples returned to client per second
### Database IO
Blocks Disk read per second: it can be physical disk read or OS buffer-cache reads
Blocks Shared Memory read per second: reads served from database shared buffers
Blocks Read Time/s: amount of time is spent reading blocks in a second (it can be greater than 1 for multi-core systems)
Blocks Write Time/s: amount of time is spent writing blocks in a second (it can be greater than 1 for multi-core systems)
Temp Files/s: number of temp files created per second
Temp Bytes/s: size of the temp files created per second
### Database Pool
Blocked
Active
Waiting: sessions actively waiting for some lock or resource (doesn't include idle waits)
Idle in Transaction: sessions idle in transaction or idle in transaction aborted
Disabled: sessions disabled
Fastpath: sessions performing fastpath operations
Idle
Total