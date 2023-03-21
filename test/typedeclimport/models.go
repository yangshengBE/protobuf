package typedeclimport

import subpkg "github.com/yangshengBE/protobuf/test/typedeclimport/subpkg"

type SomeMessage struct {
	Imported subpkg.AnotherMessage
}
