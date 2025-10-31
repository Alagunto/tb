package sendables

import "github.com/alagunto/tb/files"

type Prepared struct {
	SendMethod SendMethod

	// Params that must be provided to send<Type> method as-is
	Params map[string]string

	// Files that will be sent to the send<Type> method, if any
	Files map[string]files.FileSource
}
