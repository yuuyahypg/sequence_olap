package multi_match_recognize

type FinalState struct {
	Name          string
	OutputColumns []string
	Parents       map[string]*Parent
}

type Parent struct {
	Type         string
	MatchColumns []string
}

var finalStates = map[string]*FinalState{
	"state_y1": &FinalState{
		Name:          "P1",
		OutputColumns: []string{"x", "y"},
		Parents:       map[string]*Parent{},
	},
	"state_z": &FinalState{
		Name:          "P2",
		OutputColumns: []string{"x", "y", "z"},
		Parents: map[string]*Parent{
			"P1": &Parent{
				Type:         "forward",
				MatchColumns: []string{"x", "y"},
			},
		},
	},
	"state_y2": &FinalState{
		Name:          "P3",
		OutputColumns: []string{"x", "w", "y"},
		Parents: map[string]*Parent{
			"P1": &Parent{
				Type:         "backward",
				MatchColumns: []string{"x", "y"},
			},
		},
	},
}
