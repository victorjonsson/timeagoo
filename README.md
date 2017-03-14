# timeagoo

This package provides you with a method that describes the time between `time.Now()` and a given time.
It's an alternative package to [justincampbell/timeago](https://github.com/justincampbell/timeago).


## API

`func Format(t time.Time, timeagoo.language.Dialogs) string`

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

Lorem te ipsum...

