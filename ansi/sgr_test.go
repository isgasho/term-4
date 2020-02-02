// SPDX-License-Identifier: MIT

package ansi

import (
	"testing"

	"github.com/issue9/assert"
)

func TestSGR(t *testing.T) {
	// 多个
	sgr := SGR(SGRBBlack, SGRBBlue, SGRFRed)
	result := "\033[" + SGRBBlack + ";" + SGRBBlue + ";" + SGRFRed + "m"
	assert.Equal(t, sgr, result)

	// 传递单个值
	sgr = SGR(SGRBBlack)
	result = "\033[" + SGRBBlack + "m"
	assert.Equal(t, sgr, result)

	// 传递空值，相当于SReset
	sgr = SGR()
	assert.Equal(t, sgr, Reset)
}
