# timeagoo

This package provides you with a method that describes the time between `time.Now()` and a given time.
It's an alternative package to [justincampbell/timeago](https://github.com/justincampbell/timeago).


## API

`func Format(time time.Time, dialogs timeagoo.language.Dialogs) string`

```go
package main

import (
	"github.com/victorjonsson/timeagoo"
	"github.com/victorjonsson/timeagoo/language"
	...
)

...

func formatWhenIncidentOccurred(i Incident) {
    return timeagoo.Format(i.time, language.EnglishDialogs{})
}

```

## Multilingual support

The package comes with an English translation of all dialogs. You can use another language by
 implementing the [Dialogs interface](https://github.com/victorjonsson/timeagoo/blob/master/language/language.go#L5).

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
