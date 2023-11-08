package serendip

import (
	"fmt"
	"testing"
)

func TestGenerateRandomArticleMessage(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "standard",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateRandomArticleMessage()
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateRandomArticleMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != "" {
				fmt.Println(got)
			} else {
				t.Errorf("The got value is empty")
				return
			}
		})
	}
}

func TestGenerateSearchResultMessage(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "standard",
			args: args{"Wikipedia"},
			want: `Wikipedia
<https://ja.wikipedia.org/wiki/Wikipedia>
Wikipedia Zero
<https://ja.wikipedia.org/wiki/Wikipedia%20Zero>
Wikipedia日本語版
<https://ja.wikipedia.org/wiki/Wikipedia%E6%97%A5%E6%9C%AC%E8%AA%9E%E7%89%88>
`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateSearchResultMessage(tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateSearchResultMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GenerateSearchResultMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
