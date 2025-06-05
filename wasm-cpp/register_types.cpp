#include "gdextension_interface.h"
#include <iostream>

// 定义导出宏
#ifdef __cplusplus
extern "C" {
#endif

#if defined(_WIN32)
#define GDE_EXPORT __declspec(dllexport)
#elif defined(__GNUC__) || defined(__clang__)
#define GDE_EXPORT __attribute__((visibility("default")))
#else
#define GDE_EXPORT
#endif

void initialize_summator_types(void *userdata, GDExtensionInitializationLevel p_level)
{
	if (p_level != GDEXTENSION_INITIALIZATION_SCENE) {
		return;
	}
	std::cout << "initialize_summator_types" << std::endl;
}

void uninitialize_summator_types(void *userdata, GDExtensionInitializationLevel p_level) {
	if (p_level != GDEXTENSION_INITIALIZATION_SCENE) {
		return;
	}
}

// Initialization.
GDExtensionBool GDE_EXPORT summator_library_init(GDExtensionInterfaceGetProcAddress p_get_proc_address, 
	GDExtensionClassLibraryPtr p_library, GDExtensionInitialization *r_initialization) {
	std::cout << "summator_library_init" << std::endl;
	r_initialization->initialize = initialize_summator_types;
	r_initialization->deinitialize = uninitialize_summator_types;
	r_initialization->minimum_initialization_level = GDEXTENSION_INITIALIZATION_SCENE;
	return true;
}

#ifdef __cplusplus
}
#endif
