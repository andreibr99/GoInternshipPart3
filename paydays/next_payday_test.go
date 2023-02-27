package paydays

import (
	"testing"
	"time"
)

func TestDaysUntilNextPayday(t *testing.T) {
	type testCase struct {
		name     string
		time     time.Time
		payday   int
		wantDays int
		wantDate string
	}

	var tests = []testCase{
		{
			name:     "payday is later this month, not in a weekend day",
			time:     time.Date(2023, time.July, 9, 0, 0, 0, 0, time.Local),
			payday:   15,
			wantDays: 5,
			wantDate: "2023-07-14",
		},
		{
			name:     "payday is today, not on a weekend",
			time:     time.Date(2023, time.February, 21, 0, 0, 0, 0, time.Local),
			payday:   21,
			wantDays: 0,
			wantDate: "2023-02-21",
		},
		{
			name: "payday is today, but today is a weekend day",
			// that means that the payday was this friday, so it should compute the date for the next month
			time:     time.Date(2023, time.March, 25, 0, 0, 0, 0, time.Local),
			payday:   25,
			wantDays: 30,
			wantDate: "2023-04-25",
		},
		{
			name: "payday is today, but today is a weekend day, and next month payday is on a weekend day too",
			// same as above, but the next month date is not available, so it should put the closest friday to that date
			time:     time.Date(2023, time.February, 25, 0, 0, 0, 0, time.Local),
			payday:   25,
			wantDays: 28,
			wantDate: "2023-03-24",
		},
		{
			name:     "payday is later this month, not in a weekend day",
			time:     time.Date(2023, time.February, 5, 0, 0, 0, 0, time.Local),
			payday:   15,
			wantDays: 10,
			wantDate: "2023-02-15",
		},
		{
			name:     "payday passed this month, so it should be computed for the next month",
			time:     time.Date(2023, time.February, 21, 0, 0, 0, 0, time.Local),
			payday:   15,
			wantDays: 22,
			wantDate: "2023-03-15",
		},
		{
			name: "payday passed this month, computed for the next month, but next month payday is in a weekend",
			// this means that the next month payday should be on the closest friday to that date
			time:     time.Date(2023, time.February, 21, 0, 0, 0, 0, time.Local),
			payday:   5,
			wantDays: 10,
			wantDate: "2023-03-03",
		},
		{
			name:     "payday is on a weekend day, expected payday on the closest friday",
			time:     time.Date(2023, time.March, 20, 0, 0, 0, 0, time.Local),
			payday:   26,
			wantDays: 4,
			wantDate: "2023-03-24",
		},
		{
			name:     "payday does not exist in that month, expected payday on the closest available date",
			time:     time.Date(2023, time.February, 21, 0, 0, 0, 0, time.Local),
			payday:   30,
			wantDays: 7,
			wantDate: "2023-02-28",
		},
		{
			name:     "payday is on the 1st of the next month, but it's a weekend day",
			time:     time.Date(2023, time.March, 27, 0, 0, 0, 0, time.Local),
			payday:   1,
			wantDays: 4,
			wantDate: "2023-03-31",
		},
		{
			name: "payday is today, on the 1st of this month, but it's a weekend day",
			// this means that the payday was on the last friday of the previous month, so the next payday date
			// should be 1st of the next month if available
			time:     time.Date(2023, time.April, 1, 0, 0, 0, 0, time.Local),
			payday:   1,
			wantDays: 30,
			wantDate: "2023-05-01",
		},
		{
			name: "payday is this sunday, today is saturday",
			// this means that the payday was on friday, so the next payday date should be computed for the next month
			time:     time.Date(2023, time.April, 15, 0, 0, 0, 0, time.Local),
			payday:   16,
			wantDays: 31,
			wantDate: "2023-05-16",
		},
		{
			name: "payday is on 31st, 31st does not exist in the month, and today is the last day of the month",
			// this means that the payday was on friday, so the next payday date should be computed for the next month
			time:     time.Date(2023, time.February, 28, 0, 0, 0, 0, time.Local),
			payday:   31,
			wantDays: 0,
			wantDate: "2023-02-28",
		},
		{
			name:     "payday passed in the last day of the month, and the pay day date should be next year",
			time:     time.Date(2023, time.December, 25, 0, 0, 0, 0, time.Local),
			payday:   14,
			wantDays: 18,
			wantDate: "2024-01-12", // 13,14 Jan are on a weekend
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDays, gotDate := CalculateNextPayday(tt.time, tt.payday)
			if gotDays != tt.wantDays || gotDate != tt.wantDate {
				t.Errorf("Function result: %v days, date: %v, expected result: %v days, date: %v",
					gotDays, gotDate, tt.wantDays, tt.wantDate)
			}
		})
	}
}
