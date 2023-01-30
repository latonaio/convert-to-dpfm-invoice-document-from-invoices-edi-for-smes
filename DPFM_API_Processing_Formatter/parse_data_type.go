package dpfm_api_processing_formatter

import (
	"strconv"
	"time"
)

func ParseInt(s string) (int, error) {
	res, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	return res, nil
}

func ParseIntPtr(s *string) (*int, error) {
	if s == nil {
		return nil, nil
	} else if *s == "" {
		return nil, nil
	}

	res, err := strconv.Atoi(*s)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func GetIntPtr(i int) *int {
	return &i
}

func ParseFloat32Ptr(s *string) (*float32, error) {
	if s == nil {
		return nil, nil
	} else if *s == "" {
		return nil, nil
	}

	f, err := strconv.ParseFloat(*s, 32)
	if err != nil {
		return nil, err
	}

	res := float32(f)

	return &res, nil
}

func GetFloat32Ptr(f float32) *float32 {
	return &f
}

func ParseBoolPtr(s *string) (*bool, error) {
	if s == nil {
		return nil, nil
	} else if *s == "" {
		return nil, nil
	}

	res, err := strconv.ParseBool(*s)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func GetBoolPtr(b bool) *bool {
	return &b
}

func GetStringPtr(s string) *string {
	return &s
}

func GetSystemDatePtr() *string {
	day := time.Now()
	res := day.Format("2006-01-02")

	return &res
}

func GetSystemTimePtr() *string {
	day := time.Now()
	res := day.Format("15:04:05")

	return &res
}
