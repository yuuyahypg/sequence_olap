package multi_match_recognize

type NondeterministicFiniteAutomata struct {
	States             []string
	Edges              map[string][]*Edge
	InitState          string
	CurrentStates      []*CurrentState
	FinalStates        map[string]*FinalState
	MatchBufferColumns []string
	MiddleBuffer       *MiddleBuffer
	ResultBuffer       *ResultBuffer
}

type ResultBuffer struct {
	Count  int
	Buffer []map[string]interface{}
}

type Row struct {
	Rid    int
	Record map[string]interface{}
}

func NewNondeterministicFiniteAutomata() *NondeterministicFiniteAutomata {
	states := []string{"init", "state_x", "state_y1", "state_z", "state_w", "state_y2"}

	nfa := &NondeterministicFiniteAutomata{
		States:             states,
		Edges:              edges,
		InitState:          "init",
		CurrentStates:      []*CurrentState{},
		FinalStates:        finalStates,
		MatchBufferColumns: []string{"x", "w", "y", "z"},
		MiddleBuffer:       middleBuffer,
		ResultBuffer: &ResultBuffer{
			Count:  0,
			Buffer: []map[string]interface{}{},
		},
	}

	return nfa
}

func (nfa *NondeterministicFiniteAutomata) Run(table []*Row) bool {

	return true
}
