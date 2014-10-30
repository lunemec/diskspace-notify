Diskspace-notifier
==================
Checks periodically for free disk space and alerts user via email when there is not enough free space (threshold set in config).


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


Ideas are welcome!
