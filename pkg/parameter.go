package lada

import (
	"strconv"
	"strings"
)

type Parameter struct {
	Name         string
	ShortForm    string
	DefaultValue string
	Description  string
	Value        string
	IsFlag       bool
}

func (p Parameter) AsInt() (int, error) {
	if p.IsFlag {
		return 0, CannotUseFlagAsAValueError.New()
	}
	value, err := strconv.Atoi(p.Value)
	return value, err
}

func (p Parameter) AsRangedInt(min int, max int) (int, error) {
	value, err := p.AsInt()
	if err != nil {
		return 0, InvalidParameterValueError.New(p.Name, p.Value).CausedBy(err)
	}

	if min <= value && value <= max {
		return value, nil
	}
	return 0, InvalidParameterValueError.New(p.Name, p.Value)
}

func (p Parameter) AsBool() (bool, error) {
	if p.IsFlag {
		return false, CannotUseFlagAsAValueError.New()
	}
	value, err := strconv.ParseBool(p.Value)
	return value, err
}

func (p Parameter) AsFloat() (float64, error) {
	if p.IsFlag {
		return 0, CannotUseFlagAsAValueError.New()
	}
	value, err := strconv.ParseFloat(p.Value, 64)
	return value, err
}

func (p Parameter) AsRangedFloat(min float64, max float64) (float64, error) {
	value, err := p.AsFloat()
	if err != nil {
		return 0, InvalidParameterValueError.New(p.Name, p.Value).CausedBy(err)
	}

	if min <= value && value <= max {
		return value, nil
	}
	return 0, InvalidParameterValueError.New(p.Name, p.Value)
}

func (p Parameter) AsStringList() ([]string, error) {
	if p.IsFlag {
		return []string{}, CannotUseFlagAsAValueError.New()
	}
	return strings.Split(p.Value, ","), nil
}

func (p Parameter) AsIntList() ([]int, error) {
	if p.IsFlag {
		return []int{}, CannotUseFlagAsAValueError.New()
	}
	var result []int
	for _, item := range strings.Split(p.Value, ",") {
		value, _ := strconv.Atoi(item)
		result = append(result, value)
	}

	return result, nil
}

func (p Parameter) AsFloatList() ([]float64, error) {
	if p.IsFlag {
		return []float64{}, CannotUseFlagAsAValueError.New()
	}
	var result []float64
	for _, item := range strings.Split(p.Value, ",") {
		value, _ := strconv.ParseFloat(item, 64)
		result = append(result, value)
	}

	return result, nil
}

func (p Parameter) AsIntEnum(enumMap map[string]int) (int, error) {
	if p.IsFlag {
		return 0, CannotUseFlagAsAValueError.New()
	}

	if value, ok := enumMap[p.Value]; ok {
		return value, nil
	}

	return -1, InvalidParameterValueError.New(p.Name, p.Value)
}

func (p Parameter) AsStringEnum(enumMap map[string]string) (string, error) {
	if p.IsFlag {
		return "", CannotUseFlagAsAValueError.New()
	}

	if value, ok := enumMap[p.Value]; ok {
		return value, nil
	}

	return "", InvalidParameterValueError.New(p.Name, p.Value)
}

func (p Parameter) IsEnabled() bool {
	if p.Value == "1" {
		return true
	}

	return false
}