package config

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Group struct {
	Name  string `yaml:"name"`
	Tests []Test `yaml:"tests"`
}

const (
	TestTypeText  = "text"
	TestTypeMedia = "media"
)

type Test struct {
	// here we have to work with "input" & "internal"
	// because the inputs and post-compilation values are different types

	// Input -> parsed straight from yaml
	// Internal -> compiled value from Input

	// inputs
	TypeInput     string `yaml:"type"`
	SizeInput     string `yaml:"size"`
	IntervalInput int    `yaml:"interval"`

	// parsed values
	TypeInternal     string `yaml:"typeInternal"`
	SizeInternal     int    `yaml:"sizeInternal"`     // in KB
	IntervalInternal int    `yaml:"intervalInternal"` // in Seconds

}

func (c *NodeGroup) parseGroups() error {

	// we check if the user doesn't try to add daemon to types that are not supposed to have group
	// only types that support daemon are:
	// Peer & Replication

	switch c.NodeType {
	case NodeTypeRDVP:
		if len(c.Groups) > 0 {
			return errors.New("can't have daemon in type RDVP")
		}
	case NodeTypeBootstrap:
		if len(c.Groups) > 0 {
			return errors.New("can't have daemon in type Bootstrap")
		}
	case NodeTypeRelay:
		if len(c.Groups) > 0 {
			return errors.New("can't have daemon in type Relay")
		}
	case NodeTypeReplication:
		// allowed
	case NodeTypePeer:
		// daemon a re allowed here
		// do nothing
	}

	// parse the test cases for each individual group
	for i, group := range c.Groups {
		config.Attributes.Groups[c.Groups[i].Name] = &c.Groups[i]

		// parse the tests
		for j, test := range c.Groups[i].Tests {
			t := strings.ToLower(test.TypeInput)

			// parse group type
			switch {
			case strings.Contains(t, TestTypeText):
				c.Groups[i].Tests[j].TypeInternal = TestTypeText
			case strings.Contains(t, TestTypeMedia):
				c.Groups[i].Tests[j].TypeInternal = TestTypeMedia
			default:
				return errors.New(fmt.Sprintf("invalid test type in test in group: %s test: %v", group.Name, i+1))
			}

			// parse message size
			s := strings.ToLower(test.SizeInput)
			var unit int

			switch {
			case strings.Contains(s, "kb"):
				unit = 1000
			case strings.Contains(s, "mb"):
				unit = 1000000
			case strings.Contains(s, "gb"):
				unit = 1000000000
			case strings.Contains(s, "b"):
				// this one has to go last because of the single "b"
				unit = 1
			default:
				return errors.New(fmt.Sprintf("invalid test size in test in group: %s test: %v", group.Name, i+1))
			}

			size, err := strconv.Atoi(s[:len(s)-2])
			if err != nil {
				return errors.New(fmt.Sprintf("invalid test size in test in group: %s test: %v", group.Name, i+1))
			}

			c.Groups[i].Tests[j].SizeInternal = size * unit

			if unit > 10000000 && c.Groups[i].Tests[j].TypeInternal == TestTypeText {
				log.Printf("caution big text size group: %s test: %v\n", group.Name, i+1)
			}

			// parse interval
			if test.IntervalInput > 0 {
				c.Groups[i].Tests[j].IntervalInternal = test.IntervalInput
			} else {
				return errors.New(fmt.Sprintf("invterval needs to be bigger than 0 in group: %s test: %v", group.Name, i+1))

			}
		}
	}

	return nil
}

func (g Group) Hash () [16]byte {
	bytes, err := json.Marshal(g)
	if err != nil {
		log.Println(err)
	}

	return md5.Sum(bytes)
}

func (c Config) GetAmountOfGroups() (i int) {
	return len(c.Attributes.Groups)
}
