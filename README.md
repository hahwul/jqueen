# jqueen

## Usage
```
   __    ____                                *
   \ \  /___ \_   _  ___  ___ _ __         * | *
    \ \//  / / | | |/ _ \/ _ \ '_ \       * \|/ *
 /\_/ / \_/ /| |_| |  __/  __/ | | | * * * \|O|/ * * *
 \___/\___,_\ \__,_|\___|\___|_| |_|  \o\o\o|O|o/o/o/
  v0.1.0  by @hahwul                  (<><><>O<><><>)

jqueen is simple job queue manager

Usage:
  jqueen [command]

Available Commands:
  add         Adding job
  daemon      managing jqueen daemon
  help        Help about any command
  list        Show job lists
  rm          Remove job
  show        Show detail information of job

Flags:
      --config string   config file (default is $HOME/.jqueen.yaml)
  -h, --help            help for jqueen
      --host string     daemon mode: binding address / client mode: connect address (default "localhost")
      --port int        daemon mode: listen port / client mode: connect port (default 6886)
```

## Install 
```
$ go get -u github.com/hahwul/jqueen
$ ~/go/bin/jqueen
```
