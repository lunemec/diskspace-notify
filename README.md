Diskspace-notifier
==================
Checks periodically for free disk space and alerts user via email when there is not enough free space (threshold set in config).

It does not aim to replace any large monitoring tools, it was an excercise to learn **Go** and maybe create something useful.

Goals
-----
* Super simple and lightweight.
* Check free disk space each X seconds (configurable).
* When free disk space crosses threshold (configurable in %) send notification.
* Since checking uses Statfs, it should be async.
* Sending email should not be async, we could get in trouble - starting next cycle before the previous email was sent for example.
* It should keep timestamp of last email notification and keep it in mind in the next cycle (not to SPAM).
* It should either be able to log to a file or to stdout (for start managers).


Usage
-----

    # normal
    ./diskspace-notify -config="/path/to.conf"

    # with logging
    ./diskspace-notify -config="/path/to.conf" -log="/path/to.log"

Note: It cannot daemonize itself, you should use some startup manager (init.d or supervisor).

If you want to see how the email looks like, you can set **threshold** value in your config to 100 to make the program send you an email.


**Ideas are welcome!**


Get the code
------------

    go get github.com/lunemec/diskspace-notify/diskspace-notify

or

    git clone git@github.com:lunemec/diskspace-notify.git


Build
-----

    go build -o diskspace-notify src/*.go
