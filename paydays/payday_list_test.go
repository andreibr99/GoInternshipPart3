package paydays

import (
	"github.com/google/go-cmp/cmp"
	"testing"
	"time"
)

func TestPaydayList(t *testing.T) {
	type testCase struct {
		name   string
		time   time.Time
		payday int
		want   []string
	}
	var tests = []testCase{
		{
			name:   "payday passed for this month",
			time:   time.Date(2023, time.February, 21, 0, 0, 0, 0, time.Local),
			payday: 15,
			want: []string{
				"2023-03-15", "2023-04-14", "2023-05-15", "2023-06-15", "2023-07-14",
				"2023-08-15", "2023-09-15", "2023-10-13", "2023-11-15", "2023-12-15",
			},
		},
		{
			name:   "payday is later this month",
			time:   time.Date(2023, time.February, 12, 0, 0, 0, 0, time.Local),
			payday: 15,
			want: []string{
				"2023-02-15", "2023-03-15", "2023-04-14", "2023-05-15", "2023-06-15", "2023-07-14",
				"2023-08-15", "2023-09-15", "2023-10-13", "2023-11-15", "2023-12-15",
			},
		},
		{
			name:   "payday is on the last available day of the month",
			time:   time.Date(2023, time.February, 12, 0, 0, 0, 0, time.Local),
			payday: 31,
			want: []string{
				"2023-02-28", "2023-03-31", "2023-04-28", "2023-05-31", "2023-06-30", "2023-07-31",
				"2023-08-31", "2023-09-29", "2023-10-31", "2023-11-30", "2023-12-29",
			},
		},
		{
			name:   "payday is set to 1st of each month",
			time:   time.Date(2023, time.February, 12, 0, 0, 0, 0, time.Local),
			payday: 1,
			want: []string{
				"2023-03-01", "2023-03-31", "2023-05-01", "2023-06-01", "2023-06-30",
				"2023-08-01", "2023-09-01", "2023-09-29", "2023-11-01", "2023-12-01",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetPaydaysList(tt.time, tt.payday)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Function result: %v, expected result: %v", got, tt.want)
			}
		})
	}
}
