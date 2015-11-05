# wut
wut: *W*hat's this *U*nix *T*ime. Simple command line converter from unix timestamps to local time.

It is basically the same as running `date -r <unixtime>` on the Mac or `date -d '1970-01-01 <unixtime> sec UTC'` on Linux, according to [this StackExchange comment](http://stackoverflow.com/a/14107355/3754075).

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

