package crate

import "github.com/linkosmos/meta"

// Crate -
type Crate struct {
	*meta.Map
	unpack UnpackFunc
	filter FilterFunc
	purify PurifyFunc
	dump   DumpFunc
}

// New - initialize new crate with given size
func New(size int) *Crate {
	return &Crate{
		Map: meta.NewMap(size),
	}
}

// AttachUnpack - attaches unpacking function
func (c *Crate) AttachUnpack(f UnpackFunc) {
	c.unpack = f
}

// AttachFilter - attaches filter for filtering items that comes to crate
func (c *Crate) AttachFilter(f FilterFunc) {
	c.filter = f
}

// AttachPurify - attaches purify function to filter unwanted data
func (c *Crate) AttachPurify(f PurifyFunc) {
	c.purify = f
}

// AttachDump - attaches dump to function to expose crate contents
func (c *Crate) AttachDump(f DumpFunc) {
	c.dump = f
}

// Add - adds items to crate if filter is attached will filter out incoming data
func (c *Crate) Add(s string) bool {
	if c.filter != nil && c.filter(s) {
		return false
	}
	return c.Map.Add(s)
}

// Purify - purify crate from unwanted values
// mutates crate meta map contents
func (c *Crate) Purify() {
	if c.purify != nil {
		output := make(map[string]int, c.Map.Count()/2)
		c.Map.Dump(func(key string, value int) {
			if c.purify(key, value) {
				output[key] = value
			}
		})
		c.Map.Data = output
	}
}

// Unload - dumps crate contents to a given unpack function
func (c *Crate) Unload() {
	if c.unpack != nil {
		c.Map.Dump(c.unpack)
	}
}

// Dump - dumps crate Map Data contents
func (c *Crate) Dump() {
	if c.dump != nil {
		c.dump(c.Map.Data)
	}
}
