package best_time_to_buy_and_sell_stock

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	cases := []struct {
		prices []int
		want   int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{2, 4, 1}, 2},
	}

	for _, c := range cases {
		got := maxProfit(c.prices)
		if got != c.want {
			t.Errorf("got %d, wanted %d", got, c.want)
		}
	}
}
