package utils

import (
	"reflect"
	"testing"
)

func TestParseArgs(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name             string
		args             args
		wantWorkersCount int
		wantCmdSource    string
		wantCommands     []string
	}{
		{
			name:             "empty args",
			args:             args{args: []string{}},
			wantWorkersCount: 3,
			wantCmdSource:    "files",
			wantCommands:     []string{},
		},
		{
			name: "Only workers count",
			args: args{args: []string{
				"--workers=2",
			}},
			wantWorkersCount: 2,
			wantCmdSource:    "files",
			wantCommands:     []string{},
		},
		{
			name: "Invalid workers count",
			args: args{args: []string{
				"--workers=foo",
			}},
			wantWorkersCount: 3,
			wantCmdSource:    "files",
			wantCommands:     []string{},
		},
		{
			name: "Only workers count",
			args: args{args: []string{
				"--workers=-2",
			}},
			wantWorkersCount: 3,
			wantCmdSource:    "files",
			wantCommands:     []string{},
		},
		{
			name: "Only source (args)",
			args: args{args: []string{
				"--src=args",
			}},
			wantWorkersCount: 3,
			wantCmdSource:    "args",
			wantCommands:     []string{},
		},
		{
			name: "Only source (files)",
			args: args{args: []string{
				"--src=files",
			}},
			wantWorkersCount: 3,
			wantCmdSource:    "files",
			wantCommands:     []string{},
		},
		{
			name: "Invalid source",
			args: args{args: []string{
				"--src=foo",
			}},
			wantWorkersCount: 3,
			wantCmdSource:    "files",
			wantCommands:     []string{},
		},
		{
			name: "Command line args",
			args: args{args: []string{
				"TEST",
				"TEST_YET",
			}},
			wantWorkersCount: 3,
			wantCmdSource:    "files",
			wantCommands: []string{
				"TEST",
				"TEST_YET",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWorkersCount, gotCmdSource, gotCommands := ParseArgs(tt.args.args)
			if gotWorkersCount != tt.wantWorkersCount {
				t.Errorf("ParseArgs() gotWorkersCount = %v, want %v", gotWorkersCount, tt.wantWorkersCount)
			}
			if gotCmdSource != tt.wantCmdSource {
				t.Errorf("ParseArgs() gotCmdSource = %v, want %v", gotCmdSource, tt.wantCmdSource)
			}
			if !reflect.DeepEqual(gotCommands, tt.wantCommands) {
				t.Errorf("ParseArgs() gotCommands = %v, want %v", gotCommands, tt.wantCommands)
			}
		})
	}
}
