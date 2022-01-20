package analyzer

import (
	"reflect"
	"strings"
	"testing"
)

func TestAnalyzer_AddLine(t *testing.T) {
	type args struct {
		line string
	}
	collector := New()
	tests := []struct {
		name     string
		a        *Analyzer
		args     args
		wantText string
	}{
		{
			name: "Add 1 line",
			a:    &collector,
			args: args{
				line: "First line",
			},
			wantText: "First line",
		},
		{
			name: "Add 2 line",
			a:    &collector,
			args: args{
				line: "Second line",
			},
			wantText: "First line\nSecond line",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.AddLine(tt.args.line)
			gotText := tt.a.GetText()
			if gotText != tt.wantText {
				t.Errorf("AddLine() gotText = %v, want = %v", gotText, tt.wantText)
			}
		})
	}
}

func TestAnalyzer_Restart(t *testing.T) {
	line := "Quick brown fox jumped over the lazy dog then jumped again back!"
	collector := New()
	collector.AddLine(line)
	gotStats := collector.GetStats()
	wantStats1 := Stats{
		CharCount:            53,
		TopWord:              "jumped",
		TopWordOccurrences:   2,
		NonLettersCount:      1,
		WhitespacesCount:     11,
		TopSymbol:            rune(101),
		TopSymbolOccurrences: 5,
	}

	collector.Restart()
	gotStats = collector.GetStats()
	wantStats2 := Stats{
		CharCount:            0,
		TopWord:              "",
		TopWordOccurrences:   0,
		NonLettersCount:      0,
		WhitespacesCount:     0,
		TopSymbol:            ' ',
		TopSymbolOccurrences: 0,
	}
	if !reflect.DeepEqual(gotStats, wantStats2) {
		t.Errorf("Restart() gotStats = %v, want %v", gotStats, wantStats2)
	}

	collector.AddLine(line)
	gotStats = collector.GetStats()
	if !reflect.DeepEqual(gotStats, wantStats1) {
		t.Errorf("gotStats = %v, want %v (after Restart and AddLine again)", gotStats, wantStats1)
	}
}

func BenchmarkAnalyzer_analyze(b *testing.B) {
	collector := New()
	for i := 0; i < b.N; i++ {
		collector.analyze(strings.Repeat("Quick brown fox jumped over the lazy dog", b.N))
	}
}
