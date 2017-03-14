package timeagoo

import (
    "testing"
    "time"
    "fmt"
)

func TestFormat(t *testing.T) {
    english := &EnglishDialogs{}
    cases := map[time.Time]string{
        getTime(0): english.Now(),

        getTime(9):       english.Soon(),
        getTime(10):      fmt.Sprintf(english.InXSeconds(), "10"),
        getTime(60):      fmt.Sprintf(english.InXSeconds(), "60"),
        getTime(61):      english.InOneMinute(),
        getTime(119):     fmt.Sprintf(english.InXMinutes(), "2"),
        getTime(1800):    english.FormatTodayDate(getTime(1800)),
        getTime(86400):   english.FormatTomorrowDate(getTime(86400)),
        getTime(8640000): english.FormatDefaultDateFormat(getTime(8640000)),

        getTime(-9):       english.Recently(),
        getTime(-10):      fmt.Sprintf(english.XSecondsAgo(), "10"),
        getTime(-60):      fmt.Sprintf(english.XSecondsAgo(), "60"),
        getTime(-61):      english.OneMinuteAgo(),
        getTime(-119):     fmt.Sprintf(english.XMinutesAgo(), "2"),
        getTime(-1800):    english.FormatTodayDate(getTime(-1800)),
        getTime(-86400):   english.FormatYesterdayDate(getTime(-86400)),
        getTime(-8640000): english.FormatDefaultDateFormat(getTime(-8640000)),
    }

    for input, expected := range cases {
        output := Format(input, english)
        if output != expected {
            t.Errorf("Format(%q, english) == %q, expected %q", input, output, expected)
        }
    }
}

func getTime(secondsAgo int64) time.Time {
    return time.Unix(time.Now().Unix()+secondsAgo, 0)
}
