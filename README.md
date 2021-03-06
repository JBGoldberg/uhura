# Uhura

## Overview

*Uhura* aims to keep communications open with customers using the any channel.

It accepts commands to several different command to process connections and exchange messages using AMPQ, SMTP or Telegram.

## Configuration

*Uhura* can receive configuration data from a file called `.uhura` in the path or through enviroment variables:

* DEFAULT_MODE: Defines the default mode for the current run.

## Use

### Modes

#### SMTP

```
uhura smtp
```

* -f EMAIL:  