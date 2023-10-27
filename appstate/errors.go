// Copyright © 2023 whatsmeow
// Redeveloped by Amirul Dev

package appstate

import "errors"

// Errors that this package can return.
var (
	ErrMissingPreviousSetValueOperation = errors.New("missing value MAC of previous SET operation")
	ErrMismatchingLTHash                = errors.New("mismatching LTHash")
	ErrMismatchingPatchMAC              = errors.New("mismatching patch MAC")
	ErrMismatchingContentMAC            = errors.New("mismatching content MAC")
	ErrMismatchingIndexMAC              = errors.New("mismatching index MAC")
	ErrKeyNotFound                      = errors.New("didn't find app state key")
)
