package types

const DefaultTtl = 20

// DefaultParams returns default module parameters.
func DefaultParams() Params {
	return Params{
		Ttl: DefaultTtl,
	}
}

// Validate does the sanity check on the params.
func (p Params) Validate() error {
	// Sanity check goes here.
	if p.Ttl == 0 {
		return ErrInvalidTtl
	}
	return nil
}
