package report

import (
	"fmt"
	"github.com/json-iterator/go"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
	"weekly_report_golang/setting"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func init() {
	_, err := os.Stat(setting.JsonFilePath)
	if err == nil {
		return
	}

	initReportFile()
}

func initReportFile() {
	dirname := filepath.Dir(setting.JsonFilePath)
	dirStat, err := os.Stat(dirname)
	if err == nil && !dirStat.IsDir() {
		err := os.Remove(dirname)
		if err != nil {
			panic(err)
		}
	}

	err = os.MkdirAll(dirname, 0644)
	if err != nil {
		panic(err)
	}
	r := New()
	data := make(map[string]string, 7)
	for _, value := range r.Keys() {
		data[value] = "待定"
	}
	r.Update(data)
	r.Save()
}

func New() *Report {
	wp := WorkPeriod{}
	r := Report{
		data: map[string]string{},
		wp:   &wp,
	}
	return &r
}

type Report struct {
	data map[string]string
	wp   *WorkPeriod
}

func (r *Report) Load() *Report {
	data := make(map[string]string)
	bytes, err := ioutil.ReadFile(setting.JsonFilePath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		panic(err)
	}
	r.data = data
	return r
}

func (r *Report) Save() {
	bytes, err := json.Marshal(r.data)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(setting.JsonFilePath, bytes, 0644)
	if err != nil {
		panic(err)
	}
}

func (r *Report) Update(data map[string]string) *Report {
	r.data = data
	return r
}

func (r *Report) Works() map[string]string {
	return r.data
}

func (r *Report) Labels() map[string]string {
	return r.wp.Labels()
}

func (r *Report) Items() map[string]string {
	works := r.Works()
	items := make(map[string]string, 7)
	for key, value := range r.wp.WorkDaysLabel() {
		items[key] = value + works[key]
	}
	for key := range r.wp.DaysRangeLabel() {
		items[key] = works[key]
	}
	return items
}

func (r *Report) Keys() []string {
	return []string{"l_Fri", "Mon", "Tue", "Wed", "Thu", "t_week", "n_week"}
}

func (r *Report) Message() string {
	l := r.Labels()
	w := r.Works()
	return fmt.Sprintf("\n%s完成事项：%s。\n%s工作计划：%s。",
		l["t_week"],
		w["t_week"],
		l["n_week"],
		w["n_week"],
	)
}

func (r *Report) EndDay() time.Time {
	return r.wp.EndDay()
}
