package expressgo

type Body struct {
	byt []byte
}

// Get body as a plain byte array.
func (bd *Body) GetPlain() []byte {
	return bd.byt
}

// Get the text version of the body.
func (bd *Body) GetText() string {
	return string(bd.byt)
}

// Get the json version of the body.
func (bd *Body) GetJSON() JSONData {
	return DecodeJSON(bd.byt)
}
