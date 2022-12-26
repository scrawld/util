package util

import (
	"fmt"
	"strconv"
	"time"
)

// 获取指定日期零点时间
func FirstTimeOfDay(t time.Time) (r time.Time) {
	r = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	return
}

// 获取日期 GetDtByOffset(time.Now(), 0) return 20060102
func GetDtByOffset(tm time.Time, offset int) (r int) {
	r, _ = strconv.Atoi(tm.AddDate(0, 0, offset).Format("20060102"))
	return
}

// 日期范围 GetDtRange(20200101, 20200203) ruturn []int{20200101,20200102,20200103}
func GetDtRange(start_dt, end_dt int) (r []int, err error) {
	if start_dt > end_dt {
		return
	}
	start_dt_str := strconv.Itoa(start_dt)
	end_dt_str := strconv.Itoa(end_dt)

	var tm_layout = "20060102"
	start_tm, err := time.Parse(tm_layout, start_dt_str)
	if err != nil {
		err = fmt.Errorf("time parse(%s, %s) error(%s)", tm_layout, start_dt_str, err)
		return
	}
	end_tm, err := time.Parse(tm_layout, end_dt_str)
	if err != nil {
		err = fmt.Errorf("time parse(%s, %s) error(%s)", tm_layout, end_dt_str, err)
		return
	}
	dt_interval := (end_tm.Unix() - start_tm.Unix()) / (24 * 60 * 60)

	for i := 0; i <= int(dt_interval); i++ {
		dt_str := start_tm.AddDate(0, 0, i).Format(tm_layout)
		dt, err := strconv.Atoi(dt_str)
		if err != nil {
			err = fmt.Errorf("strconv atoi(%s) error(%s)", dt_str, err)
			return r, err
		}
		r = append(r, dt)
	}
	return
}

// max-min=days DtSub(20200302, 20200301) = 1
func DtSub(max, min int) (r int, err error) {
	if max < min {
		return
	}
	max_dt_str := strconv.Itoa(max)
	min_dt_str := strconv.Itoa(min)

	var tm_layout = "20060102"

	var max_tm time.Time
	max_tm, err = time.Parse(tm_layout, max_dt_str)
	if err != nil {
		err = fmt.Errorf("time parse(%s, %s) error(%s)", tm_layout, max_dt_str, err)
		return
	}
	var min_tm time.Time
	min_tm, err = time.Parse(tm_layout, min_dt_str)
	if err != nil {
		err = fmt.Errorf("time parse(%s, %s) error(%s)", tm_layout, min_dt_str, err)
		return
	}
	r = int(max_tm.Sub(min_tm).Hours() / 24)
	return
}

// 获取日期到小时 GetDataHour(time.Now()) return 2022122609
func GetDataHour(tm time.Time) (r int) {
	r, _ = strconv.Atoi(tm.Format("2006010215"))
	return
}

// 获取日期范围 GetDataHourRange(1672012800, 1672027997) ruturn []int{2022122608 2022122609 2022122610 2022122611 2022122612}
func GetDataHourRange(st, et int64) (r []int) {
	if st > et {
		return
	}
	var (
		stTime   = time.Unix(st, 0)
		etTime   = time.Unix(et, 0)
		stZero   = time.Date(stTime.Year(), stTime.Month(), stTime.Day(), stTime.Hour(), 0, 0, 0, time.Local)
		etZero   = time.Date(etTime.Year(), etTime.Month(), etTime.Day(), etTime.Hour(), 0, 0, 0, time.Local)
		interval = etZero.Sub(stZero).Hours()
	)
	for i := 0; i <= int(interval); i++ {
		var (
			t  = stZero.Add(time.Hour * time.Duration(i))
			dh = GetDataHour(t)
		)
		r = append(r, dh)
	}
	return
}

// 获取周次 GetDtRange(time.Unix(1672588800, 0)) return 2023.1
func GetWeek(tm time.Time) string {
	year, week := tm.ISOWeek()
	return fmt.Sprintf("%d.%d", year, week)
}

// 获取时间范围内的周次 GetWeekRange(1671790242, 1672588800) return []string{"2022.51", "2022.52", "2023.1"}
func GetWeekRange(st, et int64) (r []string) {
	if st > et {
		return
	}
	var (
		stTime   = time.Unix(st, 0)
		etTime   = time.Unix(et, 0)
		stZero   = time.Date(stTime.Year(), stTime.Month(), stTime.Day(), 0, 0, 0, 0, time.Local)
		etZero   = time.Date(etTime.Year(), etTime.Month(), etTime.Day(), 0, 0, 0, 0, time.Local)
		interval = etZero.Sub(stZero).Hours() / 24
	)
	m := map[string]byte{}
	for i := 0; i <= int(interval); i++ {
		t := stZero.AddDate(0, 0, i)
		week := GetWeek(t)
		if _, ok := m[week]; ok {
			continue
		}
		m[week] = 0
		r = append(r, week)
	}
	return
}
