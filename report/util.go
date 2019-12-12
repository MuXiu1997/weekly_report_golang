package report

import (
	"fmt"
	"time"
)

func today() *time.Time {
	layout := "2006-01-02"
	s := time.Now().Format(layout)
	t, _ := time.Parse(layout, s)
	return &t
}

func dayFormat(t time.Time) string {
	_map := map[time.Weekday]string{
		1: "一",
		2: "二",
		3: "三",
		4: "四",
		5: "五",
	}
	return t.Format("01月02日") + fmt.Sprintf("（周%s）", _map[t.Weekday()])
}

func dayRangeFormat(startTime, endTime time.Time) string {
	month := int(startTime.Month())
	weekNo := calcWeekOn(startTime)
	layout := "01.02"
	return fmt.Sprintf("%v月第%v周(%s~%s)",
		month,
		weekNo,
		startTime.Format(layout),
		endTime.Format(layout),
	)
}

func calcWeekOn(t time.Time) string {
	_map := map[int]string{
		1: "一",
		2: "二",
		3: "三",
		4: "四",
		5: "五",
	}
	s := t.Format("2006-01")
	firstDay, _ := time.Parse("2006-01-02", s+"-01")
	firstDayWeekday := int(firstDay.Weekday())

	firstMonday := firstDay
	if firstDayWeekday != 1 {
		if firstDayWeekday == 0 {
			firstDayWeekday = 7
		}
		firstMonday = firstDay.AddDate(0, 0, 8-firstDayWeekday)
	}
	//week_no_int = (t - firstMonday).days // 7 + 1
	weekNoInt := (t.Day()-firstMonday.Day())/7 + 1
	return _map[weekNoInt]
}
