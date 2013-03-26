## Sample mitosis client

This program shows how a client can use the mitosis service to issue automatic
relaunches with preservation of application state across sessions.

**WARNING**: This code is for demonstration purposes only.
Running it, will cause it to spawn a new instance of itself every 5 seconds.
This can only be stopped by issueing a `pkill testdata` command from a shell.
