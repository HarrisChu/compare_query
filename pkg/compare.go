package pkg

import (
	"fmt"
	"strconv"
	"strings"

	nebula "github.com/vesoft-inc/nebula-go/v3"
)

type comparer struct {
	l      *QueryFactory
	r      *QueryFactory
	client *nebula.Session
	pool   *nebula.ConnectionPool
}

func NewCompare(left, right string, nebulagraph string, space string) (*comparer, error) {
	nebulagraph = strings.TrimSpace(nebulagraph)
	if len(strings.Split(nebulagraph, ":")) != 2 {
		return nil, fmt.Errorf("invalid nebula graph address: %s", nebulagraph)
	}
	host, p := strings.Split(nebulagraph, ":")[0], strings.Split(nebulagraph, ":")[1]
	port, err := strconv.Atoi(p)
	if err != nil {
		return nil, err
	}
	conf := nebula.GetDefaultConf()
	pool, err := nebula.NewConnectionPool([]nebula.HostAddress{{host, port}}, conf, nebula.DefaultLogger{})
	if err != nil {
		return nil, err
	}
	session, err := pool.GetSession("root", "nebula")
	if err != nil {
		return nil, err
	}
	_, err = session.Execute(fmt.Sprintf("USE %s", space))
	if err != nil {
		return nil, err
	}

	l, err := NewQueryFactory(left)
	if err != nil {
		return nil, err
	}
	r, err := NewQueryFactory(right)
	if err != nil {
		return nil, err
	}

	return &comparer{
		l:      l,
		r:      r,
		pool:   pool,
		client: session,
	}, nil
}

func (c *comparer) compare(vid string, verbose bool) error {
	left, err := c.l.NewQuery(vid)
	if err != nil {
		return err
	}
	right, err := c.r.NewQuery(vid)
	if err != nil {
		return err
	}
	respLeft, err := c.client.Execute(left)
	if err != nil {
		return err
	}
	if !respLeft.IsSucceed() {
		return fmt.Errorf("left query: %s failed: %s", left, respLeft.GetErrorMsg())
	}
	respRight, err := c.client.Execute(right)
	if err != nil {
		return err
	}
	if !respRight.IsSucceed() {
		return fmt.Errorf("right query: %s failed: %s", right, respRight.GetErrorMsg())
	}
	if respLeft.GetRowSize() != respRight.GetRowSize() {
		return fmt.Errorf("left query: %s and right query: %s have different row size", left, right)
	}
	for i := 0; i < respLeft.GetRowSize(); i++ {
		leftRecord, err := respLeft.GetRowValuesByIndex(i)
		if err != nil {
			return err
		}
		rightRecord, err := respRight.GetRowValuesByIndex(i)
		if err != nil {
			return err
		}
		if leftRecord.String() != rightRecord.String() {
			return fmt.Errorf("left query: %s and right query: %s have different result", left, right)
		}
	}
	if verbose {
		fmt.Printf("left query: \n%s \nright query: \n%s \n\n", left, right)
	}

	return nil
}

func (c *comparer) release() {
	c.client.Release()
	c.pool.Close()
}
