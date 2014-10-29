Diskspace-notifier
==================

Checks periodically for free disk space and alerts user via email when there is not enough free space (threshold set in config).


Note: This program is still in development and should not be used (yet).


Goals
-----
* Check free disk space each X seconds (configurable).
* When free disk space crosses threshold (configurable in %) send notification.
* Since checking uses Statfs, it should be async.
* Sending email should not be async, we could get in trouble - starting next cycle before the previous email was sent for example.
* It should keep timestamp of last email notification and keep it in mind in the next cycle (not to SPAM).
* It should either be able to log to a file and run as daemon, or log to stdout and run normally.
