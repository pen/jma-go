package area_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/pen/jma-go/area"
)

func TestParse(t *testing.T) {
	t.Parallel()

	//nolint:lll
	testCases := []struct {
		name    string
		json    string
		want    []*area.Area
		wantErr bool
	}{
		{name: "エラー", json: `abc`, want: nil, wantErr: true},
		{name: "空", json: `{}`, want: nil, wantErr: false},
		{
			name: "簡単", wantErr: false,
			json: `{"centers": {}, "class10s": {}, "class15s": {},
				"offices": {"456": {"parent": "123", "officeName": "札幌", "name": "北海道", "enName": "Hokkie", "kana": "ほかいど", "children": ["789"]}},
				"class20s": {"7890": {"parent": "456", "officeName": "札幌", "name": "網走", "enName": "Absr", "kana": "あばしり"}}
			}`,
			want: []*area.Area{
				{Class: "office", Code: "456", ParentCode: "123", OfficeName: "札幌", Name: "北海道", NameEn: "Hokkie", NameKana: "ほかいど"},
				{Class: "class20", Code: "7890", ParentCode: "456", OfficeName: "札幌", Name: "網走", NameEn: "Absr", NameKana: "あばしり"},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := area.Parse(strings.NewReader(tc.json))
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
