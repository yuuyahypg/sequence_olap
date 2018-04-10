package multi_match_recognize

import (
	"fmt"
)

type Edge struct {
	To         string
	Action     string
	Name       string
	Conditions []*Condition
}

type Condition struct {
	SelfColumn   string
	ColumnType   string
	Param        interface{}
	TargetName   string
	TargetColumn string
	Operator     string
}

func (e *Edge) CheckEdge(row *Row, cs *CurrentState) (*CurrentState, bool) {
	flag := true
	for _, c := range e.Conditions {
		if ok := c.Operate(row, cs); !ok {
			flag = false
			break
		}
	}

	if flag {
		return nfa.GetNewCurrentState(), true
	}
	return &CurrentState{}, false
}

func (c *Condition) Operate(row *Row, cs *CurrentState) bool {
	switch c.ColumnType {
	case "int":
		return c.OperateOnInt(row, cs)
	case "float":
		return c.OperateOnFloat(row, cs)
	case "string":
		return c.OperateOnString(row, cs)
	}

	return false
}

func (c *Condition) OperateOnInt(row *Row, cs *CurrentState) bool {
	s, ok := row.Record[c.SelfColumn].(int)
	if !ok {
		fmt.Println("self column is not int")
		return false
	}

	var t int
	if c.TargetName != "" {
		t, ok = cs.MatchBuffer[c.TargetName][0].Record[c.TargetColumn].(int)
		if !ok {
			fmt.Println("target column is not int")
			return false
		}
	}

	var p int
	if c.Param != nil {
		p, ok = c.Param.(int)
		if !ok {
			fmt.Println("parameter is not int")
			return false
		}
	}

	switch c.Operator {
	case "==":
		if s == t {
			return true
		} else {
			return false
		}
	case "!=":
		if s != t {
			return true
		} else {
			return false
		}
	case "<":
		if s < t {
			return true
		} else {
			return false
		}
	case ">":
		if s > t {
			return true
		} else {
			return false
		}
	case "length":
		if s-t <= p {
			return true
		} else {
			return false
		}
	}
	return false
}

func (c *Condition) OperateOnFloat(row *Row, cs *CurrentState) bool {
	s, ok := row.Record[c.SelfColumn].(float32)
	if !ok {
		fmt.Println("self column is not float")
		return false
	}

	var t float32
	if c.TargetName != "" {
		t, ok = cs.MatchBuffer[c.TargetName][0].Record[c.TargetColumn].(float32)
		if !ok {
			fmt.Println("target column is not float")
			return false
		}
	}

	var p float32
	if c.Param != nil {
		p, ok = c.Param.(float32)
		if !ok {
			fmt.Println("parameter is not float")
			return false
		}
	}

	switch c.Operator {
	case "==":
		if s == t {
			return true
		} else {
			return false
		}
	case "!=":
		if s != t {
			return true
		} else {
			return false
		}
	case "<":
		if s < t {
			return true
		} else {
			return false
		}
	case ">":
		if s > t {
			return true
		} else {
			return false
		}
	case "length":
		if s-t <= p {
			return true
		} else {
			return false
		}
	}
	return false
}

func (c *Condition) OperateOnString(row *Row, cs *CurrentState) bool {
	s, ok := row.Record[c.SelfColumn].(string)
	if !ok {
		fmt.Println("self column is not string")
		return false
	}

	var t string
	if c.TargetName != "" {
		t, ok = cs.MatchBuffer[c.TargetName][0].Record[c.TargetColumn].(string)
		if !ok {
			fmt.Println("target column is not string")
			return false
		}
	}

	switch c.Operator {
	case "==":
		if s == t {
			return true
		} else {
			return false
		}
	case "!=":
		if s != t {
			return true
		} else {
			return false
		}
	}
	return false
}

var edges = map[string][]*Edge{
	"init": []*Edge{
		&Edge{
			To:         "state_x",
			Action:     "take",
			Name:       "x",
			Conditions: []*Condition{},
		},
		&Edge{
			To:         "init",
			Action:     "ignore",
			Name:       "x",
			Conditions: []*Condition{},
		},
	},
	"state_x": []*Edge{
		&Edge{
			To:     "state_y1",
			Action: "take",
			Name:   "y",
			Conditions: []*Condition{
				&Condition{
					SelfColumn:   "location",
					ColumnType:   "string",
					TargetName:   "x",
					TargetColumn: "location",
					Operator:     "!=",
				},
				&Condition{
					SelfColumn:   "rid",
					ColumnType:   "int",
					Param:        10,
					TargetName:   "x",
					TargetColumn: "rid",
					Operator:     "length",
				},
			},
		},
		&Edge{
			To:     "state_w",
			Action: "take",
			Name:   "w",
			Conditions: []*Condition{
				&Condition{
					SelfColumn:   "location",
					ColumnType:   "string",
					TargetName:   "x",
					TargetColumn: "location",
					Operator:     "!=",
				},
				&Condition{
					SelfColumn:   "rid",
					ColumnType:   "int",
					Param:        10,
					TargetName:   "x",
					TargetColumn: "rid",
					Operator:     "length",
				},
			},
		},
		&Edge{
			To:     "state_x",
			Action: "ignore",
			Name:   "x",
			Conditions: []*Condition{
				&Condition{
					SelfColumn:   "rid",
					ColumnType:   "int",
					Param:        10,
					TargetName:   "x",
					TargetColumn: "rid",
					Operator:     "length",
				},
			},
		},
	},
	"state_y1": []*Edge{
		&Edge{
			To:     "state_z",
			Action: "take",
			Name:   "z",
			Conditions: []*Condition{
				&Condition{
					SelfColumn:   "location",
					ColumnType:   "string",
					TargetName:   "y",
					TargetColumn: "location",
					Operator:     "!=",
				},
				&Condition{
					SelfColumn:   "rid",
					ColumnType:   "int",
					Param:        10,
					TargetName:   "x",
					TargetColumn: "rid",
					Operator:     "length",
				},
			},
		},
		&Edge{
			To:     "state_y1",
			Action: "ignore",
			Name:   "x",
			Conditions: []*Condition{
				&Condition{
					SelfColumn:   "rid",
					ColumnType:   "int",
					Param:        10,
					TargetName:   "x",
					TargetColumn: "rid",
					Operator:     "length",
				},
			},
		},
	},
	"state_w": []*Edge{
		&Edge{
			To:     "state_y2",
			Action: "take",
			Name:   "y",
			Conditions: []*Condition{
				&Condition{
					SelfColumn:   "location",
					ColumnType:   "string",
					TargetName:   "w",
					TargetColumn: "location",
					Operator:     "!=",
				},
				&Condition{
					SelfColumn:   "location",
					ColumnType:   "string",
					TargetName:   "x",
					TargetColumn: "location",
					Operator:     "!=",
				},
				&Condition{
					SelfColumn:   "rid",
					ColumnType:   "int",
					Param:        10,
					TargetName:   "x",
					TargetColumn: "rid",
					Operator:     "length",
				},
			},
		},
		&Edge{
			To:     "state_w",
			Action: "ignore",
			Name:   "x",
			Conditions: []*Condition{
				&Condition{
					SelfColumn:   "rid",
					ColumnType:   "int",
					Param:        10,
					TargetName:   "x",
					TargetColumn: "rid",
					Operator:     "length",
				},
			},
		},
	},
}
