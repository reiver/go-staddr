package staddr


func (address *Address) Normalize() error {

//@TODO: Add support for other countries.
	return normalizeCanada(address)
}
