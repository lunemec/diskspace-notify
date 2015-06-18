Diskspace-notifier
==================
Checks periodically for free disk space and alerts user via email when there is not enough free space (threshold set in config).

It does not aim to replace any large monitoring tools, it is an excercise to learn **Go** and maybe create something useful.

Features
--------
* Super simple and lightweight.
* Configurable plaintext email messages.
* Check free disk space each X seconds (configurable).
* Send email via Gmail or local SMTP server.
* When free disk space crosses threshold (configurable in %), send notification.
* Asynchronous free space checking.
* Sends email only after all mountpoints are checked.
* Sends email only once in X seconds (configurable).
* Can log to stdout or to logfile.


Usage
-----

    # create config (and edit it with your information)
    ./diskspace-notify -defaultconfig > /path/to.conf

    # normal
    ./diskspace-notify -config /path/to.conf

    # with logging
    ./diskspace-notify -config /path/to.conf -log /path/to.log


Note: It cannot daemonize itself, you should use some startup manager (init.d or supervisor).

If you want to see how the email looks like, you can set **threshold** value in your config to 99 to make the program send you an email.


Run as a daemon
---------------
The program itself cannot be daemonized, but this can be achieved with system tools.
For debian-like linux machines, simply copy the binary somewhere (/www/diskspace-notify for example).

Then continue as stated above - generate config, edit it to your specifications, and test-run the program itself.
After you've finished these steps, you can copy diskspace-notify.etc to /etc/init.d/diskspace-notify.
Edit the file with the correct paths to binary, config and log.

Then you can add

    /etc/init.d/diskspace-notify start

to your /etc/rc.local (before the exit 0 statement).


**Ideas are welcome!**


Get the code
------------

    go get github.com/lunemec/diskspace-notify/diskspace-notify

or

    git clone git@github.com:lunemec/diskspace-notify.git


Build
-----
Install godep:

    go get github.com/tools/godep

Restore saved dependencies.

    godep restore

Build.

    godep go build
