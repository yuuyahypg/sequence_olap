package multi_match_recognize

type CurrentState struct {
	State       string
	MatchBuffer map[string][]*Row
	Oid         map[string]int
}

type CurrentStateOption func(*CurrentState) error

func MatchBuffer(matchBuffer map[string][]*Row) CurrentStateOption {
	return func(cs *CurrentState) error {
		for key, value := range matchBuffer {
			cs.MatchBuffer[key] = []*Row{}
			copy(cs.MatchBuffer[key], value)
		}
		return nil
	}
}

func Oid(oid map[string]int) CurrentStateOption {
	return func(cs *CurrentState) error {
		for key, value := range oid {
			cs.Oid[key] = value
		}
		return nil
	}
}

func NewCurrentState(state string, options ...CurrentStateOption) *CurrentState {
	cs := &CurrentState{
		State:       state,
		MatchBuffer: map[string][]*Row{},
		Oid:         map[string]int{},
	}

	for _, option := range options {
		option(cs)
	}

	return cs
}
