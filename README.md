# PebbleBus-Backend

Source code for the backend of [PebbleBus][].

[pebblebus]: https://github.com/moosingin3space/PebbleBus

## Prerequisites

You will need [Go][] 1.4 or above installed. [gvm][] is a helpful tool for managing
your `GOPATH` variable.

[go]: http://golang.org
[gvm]: https://github.com/moovweb/gvm


## Running

To start a web server for the application, run:

    source paths.bash
    go build app
    ./app

Alternatively, to run on App Engine, run:

    source paths.bash
    goapp serve .

## License

Copyright © 2015 PebbleBus Devs
