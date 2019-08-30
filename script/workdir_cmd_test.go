package script

import (
	"fmt"
	"testing"
)

func TestCommandWORKDIR(t *testing.T) {
	tests := []commandTest{
		{
			name: "WORKDIR with single arg",
			source: func() string {
				return "WORKDIR foo/bar"
			},
			script: func(s *Script) error {
				dirs := s.Preambles[CmdWorkDir]
				if len(dirs) != 1 {
					return fmt.Errorf("Script has unexpected number of WORKDIR %d", len(dirs))
				}
				wdCmd, ok := dirs[0].(*WorkdirCommand)
				if !ok {
					return fmt.Errorf("Unexpected type %T in script", dirs[0])
				}
				if wdCmd.Dir() != "foo/bar" {
					return fmt.Errorf("WORKDIR has unexpected directory %s", wdCmd.Dir())
				}
				return nil
			},
		},
		{
			name: "Multiple WORKDIRs",
			source: func() string {
				return "WORKDIR foo/bar\nWORKDIR bazz/buzz"
			},
			script: func(s *Script) error {
				dirs := s.Preambles[CmdWorkDir]
				if len(dirs) != 1 {
					return fmt.Errorf("Script has unexpected number of WORKDIR %d", len(dirs))
				}
				wdCmd, ok := dirs[0].(*WorkdirCommand)
				if !ok {
					return fmt.Errorf("Unexpected type %T in script", dirs[0])
				}
				if wdCmd.Dir() != "bazz/buzz" {
					return fmt.Errorf("WORKDIR has unexpected directory %s", wdCmd.Dir())
				}
				return nil
			},
		},
		{
			name: "WORKDIR with multiple args",
			source: func() string {
				return "WORKDIR foo/bar bazz/buzz"
			},
			shouldFail: true,
		},
		{
			name: "WORKDIR with no args",
			source: func() string {
				return "WORKDIR"
			},
			shouldFail: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			runCommandTest(t, test)
		})
	}
}