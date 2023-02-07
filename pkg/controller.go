package pkg

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Controller struct {
	sourceFile  string
	sampleFile  string
	sampleCount int
	sampleRange int
	nebulagraph string
	space       string
}

func NewController(sourceFile, sampleFile string, sampleCount, sampleRange int, nebulagraph string, nebulaSpace string) *Controller {
	c := &Controller{
		sourceFile:  sourceFile,
		sampleFile:  sampleFile,
		sampleCount: sampleCount,
		sampleRange: sampleRange,
		nebulagraph: nebulagraph,
		space:       nebulaSpace,
	}
	return c
}

func (c *Controller) Sample() error {
	fl, err := os.Open(c.sourceFile)
	if err != nil {
		return err
	}
	defer fl.Close()
	writeFile, err := os.OpenFile(c.sampleFile, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer writeFile.Close()
	scanner := bufio.NewScanner(fl)
	var (
		i     int
		count int
	)
	for scanner.Scan() {
		if i%c.sampleRange == 0 {
			count++
			// add to sample file
			writeFile.Write(scanner.Bytes())
			writeFile.Write([]byte("\n"))
			if count > c.sampleCount {
				break
			}
		}
	}
	return nil
}

func (c *Controller) Compare() error {
	fl, err := os.Open(c.sampleFile)
	if err != nil {
		return err
	}
	defer fl.Close()
	scanner := bufio.NewScanner(fl)
	compare1, err := NewCompare("go_query1.tpl", "match_query1.tpl", c.nebulagraph, c.space)
	if err != nil {
		return err
	}
	compare2, err := NewCompare("go_query1.tpl", "match_query1.tpl", c.nebulagraph, c.space)
	if err != nil {
		return err
	}
	var i int
	for scanner.Scan() {
		if i%10 == 0 {
			fmt.Printf("Now(): %v, run %d iterations. \n", time.Now(), i)
		}
		if i == 0 {
			if err := compare1.compare(scanner.Text(), true); err != nil {
				fmt.Println(err.Error())
			}
			if err := compare2.compare(scanner.Text(), true); err != nil {
				fmt.Println(err.Error())
			}
		} else {
			if err := compare1.compare(scanner.Text(), false); err != nil {
				fmt.Println(err.Error())
			}
			if err := compare2.compare(scanner.Text(), false); err != nil {
				fmt.Println(err.Error())
			}
		}
		i++
	}
	return nil
}
