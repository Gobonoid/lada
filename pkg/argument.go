package lada

import (
	"strconv"
	"strings"
)

type Argument struct {
	Name        string
	Wildcard    bool
	Description string
	Value 		string
}


func (a Argument) AsInt() (int, error) {
	return strconv.Atoi(a.Value)
}

func (a Argument) AsRangedInt(min int, max int) (int, error) {
	value, err := a.AsInt()
	if err != nil {
		return 0, InvalidArgumentValueError.New(a.Name, a.Value).CausedBy(err)
	}

	if min <= value && value <= max {
		return value, nil
	}
	return 0, InvalidArgumentValueError.New(a.Name, a.Value)
}

func (a Argument) AsBool() (bool, error) {
	return strconv.ParseBool(a.Value)
}

func (a Argument) AsFloat() (float64, error) {
	return strconv.ParseFloat(a.Value, 64)
}

func (a Argument) AsRangedFloat(min float64, max float64) (float64, error) {
	value, err := a.AsFloat()
	if err != nil {
		return 0, InvalidArgumentValueError.New(a.Name, a.Value).CausedBy(err)
	}

	if min <= value && value <= max {
		return value, nil
	}
	return 0, InvalidArgumentValueError.New(a.Name, a.Value)
}

func (a Argument) AsStringList() []string {
	if !a.Wildcard {
		return strings.Split(a.Value, ",")
	}
	return strings.Split(a.Value, " ")
}

func (a Argument) AsIntList() ([]int, error) {
	var result []int
	for _, item := range a.AsStringList() {
		value, _ := strconv.Atoi(item)
		result = append(result, value)
	}

	return result, nil
}

func (a Argument) AsFloatList() ([]float64, error) {
	var result []float64
	for _, item := range a.AsStringList() {
		value, _ := strconv.ParseFloat(item, 64)
		result = append(result, value)
	}

	return result, nil
}

func (a Argument) AsIntEnum(enumMap map[string]int) (int, error) {
	if value, ok := enumMap[a.Value]; ok {
		return value, nil
	}

	return -1, InvalidArgumentValueError.New(a.Name, a.Value)
}

func (a Argument) AsStringEnum(enumMap map[string]string) (string, error) {
	if value, ok := enumMap[a.Value]; ok {
		return value, nil
	}

	return "", InvalidArgumentValueError.New(a.Name, a.Value)
}
