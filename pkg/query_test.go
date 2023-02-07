package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuery(t *testing.T) {
	vid := "123"
	expect := "GO FROM 123 OVER KNOWS YIELD KNOWS._dst as one_hop_id, timestamp(KNOWS.creationDate) as one_hop_edge_strength | ORDER BY $-.one_hop_edge_strength DESC | LIMIT 20 | GO FROM $-.one_hop_id OVER KNOWS YIELD DISTINCT KNOWS._dst as two_hop_id, timestamp(KNOWS.creationDate) + $-.one_hop_edge_strength as two_hop_edge_strength | ORDER BY $-.two_hop_edge_strength DESC | LIMIT 100 | GO FROM $-.two_hop_id OVER KNOWS YIELD KNOWS._dst as three_hop_id, timestamp(KNOWS.creationDate) + $-.two_hop_edge_strength as three_hop_edge_strength | ORDER BY $-.three_hop_edge_strength DESC | LIMIT 2000 | YIELD DISTINCT $-.three_hop_id as three_hop_id;"
	qf, err := NewQueryFactory("go_query1.tpl")
	if err != nil {
		t.Fatal(err)
	}
	query, err := qf.NewQuery(vid)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, expect, query)
}
