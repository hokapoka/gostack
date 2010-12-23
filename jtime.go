package gostack
import (
	"json"
	"time"
	"os"
)

const Fmt = "2006-01-02T15:04:05"
type jTime time.Time
func (jt *jTime) UnmarshalJSON(data []byte) os.Error {
	var n int64
	if err := json.Unmarshal(data, &n); err != nil {
        return err
	}
	t := time.SecondsToLocalTime(n)
	*jt = (jTime)(*t)
	return nil
}

func (jt jTime) MarshalJSON() ([]byte, os.Error) {
	return json.Marshal((*time.Time)(&jt).Format(Fmt))
}

