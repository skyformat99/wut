# wut
wut: What's this Unix time. Simple command line converter from unix timestamps to local time.

Unix time is the number of seconds that have elapsed since 00:00:00 Coordinated Universal Time (UTC), Thursday, 1 January 1970, ignoring leap seconds.

## Running

`wut` without an argument uses the current Unix time. `wut` with an argument prints the Unix time given as the a local time.

    > wut 
    1446739220 is 2015-11-05 11:00:20 -0500 EST
    > wut 1000000000
    1000000000 is 2001-09-08 21:46:40 -0400 EDT
    
## Installing

Binaries are in the `bin` directory (for Mac, Linux, and Windows). You can use [gb](https://getgb.io/) to build it as well:

    > cd /path/to/wut
    > gb build all
    > bin/wut
