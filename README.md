# timeagoo
[![Build Status](https://travis-ci.org/victorjonsson/timeagoo.svg?branch=master)](https://travis-ci.org/victorjonsson/timeagoo)

This package provides you with a method that describes the time between `time.Now()` and a given time.
It's an alternative package to [justincampbell/timeago](https://github.com/justincampbell/timeago).


```
↓ Formatted date (eg. Mon Jan 2 2015)
↓ Yesterday hh:mm 
↓ Today hh:mm
↓ About %s minutes ago 
↓ About one minute ago 
↓ About %s seconds ago
↓ Recently
- Now
↑ Soon
↑ In about %s seconds
↑ In about one minute
↑ In about %s minutes
↑ Today hh:mm
↑ Tomorrow hh:mm
↑ Formatted date (eg. Mon Jan 2 2021)
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
