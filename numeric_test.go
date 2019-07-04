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

	result = NumericFormatString(22222222222222222222222222222222)
	assert.Equal(t, "22,222,222,222,222,224,000,000,000,000,000", result)

	result = NumericFormatString(22222222222222222222222.22222222222222222)
	assert.Equal(t, "22,222,222,222,222,223,000,000", result)

	result = NumericFormatString(2222.222222222222222222222222222)
	assert.Equal(t, "2,222.222222222222", result)
}
