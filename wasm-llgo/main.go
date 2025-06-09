package main

import (
	"unsafe"

	"github.com/goplus/gdext/gdext"
	"github.com/goplus/lib/c"
)

const (
	LLGoPackage = "link: -s SIDE_MODULE=1 -s IMPORTED_MEMORY=1 -s USE_PTHREADS=1"
)

var i = 100

// typedef struct {
// 	/* Minimum initialization level required.
// 	 * If Core or Servers, the extension needs editor or game restart to take effect */
// 	GDExtensionInitializationLevel minimum_initialization_level;
// 	/* Up to the user to supply when initializing */
// 	void *userdata;
// 	/* This function will be called multiple times for each initialization level. */
// 	void (*initialize)(void *userdata, GDExtensionInitializationLevel p_level);
// 	void (*deinitialize)(void *userdata, GDExtensionInitializationLevel p_level);
// } GDExtensionInitialization;

func initialize(userdata unsafe.Pointer, p_level gdext.InitializationLevel) {
	if p_level != gdext.INITIALIZATION_SCENE {
		return
	}
	c.Printf(c.Str("llgo initialize\n"))
}

func uninitialize(userdata unsafe.Pointer, p_level gdext.InitializationLevel) {
	c.Printf(c.Str("uninitialize\n"))
}

//llgo:type C
type Initialize func(userdata unsafe.Pointer, p_level gdext.InitializationLevel)

//llgo:type C
type Deinitialize func(userdata unsafe.Pointer, p_level gdext.InitializationLevel)

//llgo:type C
type Initialization struct {
	MinimumInitializationLevel gdext.InitializationLevel
	Userdata                   c.Pointer
	Initialize                 Initialize
	Deinitialize               Deinitialize
}

//export summator_library_init
func summator_library_init(p_get_proc_address gdext.InterfaceGetProcAddress,
	p_library gdext.ClassLibraryPtr, r_initialization *Initialization) c.Int {
	pkginit()
	c.Printf(c.Str("llgo library_init\n"))
	c.Printf(c.Str("global i=%d\n"), i)
	r_initialization.MinimumInitializationLevel = gdext.INITIALIZATION_SCENE
	r_initialization.Initialize = initialize
	r_initialization.Deinitialize = uninitialize
	return 1
}

func init() {
	c.Printf(c.Str("github.com/goplus/gdext.init\n"))
}

//go:linkname pkginit github.com/goplus/gdext.init
func pkginit()

func main() {
}
