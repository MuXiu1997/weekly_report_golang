package report

import "time"

type WorkPeriod struct {
}

func (wp *WorkPeriod) StartDay() time.Time {
	_weekday := int(today().Weekday())
	if _weekday == 0 {
		_weekday = 7
	}
	daysDelta := -2 - _weekday
	if _weekday > 4 {
		daysDelta += 7
	}
	return today().AddDate(0, 0, daysDelta)
}

func (wp *WorkPeriod) EndDay() time.Time {
	return wp.StartDay().AddDate(0, 0, 6)
}

func (wp *WorkPeriod) workDays() map[string]time.Time {
	sd := wp.StartDay()
	return map[string]time.Time{
		"l_Fri": sd.AddDate(0, 0, 0),
		"Mon":   sd.AddDate(0, 0, 3),
		"Tue":   sd.AddDate(0, 0, 4),
		"Wed":   sd.AddDate(0, 0, 5),
		"Thu":   sd.AddDate(0, 0, 6),
	}
}

func (wp *WorkPeriod) daysRange() map[string]time.Time {
	sd := wp.StartDay()
	return map[string]time.Time{
		"t_week_start": sd.AddDate(0, 0, 3),
		"t_week_end":   sd.AddDate(0, 0, 7),
		"n_week_start": sd.AddDate(0, 0, 10),
		"n_week_end":   sd.AddDate(0, 0, 14),
	}
}

func (wp *WorkPeriod) WorkDaysLabel() map[string]string {
	wdl := make(map[string]string, 5)
	for key, value := range wp.workDays() {
		wdl[key] = dayFormat(value)
	}
	return wdl
}

func (wp *WorkPeriod) DaysRangeLabel() map[string]string {
	r := wp.daysRange()
	drl := map[string]string{
		"t_week": dayRangeFormat(r["t_week_start"], r["t_week_end"]),
		"n_week": dayRangeFormat(r["n_week_start"], r["n_week_end"]),
	}
	return drl
}

func (wp *WorkPeriod) Labels() map[string]string {
	l := make(map[string]string, 7)
	for key, value := range wp.WorkDaysLabel() {
		l[key] = value
	}
	for key, value := range wp.DaysRangeLabel() {
		l[key] = value
	}
	return l
}
