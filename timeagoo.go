package timeagoo

import (
    "time"
    "math"
    "fmt"
    "strconv"
)


/* * * Format Function * * */


func Format(t time.Time, dialogs Dialogs) string {
    dateComparison := newDateComparison(t)
    if t.Unix() == dateComparison.now.Unix() {
        return dialogs.Now()
    } else if dateComparison.now.After(t) {
        return getDialogForHistoricTime(dateComparison, dialogs)
    } else {
        return getDialogForFutureTime(dateComparison, dialogs)
    }
}

func getDialogForFutureTime(dateComparison dateComparison, dialogs Dialogs) string {
    var str string
    if dateComparison.diff < 10 {
        str = dialogs.Soon()
    } else if dateComparison.diff <= 60 {
        str = fmt.Sprintf(dialogs.InXSeconds(), strconv.Itoa(dateComparison.diff))
    } else if dateComparison.diff <= 90 {
        str = dialogs.InOneMinute()
    } else if dateComparison.diff <= 600 {
        str = fmt.Sprintf(dialogs.InXMinutes(), toMinutes(dateComparison.diff))
    } else if dateComparison.isTodayDate() {
        str = dialogs.FormatTodayDate(dateComparison.givenTime)
    } else if dateComparison.isTomorrowDate() {
        str = dialogs.FormatTomorrowDate(dateComparison.givenTime)
    } else {
        str = dialogs.FormatDefaultDateFormat(dateComparison.givenTime)
    }
    return str
}

func getDialogForHistoricTime(dateComparison dateComparison, dialogs Dialogs) string {
    var str string
    if dateComparison.diff < 10 {
        str = dialogs.Recently()
    } else if dateComparison.diff <= 60 {
        str = fmt.Sprintf(dialogs.XSecondsAgo(), strconv.Itoa(dateComparison.diff))
    } else if dateComparison.diff <= 90 {
        str = dialogs.OneMinuteAgo()
    } else if dateComparison.diff <= 600 {
        str = fmt.Sprintf(dialogs.XMinutesAgo(), toMinutes(dateComparison.diff))
    } else if dateComparison.isTodayDate() {
        str = dialogs.FormatTodayDate(dateComparison.givenTime)
    } else if dateComparison.isYesterdayDate() {
        str = dialogs.FormatYesterdayDate(dateComparison.givenTime)
    } else {
        str = dialogs.FormatDefaultDateFormat(dateComparison.givenTime)
    }
    return str
}


/* * * Dialog interface * * */

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


/* * * * dateComparison  * * * */

type dateComparison struct {
    now       time.Time
    givenTime time.Time
    diff      int
}

func newDateComparison(givenTime time.Time) dateComparison {
    d := dateComparison{now: time.Now(), givenTime: givenTime}
    d.diff = int(math.Abs(float64(d.now.Unix() - givenTime.Unix())))
    return d
}

func (d *dateComparison) isTodayDate() bool {
    return isSameDate(d.now, d.givenTime)
}

func (d *dateComparison) isTomorrowDate() bool {
    return isSameDate(d.now.AddDate(0, 0, 1), d.givenTime);
}

func (d *dateComparison) isYesterdayDate() bool {
    return isSameDate(d.now.AddDate(0, 0, -1), d.givenTime);
}

/* * * * Utility methods * * * */

func toMinutes(seconds int) string {
    i := int(round(float64(seconds) / float64(60)))
    return strconv.Itoa(i)
}

func round(val float64) (newVal float64) {
    var round float64
    places := 1
    pow := math.Pow(10, float64(places))
    digit := pow * val
    _, div := math.Modf(digit)
    if div >= 0.5 {
        round = math.Ceil(digit)
    } else {
        round = math.Floor(digit)
    }
    newVal = round / pow
    return
}

func isSameDate(t1 time.Time, t2 time.Time) bool {
    return t1.Day() == t2.Day() && t1.Month() == t2.Month() && t1.Year() == t2.Year()
}
