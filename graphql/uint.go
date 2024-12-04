package graphql

import (
	"github.com/goccy/go-json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

func MarshalUint(i uint) Marshaler {
	return WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, strconv.FormatUint(uint64(i), 10))
	})
}

func UnmarshalUint(v any) (uint, error) {
	switch v := v.(type) {
	case string:
		u64, err := strconv.ParseUint(v, 10, 64)
		return uint(u64), err
	case int:
		if v < 0 {
			return 0, errors.New("cannot convert negative numbers to uint")
		}

		return uint(v), nil
	case int64:
		if v < 0 {
			return 0, errors.New("cannot convert negative numbers to uint")
		}

		return uint(v), nil
	case json.Number:
		u64, err := strconv.ParseUint(string(v), 10, 64)
		return uint(u64), err
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("%T is not an uint", v)
	}
}

func MarshalUint64(i uint64) Marshaler {
	return WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, strconv.FormatUint(i, 10))
	})
}

func UnmarshalUint64(v any) (uint64, error) {
	switch v := v.(type) {
	case string:
		return strconv.ParseUint(v, 10, 64)
	case int:
		if v < 0 {
			return 0, errors.New("cannot convert negative numbers to uint64")
		}

		return uint64(v), nil
	case int64:
		if v < 0 {
			return 0, errors.New("cannot convert negative numbers to uint64")
		}

		return uint64(v), nil
	case json.Number:
		return strconv.ParseUint(string(v), 10, 64)
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("%T is not an uint", v)
	}
}

func MarshalUint32(i uint32) Marshaler {
	return WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, strconv.FormatUint(uint64(i), 10))
	})
}

func UnmarshalUint32(v any) (uint32, error) {
	switch v := v.(type) {
	case string:
		iv, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return 0, err
		}
		return uint32(iv), nil
	case int:
		if v < 0 {
			return 0, errors.New("cannot convert negative numbers to uint32")
		}

		return uint32(v), nil
	case int64:
		if v < 0 {
			return 0, errors.New("cannot convert negative numbers to uint32")
		}

		return uint32(v), nil
	case json.Number:
		iv, err := strconv.ParseUint(string(v), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint32(iv), nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("%T is not an uint", v)
	}
}
