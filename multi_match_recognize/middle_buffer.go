package multi_match_recognize

type MiddleBuffer struct {
	JoinPlans []*JoinPlan
	Buffers   map[string]*Buffer
}

type JoinPlan struct {
	Left         string
	Right        string
	MatchColumns []string
	ReferenceOid []string
}

type Buffer struct {
	Count int
	Rows  []*BufferRow
}

type BufferRow struct {
	Rid     map[string][]int
	Columns map[string][]map[string]interface{}
}

var middleBuffer = &MiddleBuffer{
	JoinPlans: []*JoinPlan{
		&JoinPlan{
			Left:         "P1",
			Right:        "P2",
			MatchColumns: []string{"x", "y"},
			ReferenceOid: []string{"P1_oid"},
		},
		&JoinPlan{
			Left:         "P1",
			Right:        "P3",
			MatchColumns: []string{"x", "y"},
			ReferenceOid: []string{"P1_oid"},
		},
	},
	Buffers: map[string]*Buffer{
		"P1": &Buffer{
			Count: 0,
			Rows:  []*BufferRow{},
		},
		"P2": &Buffer{
			Count: 0,
			Rows:  []*BufferRow{},
		},
		"P3": &Buffer{
			Count: 0,
			Rows:  []*BufferRow{},
		},
	},
}
