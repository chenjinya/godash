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

func (nt NullTime) MarshalJSON()([]byte, error){
	if nt.Valid {
		return json.Marshal(nt.Time)
	} else {
		return []byte(""), nil
	}
}

func (nt *NullTime) UnmarshalJSON(data []byte) error{
	if nil != data {
		nt.Valid = true
		return json.Unmarshal(data, &nt.Time)
	} else {
		nt.Time = time.Time{}
		nt.Valid = false
		return nil
	}
}
