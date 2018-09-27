---
title: Administration
---

Frontier is responsible for providing an HTTP API to data in the DigitalBits network. It ingests and re-serves the data produced by the DigitalBits network in a form that is easier to consume than the performance-oriented data representations used by digitalbits-core.

## Why run frontier?

The DigitalBits Foundation runs two frontier servers, one for the public network and one for the test network, free for anyone's use at https://frontier.livenet.digitalbits.io and https://frontier.livenet.digitalbits.io.  These servers should be fine for development and small scale projects, but is not recommended that you use them for production services that need strong reliability.  By running frontier within your own infrastructure provides a number of benefits:

  - Multiple instances can be run for redundancy and scalability.
  - Request rate limiting can be disabled.
  - Full operational control without dependency on the DigitalBits Foundations operations.

## Prerequisites

Frontier is a dependent upon a digitalbits-core server.  Frontier needs access to both the SQL database and the HTTP API that is published by digitalbits-core. See [the administration guide](https://developer.digitalbits.io/digitalbits-core/software/admin.html
) to learn how to set up and administer a digitalbits-core server.  Secondly, frontier is dependent upon a postgresql server, which it uses to store processed core data for ease of use. Frontier requires postgres version >= 9.3.

In addition to the two required prerequisites above, you may optionally install a redis server to be used for rate limiting requests.

## Installing

To install frontier, you have a choice: either downloading a [prebuilt release for your target architecture](https://github.com/digitalbitsorg/go/releases) and operation system, or [building frontier yourself](#Building).  When either approach is complete, you will find yourself with a directory containing a file named `frontier`.  This file is a native binary.

After building or unpacking frontier, you simply need to copy the native binary into a directory that is part of your PATH.  Most unix-like systems have `/usr/local/bin` in PATH by default, so unless you have a preference or know better, we recommend you copy the binary there.

To test the installation, simply run `frontier --help` from a terminal.  If the help for frontier is displayed, your installation was successful. Note: some shells, such as zsh, cache PATH lookups.  You may need to clear your cache  (by using `rehash` in zsh, for example) before trying to run `frontier --help`.


## Building

Should you decide not to use one of our prebuilt releases, you may instead build frontier from source.  To do so, you need to install some developer tools:

- A unix-like operating system with the common core commands (cp, tar, mkdir, bash, etc.)
- A compatible distribution of go (we officially support go 1.6 and later)
- [glide](https://glide.sh/)
- [git](https://git-scm.com/)

Provided your workstation satisfies the requirements above, follow the steps below:

1. Clone frontier's source:  `git clone https://github.com/digitalbitsorg/go.git && cd go`
2. Download external dependencies: `glide install`
3. Build the binary: `go install github.com/digitalbitsorg/go/services/frontier`

After running the above commands have succeeded, the built frontier will have be written into the `bin` subdirectory of the current directory.

Note:  Building directly on windows is not supported.


## Configuring

Frontier is configured using command line flags or environment variables.  To see the list of command line flags that are available (and their default values) for your version of frontier, run:

`frontier --help`

As you will see if you run the command above, frontier defines a large number of flags, however only three are required:

| flag                    | envvar                      | example                              |
|-------------------------|-----------------------------|--------------------------------------|
| `--db-url`              | `DATABASE_URL`              | postgres://localhost/horizon_testnet |
| `--stellar-core-db-url` | `STELLAR_CORE_DATABASE_URL` | postgres://localhost/core_testnet    |
| `--stellar-core-url`    | `STELLAR_CORE_URL`          | http://localhost:11626               |

`--db-url` specifies the frontier database, and its value should be a valid [PostgreSQL Connection URI](http://www.postgresql.org/docs/9.2/static/libpq-connect.html#AEN38419).  `--stellar-core-db-url` specifies a stellar-core database which will be used to load data about the stellar ledger.  Finally, `--stellar-core-url` specifies the HTTP control port for an instance of stellar-core.  This URL should be associated with the stellar-core that is writing to the database at `--stellar-core-db-url`.

Specifying command line flags every time you invoke frontier can be cumbersome, and so we recommend using environment variables.  There are many tools you can use to manage environment variables:  we recommend either [direnv](http://direnv.net/) or [dotenv](https://github.com/bkeepers/dotenv).  A template configuration that is compatible with dotenv can be found in the [frontier git repo](https://github.com/digitalbitsorg/go/blob/master/services/frontier/.env.template).



## Preparing the database

Before the frontier server can be run, we must first prepare the frontier database.  This database will be used for all of the information produced by frontier, notably historical information about successful transactions that have occurred on the DigitalBits network.  

To prepare a database for frontier's use, first you must ensure the database is blank.  It's easiest to simply create a new database on your postgres server specifically for frontier's use.  Next you must install the schema by running `frontier db init`.  Remember to use the appropriate command line flags or environment variables to configure frontier as explained in [Configuring ](#Configuring).  This command will log any errors that occur.

## Running

Once your frontier database is configured, you're ready to run frontier.  To run frontier you simply run `frontier` or `frontier serve`, both of which start the HTTP server and start logging to standard out.  When run, you should see some output that similar to:

```
INFO[0000] Starting frontier on :8000                     pid=29013
```

The log line above announces that frontier is ready to serve client requests. Note: the numbers shown above may be different for your installation.  Next we can confirm that frontier is responding correctly by loading the root resource.  In the example above, that URL would be [http://127.0.0.1:8000/] and simply running `curl http://127.0.0.1:8000/` shows you that the root resource can be loaded correctly.


## Ingesting digitalbits-core data

Frontier provides most of its utility through ingested data.  Your frontier server can be configured to listen for and ingest transaction results from the connected digitalbits-core.  We recommend that within your infrastructure you run one (and only one) frontier process that is configured in this way.   While running multiple ingestion processes will not corrupt the frontier database, your error logs will quickly fill up as the two instances race to ingest the data from digitalbits-core.  We may develop a system that coordinates multiple frontier processes in the future, but we would also be happy to include an external contribution that accomplishes this.

To enable ingestion, you must either pass `--ingest=true` on the command line or set the `INGEST` environment variable to "true".

### Managing storage for historical data

Given an empty frontier database, any and all available history on the attached digitalbits-core instance will be ingested. Over time, this recorded history will grow unbounded, increasing storage used by the database.  To keep you costs down, you may configure frontier to only retain a certain number of ledgers in the historical database.  This is done using the `--history-retention-count` flag or the `HISTORY_RETENTION_COUNT` environment variable.  Set the value to the number of recent ledgers you with to keep around, and every hour the frontier subsystem will reap expired data.  Alternatively, you may execute the command `frontier db reap` to force a collection.

### Surviving digitalbits-core downtime

Frontier tries to maintain a gap-free window into the history of the DigitalBits network.  This reduces the number of edge cases that frontier-dependent software must deal with, aiming to make the integration process simpler.  To maintain a gap-free history, frontier needs access to all of the metadata produced by digitalbits-core in the process of closing a ledger, and there are instances when this metadata can be lost.  Usually, this loss of metadata occurs because the digitalbits-core node went offline and performed a catchup operation when restarted.

To ensure that the metadata required by frontier is maintained, you have several options: You may either set the `CATCHUP_COMPLETE` digitalbits-core configuration option to `true` or configure `CATCHUP_RECENT` to determine the amount of time your digitalbits-core can be offline without having to rebuild your frontier database.

We _do not_ recommend using the `CATCHUP_COMPLETE` method, as this will force digitalbits-core to apply every transaction from the beginning of the ledger, which will take an ever increasing amount of time.  Instead, we recommend you set the `CATCHUP_RECENT` config value.  To do this, determine how long of a downtime you would like to survive (expressed in seconds) and divide by ten.  This roughly equates to the number of ledgers that occur within you desired grace period (ledgers roughly close at a rate of one every ten seconds).  With this value set, digitalbits-core will replay transactions for ledgers that are recent enough, ensuring that the metadata needed by frontier is present.

### Correcting gaps in historical data

In the section above, we mentioned that frontier _tries_ to maintain a gap-free window.  Unfortunately, it cannot directly control the state of digitalbits-core and so gaps may form due to extended down time.  When a gap is encountered, frontier will stop ingesting historical data and complain loudly in the log with error messages (log lines will include "ledger gap detected").  To resolve this situation, you must re-establish the expected state of the digitalbits-core database and purge historical data from frontier's database.  We leave the details of this process up to the reader as it is dependent upon your operating needs and configuration, but we offer one potential solution:

We recommend you configure the HISTORY_RETENTION_COUNT in frontier to a value less than or equal to the configured value for CATCHUP_RECENT in digitalbits-core.  Given this situation any downtime that would cause a ledger gap will require a downtime greater than the amount of historical data retained by frontier.  To re-establish continuity, simply:

1.  Stop frontier.
2.  run `frontier db reap` to clear the historical database.
3.  Clear the cursor for frontier by running `digitalbits-core -c "dropcursor?id=HORIZON"` (ensure capitilization is maintained).
4.  Clear ledger metadata from before the gap by running `digitalbits-core -c "maintenance?queue=true"`.
5.  Restart frontier.    

## Managing Stale Historical Data

Frontier ingests ledger data from a connected instance of digitalbits-core.  In the event that digitalbits-core stops running (or if frontier stops ingesting data for any other reason), the view provided by frontier will start to lag behind reality.  For simpler applications, this may be fine, but in many cases this lag is unacceptable and the application should not continue operating until the lag is resolved.

To help applications that cannot tolerate lag, frontier provides a configurable "staleness" threshold.  Given that enough lag has accumulated to surpass this threshold (expressed in number of ledgers), frontier will only respond with an error: [`stale_history`](./errors/stale-history.md).  To configure this option, use either the `--history-stale-threshold` command line flag or the `HISTORY_STALE_THRESHOLD` environment variable.  NOTE:  non-historical requests (such as submitting transactions or finding payment paths) will not error out when the staleness threshold is surpassed.

## Monitoring

To ensure that your instance of frontier is performing correctly we encourage you to monitor it, and provide both logs and metrics to do so.  

Frontier will output logs to standard out.  Information about what requests are coming in will be reported, but more importantly and warnings or errors will also be emitted by default.  A correctly running frontier instance will not ouput any warning or error log entries.

Metrics are collected while a frontier process is running and they are exposed at the `/metrics` path.  You can see an example at (https://frontier.testnet.digitalbits.io/metrics).

## I'm Stuck! Help!

If any of the above steps don't work or you are otherwise prevented from correctly setting up frontier, please come to our community and tell us. [Chat with us on slack](http://slack.digitalbits.io/) to ask for help.
