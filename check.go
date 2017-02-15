package IDCardCheck

import (
	"errors"
	"strconv"
	"strings"
)

var (
	wi        = [17]uint32{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	ci        = [11]string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}
	toal, mod uint32
)

func NumberCheck(numStr string) (err error) {
	if len(numStr) != 18 {
		return errors.New("length wrong")
	}
	numStr = strings.ToUpper(numStr)
	var tmp uint64
	//check month and day
	if tmp, err = strconv.ParseUint(numStr[10:12], 10, 32); err != nil {
		return errors.New("check month wrong")
	}
	if tmp < 1 || tmp > 12 {
		return errors.New("month wrong")
	}
	if tmp, err = strconv.ParseUint(numStr[12:14], 10, 32); err != nil {
		return errors.New("check day wrong")
	}
	if tmp < 1 || tmp > 31 {
		return errors.New("day wrong")
	}
	//format, mul and sum
	toal = 0
	for i := 0; i < 17; i++ {
		if tmp, err = strconv.ParseUint(numStr[i:i+1], 10, 32); err != nil {
			return err
		}
		toal += (uint32(tmp) * wi[i])
	}
	//mod
	mod = toal % 11
	if numStr[17:18] != ci[mod] {
		return errors.New("check wrong")
	}
	return
}
