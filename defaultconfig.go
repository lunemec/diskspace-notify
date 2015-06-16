package main

import (
	"fmt"
)

const defaultConf string = `[mail]
from = nospam@test.com
sendto = nospam@test.com
;sendto = other@address.com
subject = My subject
message = Running out of space: mountpoint %%v, remaining free space %%v%%%% (%%v of %%v)

[smtp]
;auth - enable/disable SMTP Authentication
auth = True
address = smtp.test.com
username = usename@test.com
password = xxx

;port = 25
;antispamdelay = 3600  ; send another notification after x seconds

[check]
mountpoint = /
;mountpoint = /mnt/other_drive

;threshold - percent of minimum free space, notification will occur after crossing the threshold
;threshold = 10
;delay = 10  ;check every x seconds
`

// Prints default config to stdout.
func PrintDefaultConfig() {
	fmt.Printf(defaultConf)
}
