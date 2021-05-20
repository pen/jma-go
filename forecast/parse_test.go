package forecast_test

import (
	"encoding/json"
	"os"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/pen/jma-go/forecast"
)

func TestParse(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		json    string
		want    []*forecast.Forecast
		wantErr bool
	}{
		{name: "エラー", json: `abc`, want: nil, wantErr: true},
		{name: "空", json: `[]`, want: []*forecast.Forecast{}, wantErr: false},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := forecast.Parse(strings.NewReader(tc.json))
			if (err != nil) != tc.wantErr {
				t.Errorf("Parse(): error: %v, wantErr: %v", err, tc.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Parse(): got: %+v, want: %+v", got, tc.want)
			}
		})
	}
}

func TestParse_File(t *testing.T) { //nolint:funlen
	t.Parallel()

	name := "130000"

	//nolint:lll
	testCases := []struct {
		name  string
		index int
		want  string
	}{
		{
			name:  "サンプル10",
			index: 10,
			want:  `{"Area":{"Code":"130010","Name":"東京地方"},"ReportedAt":"2021-05-20T17:00:00+09:00","ComesAt":"2021-05-26T00:00:00+09:00","Weather":{"Code":"201","Text":null},"Wind":null,"Wave":null,"Precipitation":{"Probability":30,"Reliability":"A"},"Temperature":null}`,
		},
		{
			name:  "サンプル20",
			index: 20,
			want:  `{"Area":{"Code":"130030","Name":"伊豆諸島南部"},"ReportedAt":"2021-05-20T17:00:00+09:00","ComesAt":"2021-05-20T18:00:00+09:00","Weather":null,"Wind":null,"Wave":null,"Precipitation":{"Probability":60,"Reliability":null},"Temperature":null}`,
		},
	}

	f, err := os.Open("testdata/" + name + ".json")
	if err != nil {
		t.Errorf("Open(): error: %v", err)
		return
	}

	forecasts, err := forecast.Parse(f)
	if err != nil {
		t.Errorf("Parse(): error: %v", err)
		return
	}

	gotLen := len(forecasts)
	wantLen := 71

	if gotLen != wantLen {
		t.Errorf("Parse(): length: got: %+v, want: %+v", gotLen, wantLen)
	}

	sort.Slice(forecasts, func(i, j int) bool {
		if forecasts[i].Area.Code < forecasts[j].Area.Code {
			return true
		}
		if forecasts[i].Area.Code > forecasts[j].Area.Code {
			return false
		}
		return forecasts[i].ComesAt.Before(forecasts[j].ComesAt)
	})

	forecastStr := func(f *forecast.Forecast) string {
		b, err := json.Marshal(f)
		if err != nil {
			return err.Error()
		}

		return string(b)
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := forecastStr(forecasts[tc.index])
			if got != tc.want {
				t.Errorf("Parse(): got: %s, want: %s", got, tc.want)
			}
		})
	}
}
