Diskspace-notifier
==================
Checks periodically for free disk space and alerts user via email when there is not enough free space (threshold set in config).

It does not aim to replace any large monitoring tools, it is an excercise to learn **Go** and maybe create something useful.

Features
--------
* Super simple and lightweight.
* Configurable plaintext email messages.
* Check free disk space each X seconds (configurable).
* When free disk space crosses threshold (configurable in %), send notification.
* Asynchronous free space checking.
* Sends email only after all mountpoints are checked.
* Sends email only once in X seconds (configurable).
* Can log to stdout or to logfile.


Usage
-----

    # normal
    ./diskspace-notify -config="/path/to.conf"

    # with logging
    ./diskspace-notify -config="/path/to.conf" -log="/path/to.log"

Note: It cannot daemonize itself, you should use some startup manager (init.d or supervisor).

If you want to see how the email looks like, you can set **threshold** value in your config to 99 to make the program send you an email.


**Ideas are welcome!**


Get the code
------------

    go get github.com/lunemec/diskspace-notify/diskspace-notify

or

    git clone git@github.com:lunemec/diskspace-notify.git


Build
-----

    go build -o diskspace-notify src/*.go
