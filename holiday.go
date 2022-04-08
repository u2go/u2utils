package u2utils

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"time"
)

type holidayRs struct {
	Code    int `json:"code"`
	Holiday map[string]struct {
		Holiday bool   `json:"holiday"`
		Date    string `json:"date"`
	}
}

// HolidayChina 判断日期是否是节假日。
// 日期格式：yyyy-mm-dd，如：2022-04-08
func HolidayChina(date string) (bool, error) {
	var holidays map[string]bool
	cacheKey := "__holidays"
	err := TmpCacheGet(cacheKey, &holidays)
	if err != nil {
		return false, err
	}
	if holidays == nil {
		r, err := http.Get("http://timor.tech/api/holiday/year/" + date[:4])
		if err != nil {
			return false, err
		}
		if r.StatusCode != 200 {
			return false, errors.Errorf("git holiday from api failed status_code=%d", r.StatusCode)
		}
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return false, err
		}
		var rs holidayRs
		err = json.Unmarshal(b, &rs)
		if err != nil {
			return false, err
		}
		holidays = map[string]bool{}
		for _, v := range rs.Holiday {
			holidays[v.Date] = v.Holiday
		}
		err = TmpCacheSet(cacheKey, holidays)
		if err != nil {
			return false, err
		}
	}

	if v, ok := holidays[date]; ok {
		return v, nil
	}

	// is weekend
	return IsWeekend(date)
}

func IsWeekend(date string) (bool, error) {
	tt, err := time.Parse("2006-01-02", date)
	if err != nil {
		return false, err
	}
	switch tt.Weekday() {
	case time.Saturday, time.Sunday:
		return true, nil
	}
	return false, nil
}
