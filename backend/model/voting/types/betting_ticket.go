// Package types
package types

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
)

// BettingTicket は、着順ごとの馬番による買い目を表す。
// DBには着順を「-」、同じ着順の複数指定を「,」で区切って保存する。
//
// 例として、[][]int{{1}, {2}, {3}} は「1-2-3」となり、
// [][]int{{1}, {2, 3, 4}, {3, 4, 5, 6}} は
// 「1-2,3,4-3,4,5,6」となる。
type BettingTicket string

// NewBettingTicket 指定された馬番から買い目を生成する。
func NewBettingTicket(selections [][]int) (BettingTicket, error) {
	if err := validateSelections(selections); err != nil {
		return "", err
	}

	positions := make([]string, len(selections))
	for i, selection := range selections {
		horseNumbers := make([]string, len(selection))
		for j, horseNumber := range selection {
			horseNumbers[j] = strconv.Itoa(horseNumber)
		}
		positions[i] = strings.Join(horseNumbers, ",")
	}

	return BettingTicket(strings.Join(positions, "-")), nil
}

// ParseBettingTicket DB保存形式の文字列から買い目を生成する。
func ParseBettingTicket(value string) (BettingTicket, error) {
	selections, err := parseSelections(value)
	if err != nil {
		return "", err
	}
	return NewBettingTicket(selections)
}

// Selections 買い目を着順ごとの馬番に変換する。
func (ticket BettingTicket) Selections() ([][]int, error) {
	return parseSelections(string(ticket))
}

// Scan DBから取得した値を買い目に変換する。
func (ticket *BettingTicket) Scan(value any) error {
	if ticket == nil {
		return fmt.Errorf("cannot scan betting ticket into nil receiver")
	}

	switch value := value.(type) {
	case nil:
		return fmt.Errorf("betting ticket cannot be NULL")
	case string:
		parsed, err := ParseBettingTicket(value)
		if err != nil {
			return err
		}
		*ticket = parsed
		return nil
	case []byte:
		parsed, err := ParseBettingTicket(string(value))
		if err != nil {
			return err
		}
		*ticket = parsed
		return nil
	default:
		return fmt.Errorf("cannot scan %T into betting ticket", value)
	}
}

// toString 買い目をDB保存形式の文字列に変換する。
func (ticket BettingTicket) toString() string {
	return string(ticket)
}

// toValue 買い目をDB保存用の値に変換する。
func (ticket BettingTicket) toValue() (driver.Value, error) {
	if _, err := ticket.Selections(); err != nil {
		return nil, err
	}
	return ticket.toString(), nil
}

// parseSelections DB保存形式の文字列を着順ごとの馬番に変換する。
func parseSelections(value string) ([][]int, error) {
	if strings.TrimSpace(value) == "" {
		return nil, fmt.Errorf("betting ticket cannot be empty")
	}

	positions := strings.Split(value, "-")
	selections := make([][]int, len(positions))
	for i, position := range positions {
		if position == "" {
			return nil, fmt.Errorf("betting ticket contains an empty position")
		}

		horseNumbers := strings.Split(position, ",")
		selections[i] = make([]int, len(horseNumbers))
		for j, horseNumber := range horseNumbers {
			if horseNumber == "" {
				return nil, fmt.Errorf("betting ticket contains an empty horse number")
			}
			parsed, err := strconv.Atoi(horseNumber)
			if err != nil || parsed <= 0 {
				return nil, fmt.Errorf("invalid horse number %q", horseNumber)
			}
			selections[i][j] = parsed
		}
	}
	return selections, nil
}

// validateSelections 着順ごとの馬番を検証する。
func validateSelections(selections [][]int) error {
	if len(selections) == 0 {
		return fmt.Errorf("betting ticket must contain at least one position")
	}
	for _, selection := range selections {
		if len(selection) == 0 {
			return fmt.Errorf("betting ticket positions cannot be empty")
		}
		for _, horseNumber := range selection {
			if horseNumber <= 0 {
				return fmt.Errorf("horse number must be positive: %d", horseNumber)
			}
		}
	}
	return nil
}
