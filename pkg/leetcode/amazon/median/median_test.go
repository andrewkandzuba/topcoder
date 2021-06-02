package median

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Median_Case1(t *testing.T) {
	mf := Constructor()
	mf.AddNum(1);    // arr = [1]
	mf.AddNum(2);    // arr = [1, 2]
	assert.EqualValues(t, 1.5, mf.FindMedian()); // return 1.5 (i.e., (1 + 2) / 2)
	mf.AddNum(3);    // arr[1, 2, 3]
	assert.EqualValues(t, 2.0, mf.FindMedian());
}
