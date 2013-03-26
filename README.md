## Mitosis

### Short description

Mitosis allows Go applications to easily fork themselves while preserving
arbitrary application state and inherit file descriptors.


### Longer description

Mitosis supplies a simple mechanism by which any Go application can have
itself relaunched into a new session, while allowing it to pass on arbitrary
application state to the new instance. This includes inheritance of open
file descriptors.

Any aplication wanting to hook into this service, must import this package,
and implement the API it exposes. An example of this can be seen in
`testdata/main.go`.


### Why?

I wrote this because I have several long running applications which need to
be updated occasionally. One of which is an IRC bot. Its functionality consists
of plugin modules. Because Go does not support hit-plugging of code, I am
relegated to recompiling the bot, stopping the original and relaunching it.

This naturally means it will lose its connection to the IRC server. Not only
does this generate JOIN/PART noise in the channels it occupies, it is simply
unnecessary if I am able to launch a new bot instance, hand it the existing
connection and then gracefuly shut down the original.

This is what Mitosis seeks to offer in a generalized plug-and-play fashion.


### Usage

    $ go get github.com/jteeuwen/mitosis


### License

Unless otherwise stated, all of the work in this project is subject to a
1-clause BSD license. Its contents can be found in the enclosed LICENSE file.

