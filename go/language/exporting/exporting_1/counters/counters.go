// Package counters provides alert counter support.
package counters

// 總結：大寫開頭可以被其他package使用，小寫不能

// AlertCounter is an exported named type that contains an integer counter for alerts.
// The first character is in upper-case format so it is considered to be exported.
type AlertCounter int

// alertCounter is an unexported named type that contains an integer counter for alerts.
// The first character is in lower-case format so it is considered to be unexported.
// It is not accessible for other packages, unless they are part of the package counters themselves.
type alertCounter int
