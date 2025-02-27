//go:generate go-enum --marshal --forceupper --mustparse --nocomments --names --values
package cfg

// Specifies the type of config file that corresponds to this configuration type.
/*
ENUM(
	UNKNOWN	= 0x00000000	//Unknown configuration type; this can be used in lieu of an actual config.
	SYSENV	= 0x00000001	//System ENV configuration type; fills from system env or a `.env` file in the root working directory. This configuration type always has priority over other types.
	JSON	= 0x00000010	//JSON config; uses Go's in-built JSON marshalling API.
	ENV		= 0x00000100	//
	TOML	= 0x00001000	//
	YAML	= 0x00010000	//
)
*/
type Ctype int

// Masks this `Ctype` object with multiple inputs.
func (c Ctype) Mask(cts ...Ctype) Ctype {
	val := c
	for _, ct := range cts {
		val |= ct
	}
	return val
}

// Unmasks the value of this `Ctype` to get the file types that can be used to initialize a configuration object.
func (c Ctype) Unmask() []Ctype {
	hits := make([]Ctype, 0, len(CtypeValues()))
	for _, val := range CtypeValues() {
		if val&c != 0 {
			hits = append(hits, val)
		}
	}
	return hits
}
