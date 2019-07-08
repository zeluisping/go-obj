package main

// MarshalOptions gives finer control over the final result
type MarshalOptions struct {
	OmitDefaultOptional bool
	FloatPrecision      int
}

// DefaultMarshalOptions _
var DefaultMarshalOptions = MarshalOptions{
	OmitDefaultOptional: true,
	FloatPrecision:      6,
}
