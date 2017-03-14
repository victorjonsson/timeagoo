package language;

import "time"

type Dialogs interface {
    Now() string
    Recently() string
    Soon() string
    InXSeconds() string
    XSecondsAgo() string
    InOneMinute() string
    OneMinuteAgo() string
    InXMinutes() string
    XMinutesAgo() string
    FormatTodayDate(t time.Time) string
    FormatYesterdayDate(t time.Time) string
    FormatTomorrowDate(t time.Time) string
    FormatDefaultDateFormat(t time.Time) string
}

type EnglishDialogs struct {
}

func (EnglishDialogs) Now() string {
    return "Right now";
}

func (EnglishDialogs) Recently() string {
    return "Recently"
}

func (EnglishDialogs) Soon() string {
    return "Soon"
}

func (EnglishDialogs) InXSeconds() string {
    return "In about %s seconds"
}

func (EnglishDialogs) XSecondsAgo() string {
    return "About %s seconds ago"
}

func (EnglishDialogs) InXMinutes() string {
    return "In about %s minutes"
}

func (EnglishDialogs) InOneMinute() string {
    return "In about 1 minute"
}
func (EnglishDialogs) OneMinuteAgo() string {
    return "About 1 minute ago"
}

func (EnglishDialogs) XMinutesAgo() string {
    return "About %s minutes ago"
}

func (EnglishDialogs) FormatTodayDate(t time.Time) string {
    return t.Format("Today 3:04PM")
}

func (EnglishDialogs) FormatYesterdayDate(t time.Time) string {
    return t.Format("Yesterday 3:04PM")
}

func (EnglishDialogs) FormatTomorrowDate(t time.Time) string {
    return t.Format("Tomorrow 3:04PM")
}

func (EnglishDialogs) FormatDefaultDateFormat(t time.Time) string {
    return t.Format("Mon Jan 2 15:04:05")
}
