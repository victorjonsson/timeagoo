# timeagoo
[![Build Status](https://travis-ci.org/victorjonsson/timeagoo.svg?branch=master)](https://travis-ci.org/victorjonsson/timeagoo)

This package provides you with a method that describes the time between `time.Now()` and a given time.
It's an alternative package to [justincampbell/timeago](https://github.com/justincampbell/timeago).


```
... <-> -48h    # => Formatted date (eg. Mon Jan 2 15:04:05)
-48h <-> -24h   # => Yesterday hh:mm 
-24h <-> -10m   # => Today hh:mm
-10m <-> -90s   # => About %s minutes ago 
-90s <-> -60s   # => About one minute ago 
-60s <-> -10s   # => About %s seconds ago
-10s <-> 0s     # => Recently
0s              # => Now
1s <-> 10s      # => Soon
10sec <-> 60s   # => In about %s seconds
60sec <-> 90s   # => In about one minute
90s <-> 10m     # => In about %s minutes
10m <-> 24h     # => Today hh:mm
24h <-> 48h     # => Tomorrow hh:mm
48h <-> ...     # => Formatted date (eg. Mon Jan 2 15:04:05)
```

## API

`func Format(time time.Time, dialogs timeagoo.language.Dialogs) string`

```go
package main

import (
    "github.com/victorjonsson/timeagoo"
    ...
)

...

func formatWhenIncidentOccurred(i Incident) {    
    return timeagoo.Format(i.time, timeagoo.EnglishDialogs{})
}

```

## Multilingual support

The package comes with an English translation of all dialogs. You can use another language by
 implementing the [Dialogs interface](https://github.com/victorjonsson/timeagoo/blob/master/timeagoo.go#L68).

```go
import "time"

type swedishDialogs struct {
}

func (swedishDialogs) Recently() string {
    return "Nyligen";
}
...

func FormatTimeagoo(t time.Time) string {
    return timeagoo.Format(t, swedishDialogs{})
}

```
