package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)


var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")


type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {

	var jsonValue string = strconv.Quote(fmt.Sprintf("%d mins", r))
	
	return []byte(jsonValue), nil

}

func (r *Runtime) UnmarshalJSON(data []byte) (error) {

	unqoutedJSONValue, err := strconv.Unquote(string(data))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}
	
	var parts []string = strings.Split(unqoutedJSONValue, " ")

	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}

	minutes, err := strconv.ParseInt(parts[0], 10 ,32)
	
	if err != nil  {
		return ErrInvalidRuntimeFormat
	}

	*r = Runtime(minutes)

	return nil

}

