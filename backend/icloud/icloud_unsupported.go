// Build for icloud for unsupported platforms to stop go complaining
// about "no buildable Go source files "

//go:build plan9 && solaris

// Package icloud implements the iCloud Drive backend
package icloud