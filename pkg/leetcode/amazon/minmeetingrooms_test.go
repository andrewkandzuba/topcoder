package amazon

import (
"github.com/stretchr/testify/assert"
"testing"
)

func Test_MinMeetingRooms_Case1(t *testing.T) {
	assert.EqualValues(t, 2, minMeetingRooms([][]int{{0,30},{5,10},{15,20}}))
	assert.EqualValues(t, 1, minMeetingRooms( [][]int{{7,10},{2,4}}))
}

func Test_MinMeetingRooms_Case2(t *testing.T) {
	assert.EqualValues(t, 2, minMeetingRooms([][]int{{0,8},{5,10},{9,20}}))
	assert.EqualValues(t, 1, minMeetingRooms([][]int{{0,8}}))
	assert.EqualValues(t, 1, minMeetingRooms([][]int{{0,8}, {10, 15}}))
}


func Test_MinMeetingRooms_Case3(t *testing.T) {
	assert.EqualValues(t, 1, minMeetingRooms([][]int{{13,15},{1, 13}}))
}