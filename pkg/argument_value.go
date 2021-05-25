package lada

import (
	"strconv"
	"strings"
)

type ArgumentValue struct {
	argument *Argument
	value string
}

func NewArgumentValue(a *Argument, v string) ArgumentValue {
	return ArgumentValue{
		argument: a,
		value: v,
	}
}

func (a ArgumentValue) Value() string {
	if a.value != "" {
		return a.value
	}

	return a.argument.DefaultValue()
}

func (a ArgumentValue) AsString() (string, error) {
	return a.value, nil
}

func (a ArgumentValue) AsInt() (int, error) {
	value, err := strconv.Atoi(a.value)
	return value, err
}

func (a ArgumentValue) AsRangedInt(min int, max int) (int, error) {
	value, err := a.AsInt()
	if err != nil {
		return 0, InvalidArgumentValueError.New(a.argument.Name, a.value).CausedBy(err)
	}

	if min <= value && value <= max {
		return value, nil
	}

	return 0, InvalidArgumentValueError.New(a.argument.Name, a.value)
}

func (a ArgumentValue) AsBool() (bool, error) {
	value, err := strconv.ParseBool(a.value)

	return value, err
}

func (a ArgumentValue) AsFloat() (float64, error) {
	value, err := strconv.ParseFloat(a.value, 64)

	return value, err
}

func (a ArgumentValue) AsRangedFloat(min float64, max float64) (float64, error) {
	value, err := a.AsFloat()
	if err != nil {
		return 0, InvalidArgumentValueError.New(a.argument.Name, a.value).CausedBy(err)
	}

	if min <= value && value <= max {
		return value, nil
	}
	return 0, InvalidArgumentValueError.New(a.argument.Name, a.value)
}

func (a ArgumentValue) AsStringList() ([]string, error) {
	return strings.Split(a.value, ","), nil
}

func (a ArgumentValue) AsIntList() ([]int, error) {
	var result []int
	for _, item := range strings.Split(a.value, ",") {
		value, _ := strconv.Atoi(item)
		result = append(result, value)
	}

	return result, nil
}

func (a ArgumentValue) AsFloatList() ([]float64, error) {
	var result []float64
	for _, item := range strings.Split(a.value, ",") {
		value, _ := strconv.ParseFloat(item, 64)
		result = append(result, value)
	}

	return result, nil
}

func (a ArgumentValue) AsIntEnum(enumMap map[string]int) (int, error) {
	if value, ok := enumMap[a.value]; ok {
		return value, nil
	}

	return -1, InvalidArgumentValueError.New(a.argument.Name, a.value)
}

func (a ArgumentValue) AsStringEnum(enumMap map[string]string) (string, error) {
	if value, ok := enumMap[a.value]; ok {
		return value, nil
	}

	return "", InvalidArgumentValueError.New(a.argument.Name, a.value)
}

func (a ArgumentValue) IsEnabled() bool {
	if a.value == "1" {
		return true
	}

	return false
}