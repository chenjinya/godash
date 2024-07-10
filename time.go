package godash

import (
    "database/sql"
    "encoding/json"
    "log"
    "math/rand"
    "time"
)

func RandSleep() {

    r := rand.New(rand.NewSource(time.Now().Unix()))
    seed := r.Int() % 1000
    log.Println("Rand Sleep:", seed)
    time.Sleep(time.Duration(seed) * time.Millisecond)
}

func TimePtr(v time.Time) *time.Time {
    return &v
}

func PtrTime(v *time.Time) time.Time {
    if nil == v {
        return time.Time{}
    }
    return *v
}

type NullTime struct {
    sql.NullTime
}

func (nt NullTime) MarshalJSON() ([]byte, error) {
    if nt.Valid {
        return json.Marshal(nt.Time)
    } else {
        return []byte(""), nil
    }
}

func (nt *NullTime) UnmarshalJSON(data []byte) error {
    if nil != data {
        nt.Valid = true
        return json.Unmarshal(data, &nt.Time)
    } else {
        nt.Time = time.Time{}
        nt.Valid = false
        return nil
    }
}

func GetMonday() time.Time {
    now := time.Now()
    year1 := time.Date(now.Year(), time.January, 1, 0, 0, 0, 0, now.Location())
    days := int(time.Since(year1) / time.Hour / 24)
    passedweeks := days / 7
    monday := year1.Add(time.Duration(passedweeks) * 7 * 24 * time.Hour)
    return monday
}

func GetMonthDay() time.Time {
    now := time.Now()
    monthday := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
    return monthday
}
