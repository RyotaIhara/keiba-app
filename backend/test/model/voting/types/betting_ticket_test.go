package types_test

import (
	"reflect"
	"testing"

	"keiba-app-backend/model/voting/types"
)

// TestNewBettingTicket 買い目を生成して選択内容を取得できることを検証する。
func TestNewBettingTicket(t *testing.T) {
	ticket, err := types.NewBettingTicket([][]int{{1}, {2, 3}, {3, 4, 5, 6}})
	if err != nil {
		t.Fatalf("NewBettingTicket returned error: %v", err)
	}
	want := [][]int{{1}, {2, 3}, {3, 4, 5, 6}}
	if selections, err := ticket.Selections(); err != nil || !reflect.DeepEqual(selections, want) {
		t.Fatalf("Selections() = %#v, %v; want %#v, nil", selections, err, want)
	}
}

// TestParseBettingTicket DB保存形式から買い目を復元できることを検証する。
func TestParseBettingTicket(t *testing.T) {
	got, err := types.ParseBettingTicket("1-2,3-3,4,5,6")
	if err != nil {
		t.Fatalf("ParseBettingTicket returned error: %v", err)
	}
	want := [][]int{{1}, {2, 3}, {3, 4, 5, 6}}
	if selections, err := got.Selections(); err != nil || !reflect.DeepEqual(selections, want) {
		t.Fatalf("Selections() = %#v, %v; want %#v, nil", selections, err, want)
	}
}

// TestBettingTicketScan DB取得時に買い目を復元できることを検証する。
func TestBettingTicketScan(t *testing.T) {
	var ticket types.BettingTicket
	if err := ticket.Scan([]byte("1-2-3")); err != nil {
		t.Fatalf("Scan() returned error: %v", err)
	}
	want := [][]int{{1}, {2}, {3}}
	if selections, err := ticket.Selections(); err != nil || !reflect.DeepEqual(selections, want) {
		t.Fatalf("Selections() = %#v, %v; want %#v, nil", selections, err, want)
	}
}

// TestBettingTicketRejectsInvalidValues 不正な買い目を拒否することを検証する。
func TestBettingTicketRejectsInvalidValues(t *testing.T) {
	for _, value := range []string{"", "1-", "1,,2", "1-a", "0-2"} {
		if _, err := types.ParseBettingTicket(value); err == nil {
			t.Errorf("ParseBettingTicket(%q) returned nil error", value)
		}
	}
}
