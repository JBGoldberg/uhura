# Uhura

## Overview

*Uhura* aims to keep communications open with customers using the any channel.

It accepts commands to several different command to process connections and exchange messages using AMPQ, SMTP or Telegram.

## Configuration

*Uhura* can receive configuration data from a file called `.uhura.yaml` in current path, user home or through enviroment variables.

Check the file `dot_uhura.yaml` for examples.

## Use

### Actions

You should define a action to be executed. There are 2 diffent possibilites: receive or send messages.

Each action must specify one or several communication channels to use.

### Communication Channels

Channels are any way to exchange messages with users. They can vary from traditional ones like SMTP/POP/Facebook, some more updated as Linkedin/Telegram/Signal or new ones like Unstoppable Chat.

For now, I am working to implement smtp as the first channel available.
#### SMTP

```bash
uhura send smtp
```