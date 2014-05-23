package jsmin

import (
	"bytes"
	"testing"

	assert "github.com/pilu/miniassert"
)

func TestMinifier_min(t *testing.T) {
	original := `// is.js

// (c) 2001 Douglas Crockford
// 2001 June 3


// is

// The -is- object is used to identify the browser.  Every browser edition
// identifies itself, but there is no standard way of doing it, and some of
// the identification is deceptive. This is because the authors of web
// browsers are liars. For example, Microsoft's IE browsers claim to be
// Mozilla 4. Netscape 6 claims to be version 5.

// Warning: Do not use this awful, awful code or any other thing like it.
// Seriously.

var is = {
    ie:      navigator.appName == 'Microsoft Internet Explorer',
    java:    navigator.javaEnabled(),
    ns:      navigator.appName == 'Netscape',
    ua:      navigator.userAgent.toLowerCase(),
    version: parseFloat(navigator.appVersion.substr(21)) ||
             parseFloat(navigator.appVersion),
    win:     navigator.platform == 'Win32'
}

is.mac = is.ua.indexOf('mac') >= 0;

if (is.ua.indexOf('opera') >= 0) {
    is.ie = is.ns = false;
    is.opera = true;
}

if (is.ua.indexOf('gecko') >= 0) {
    is.ie = is.ns = false;
    is.gecko = true;
}`

	expected := `
var is={ie:navigator.appName=='Microsoft Internet Explorer',java:navigator.javaEnabled(),ns:navigator.appName=='Netscape',ua:navigator.userAgent.toLowerCase(),version:parseFloat(navigator.appVersion.substr(21))||parseFloat(navigator.appVersion),win:navigator.platform=='Win32'}
is.mac=is.ua.indexOf('mac')>=0;if(is.ua.indexOf('opera')>=0){is.ie=is.ns=false;is.opera=true;}
if(is.ua.indexOf('gecko')>=0){is.ie=is.ns=false;is.gecko=true;}`

	output := bytes.NewBufferString("")
	m := newMinifier(bytes.NewBufferString(original), output)
	err := m.min()
	assert.Nil(t, err)
	assert.Equal(t, expected, string(output.Bytes()))

	m = newMinifier(bytes.NewBufferString(`var x = 0; /*`), output)
	err = m.min()
	assert.NotNil(t, err)
	assert.Equal(t, errorUnterminatedComment, err)

	m = newMinifier(bytes.NewBufferString(`var x = "unterminated`), output)
	err = m.min()
	assert.NotNil(t, err)
	assert.Equal(t, errorUnterminatedStringLiteral, err)

	m = newMinifier(bytes.NewBufferString(`var x = /[1-2/`), output)
	err = m.min()
	assert.NotNil(t, err)
	assert.Equal(t, errorUnterminatedSetInRegexpLiteral, err)

	m = newMinifier(bytes.NewBufferString(`var x = /hello`), output)
	err = m.min()
	assert.NotNil(t, err)
	assert.Equal(t, errorUnterminatedRegexpLiteral, err)
}
