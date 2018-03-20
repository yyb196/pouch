package plugins

import "io"

// ContainerPlugin defines in which place a plugin will be triggered in container lifecycle
type ContainerPlugin interface {
	// PreCreate defines plugin point where recevives an container create request, in this plugin point user
	// could change the container create body passed-in by http request body
	PreCreate(io.ReadCloser) (io.ReadCloser, error)

	// PreStart returns an array of PreStartHook which will pass to runc, in PreStartHook there is a Priority which
	// used to sort the pre start array that pass to runc, network plugin hook has priority value 0.
	PreStart(interface{}) ([]int, [][]string, error)
}
