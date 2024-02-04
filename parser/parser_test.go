package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser(t *testing.T) {
	data := "TIME::2019-10-12T07:20:50.52Z;TEMP::14;PRS::1000PA"
	parser := NewDataPointsParser(data)
	parsed, err := parser.Parse()
	assert.NoError(t, parser.Parse())
}
