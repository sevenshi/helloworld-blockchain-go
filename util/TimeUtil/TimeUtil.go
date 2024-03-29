package TimeUtil

/*
 @author king 409060350@qq.com
*/

import (
	"time"
)

func MillisecondTimestamp() uint64 {
	return uint64(time.Now().UnixNano()) / uint64(1000000)
}
func FormatMillisecondTimestamp(millisecondTimestamp uint64) string {
	return time.Unix(0, int64(millisecondTimestamp)*int64(1000000)).Format("2006-01-02T15:04:05.999")
}
