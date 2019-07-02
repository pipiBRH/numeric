package numeric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumericFormatString(t *testing.T) {
	result := NumericFormatString(123)
	assert.Equal(t, "123", result)

	result = NumericFormatString(-05)
	assert.Equal(t, "-5", result)

	result = NumericFormatString(05)
	assert.Equal(t, "5", result)

	result = NumericFormatString(05.123)
	assert.Equal(t, "5.123", result)

	result = NumericFormatString(123.0)
	assert.Equal(t, "123", result)

	result = NumericFormatString(-123)
	assert.Equal(t, "-123", result)

	result = NumericFormatString(1234)
	assert.Equal(t, "1,234", result)

	result = NumericFormatString(123456789)
	assert.Equal(t, "123,456,789", result)

	result = NumericFormatString(123.456789)
	assert.Equal(t, "123.456789", result)

	result = NumericFormatString(-123.456789)
	assert.Equal(t, "-123.456789", result)

	result = NumericFormatString(-123456.456789)
	assert.Equal(t, "-123,456.456789", result)

	result = NumericFormatString(0.456789)
	assert.Equal(t, "0.456789", result)

	result = NumericFormatString(-0.456789)
	assert.Equal(t, "-0.456789", result)

	result = NumericFormatString(0.000001)
	assert.Equal(t, "0.000001", result)
}
