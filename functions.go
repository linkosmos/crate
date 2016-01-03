package crate

// UnpackFunc - will execute against given function when
// will perform container dump
type UnpackFunc func(string, int)

// FilterFunc - filters incoming data that enters crate
type FilterFunc func(string) bool

// PurifyFunc - purifies bucket from unwanted data
type PurifyFunc func(string, int) bool

// DumpFunc - passes raw data struct
type DumpFunc func(s map[string]int)
