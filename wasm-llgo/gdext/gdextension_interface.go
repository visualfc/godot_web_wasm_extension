package gdext

import (
	_ "unsafe"

	"github.com/goplus/lib/c"
)

type Char32T c.Uint32T
type Char16T c.Uint16T
type VariantType c.Int

const (
	VARIANT_TYPE_NIL                  VariantType = 0
	VARIANT_TYPE_BOOL                 VariantType = 1
	VARIANT_TYPE_INT                  VariantType = 2
	VARIANT_TYPE_FLOAT                VariantType = 3
	VARIANT_TYPE_STRING               VariantType = 4
	VARIANT_TYPE_VECTOR2              VariantType = 5
	VARIANT_TYPE_VECTOR2I             VariantType = 6
	VARIANT_TYPE_RECT2                VariantType = 7
	VARIANT_TYPE_RECT2I               VariantType = 8
	VARIANT_TYPE_VECTOR3              VariantType = 9
	VARIANT_TYPE_VECTOR3I             VariantType = 10
	VARIANT_TYPE_TRANSFORM2D          VariantType = 11
	VARIANT_TYPE_VECTOR4              VariantType = 12
	VARIANT_TYPE_VECTOR4I             VariantType = 13
	VARIANT_TYPE_PLANE                VariantType = 14
	VARIANT_TYPE_QUATERNION           VariantType = 15
	VARIANT_TYPE_AABB                 VariantType = 16
	VARIANT_TYPE_BASIS                VariantType = 17
	VARIANT_TYPE_TRANSFORM3D          VariantType = 18
	VARIANT_TYPE_PROJECTION           VariantType = 19
	VARIANT_TYPE_COLOR                VariantType = 20
	VARIANT_TYPE_STRING_NAME          VariantType = 21
	VARIANT_TYPE_NODE_PATH            VariantType = 22
	VARIANT_TYPE_RID                  VariantType = 23
	VARIANT_TYPE_OBJECT               VariantType = 24
	VARIANT_TYPE_CALLABLE             VariantType = 25
	VARIANT_TYPE_SIGNAL               VariantType = 26
	VARIANT_TYPE_DICTIONARY           VariantType = 27
	VARIANT_TYPE_ARRAY                VariantType = 28
	VARIANT_TYPE_PACKED_BYTE_ARRAY    VariantType = 29
	VARIANT_TYPE_PACKED_INT32_ARRAY   VariantType = 30
	VARIANT_TYPE_PACKED_INT64_ARRAY   VariantType = 31
	VARIANT_TYPE_PACKED_FLOAT32_ARRAY VariantType = 32
	VARIANT_TYPE_PACKED_FLOAT64_ARRAY VariantType = 33
	VARIANT_TYPE_PACKED_STRING_ARRAY  VariantType = 34
	VARIANT_TYPE_PACKED_VECTOR2_ARRAY VariantType = 35
	VARIANT_TYPE_PACKED_VECTOR3_ARRAY VariantType = 36
	VARIANT_TYPE_PACKED_COLOR_ARRAY   VariantType = 37
	VARIANT_TYPE_VARIANT_MAX          VariantType = 38
)

type VariantOperator c.Int

const (
	VARIANT_OP_EQUAL         VariantOperator = 0
	VARIANT_OP_NOT_EQUAL     VariantOperator = 1
	VARIANT_OP_LESS          VariantOperator = 2
	VARIANT_OP_LESS_EQUAL    VariantOperator = 3
	VARIANT_OP_GREATER       VariantOperator = 4
	VARIANT_OP_GREATER_EQUAL VariantOperator = 5
	VARIANT_OP_ADD           VariantOperator = 6
	VARIANT_OP_SUBTRACT      VariantOperator = 7
	VARIANT_OP_MULTIPLY      VariantOperator = 8
	VARIANT_OP_DIVIDE        VariantOperator = 9
	VARIANT_OP_NEGATE        VariantOperator = 10
	VARIANT_OP_POSITIVE      VariantOperator = 11
	VARIANT_OP_MODULE        VariantOperator = 12
	VARIANT_OP_POWER         VariantOperator = 13
	VARIANT_OP_SHIFT_LEFT    VariantOperator = 14
	VARIANT_OP_SHIFT_RIGHT   VariantOperator = 15
	VARIANT_OP_BIT_AND       VariantOperator = 16
	VARIANT_OP_BIT_OR        VariantOperator = 17
	VARIANT_OP_BIT_XOR       VariantOperator = 18
	VARIANT_OP_BIT_NEGATE    VariantOperator = 19
	VARIANT_OP_AND           VariantOperator = 20
	VARIANT_OP_OR            VariantOperator = 21
	VARIANT_OP_XOR           VariantOperator = 22
	VARIANT_OP_NOT           VariantOperator = 23
	VARIANT_OP_IN            VariantOperator = 24
	VARIANT_OP_MAX           VariantOperator = 25
)

type VariantPtr c.Pointer
type ConstVariantPtr c.Pointer
type UninitializedVariantPtr c.Pointer
type StringNamePtr c.Pointer
type ConstStringNamePtr c.Pointer
type UninitializedStringNamePtr c.Pointer
type StringPtr c.Pointer
type ConstStringPtr c.Pointer
type UninitializedStringPtr c.Pointer
type ObjectPtr c.Pointer
type ConstObjectPtr c.Pointer
type UninitializedObjectPtr c.Pointer
type TypePtr c.Pointer
type ConstTypePtr c.Pointer
type UninitializedTypePtr c.Pointer
type MethodBindPtr c.Pointer
type Int c.Int64T
type Bool c.Uint8T
type GDObjectInstanceID c.Uint64T
type RefPtr c.Pointer
type ConstRefPtr c.Pointer
type CallErrorType c.Int

const (
	CALL_OK                       CallErrorType = 0
	CALL_ERROR_INVALID_METHOD     CallErrorType = 1
	CALL_ERROR_INVALID_ARGUMENT   CallErrorType = 2
	CALL_ERROR_TOO_MANY_ARGUMENTS CallErrorType = 3
	CALL_ERROR_TOO_FEW_ARGUMENTS  CallErrorType = 4
	CALL_ERROR_INSTANCE_IS_NULL   CallErrorType = 5
	CALL_ERROR_METHOD_NOT_CONST   CallErrorType = 6
)

type CallError struct {
	Error    CallErrorType
	Argument c.Int32T
	Expected c.Int32T
}

// llgo:type C
type VariantFromTypeConstructorFunc func(UninitializedVariantPtr, TypePtr)

// llgo:type C
type TypeFromVariantConstructorFunc func(UninitializedTypePtr, VariantPtr)

// llgo:type C
type PtrOperatorEvaluator func(ConstTypePtr, ConstTypePtr, TypePtr)

// llgo:type C
type PtrBuiltInMethod func(TypePtr, *ConstTypePtr, TypePtr, c.Int)

// llgo:type C
type PtrConstructor func(UninitializedTypePtr, *ConstTypePtr)

// llgo:type C
type PtrDestructor func(TypePtr)

// llgo:type C
type PtrSetter func(TypePtr, ConstTypePtr)

// llgo:type C
type PtrGetter func(ConstTypePtr, TypePtr)

// llgo:type C
type PtrIndexedSetter func(TypePtr, Int, ConstTypePtr)

// llgo:type C
type PtrIndexedGetter func(ConstTypePtr, Int, TypePtr)

// llgo:type C
type PtrKeyedSetter func(TypePtr, ConstTypePtr, ConstTypePtr)

// llgo:type C
type PtrKeyedGetter func(ConstTypePtr, ConstTypePtr, TypePtr)

// llgo:type C
type PtrKeyedChecker func(ConstVariantPtr, ConstVariantPtr) c.Uint32T

// llgo:type C
type PtrUtilityFunction func(TypePtr, *ConstTypePtr, c.Int)

// llgo:type C
type ClassConstructor func(__llgo_va_list ...interface{}) ObjectPtr

// llgo:type C
type InstanceBindingCreateCallback func(c.Pointer, c.Pointer) c.Pointer

// llgo:type C
type InstanceBindingFreeCallback func(c.Pointer, c.Pointer, c.Pointer)

// llgo:type C
type InstanceBindingReferenceCallback func(c.Pointer, c.Pointer, Bool) Bool

type InstanceBindingCallbacks struct {
	CreateCallback    InstanceBindingCreateCallback
	FreeCallback      InstanceBindingFreeCallback
	ReferenceCallback InstanceBindingReferenceCallback
}
type ClassInstancePtr c.Pointer

// llgo:type C
type ClassSet func(ClassInstancePtr, ConstStringNamePtr, ConstVariantPtr) Bool

// llgo:type C
type ClassGet func(ClassInstancePtr, ConstStringNamePtr, VariantPtr) Bool

// llgo:type C
type ClassGetRID func(ClassInstancePtr) c.Uint64T

type PropertyInfo struct {
	Type       VariantType
	Name       StringNamePtr
	ClassName  StringNamePtr
	Hint       c.Uint32T
	HintString StringPtr
	Usage      c.Uint32T
}

type MethodInfo struct {
	Name                 StringNamePtr
	ReturnValue          PropertyInfo
	Flags                c.Uint32T
	Id                   c.Int32T
	ArgumentCount        c.Uint32T
	Arguments            *PropertyInfo
	DefaultArgumentCount c.Uint32T
	DefaultArguments     *VariantPtr
}

// llgo:type C
type ClassGetPropertyList func(ClassInstancePtr, *c.Uint32T) *PropertyInfo

// llgo:type C
type ClassFreePropertyList func(ClassInstancePtr, *PropertyInfo)

// llgo:type C
type ClassPropertyCanRevert func(ClassInstancePtr, ConstStringNamePtr) Bool

// llgo:type C
type ClassPropertyGetRevert func(ClassInstancePtr, ConstStringNamePtr, VariantPtr) Bool

// llgo:type C
type ClassValidateProperty func(ClassInstancePtr, *PropertyInfo) Bool

// llgo:type C
type ClassNotification func(ClassInstancePtr, c.Int32T)

// llgo:type C
type ClassNotification2 func(ClassInstancePtr, c.Int32T, Bool)

// llgo:type C
type ClassToString func(ClassInstancePtr, *Bool, StringPtr)

// llgo:type C
type ClassReference func(ClassInstancePtr)

// llgo:type C
type ClassUnreference func(ClassInstancePtr)

// llgo:type C
type ClassCallVirtual func(ClassInstancePtr, *ConstTypePtr, TypePtr)

// llgo:type C
type ClassCreateInstance func(c.Pointer) ObjectPtr

// llgo:type C
type ClassFreeInstance func(c.Pointer, ClassInstancePtr)

// llgo:type C
type ClassRecreateInstance func(c.Pointer, ObjectPtr) ClassInstancePtr

// llgo:type C
type ClassGetVirtual func(c.Pointer, ConstStringNamePtr) ClassCallVirtual

// llgo:type C
type ClassGetVirtualCallData func(c.Pointer, ConstStringNamePtr) c.Pointer

// llgo:type C
type ClassCallVirtualWithData func(ClassInstancePtr, ConstStringNamePtr, c.Pointer, *ConstTypePtr, TypePtr)

type ClassCreationInfo struct {
	IsVirtual             Bool
	IsAbstract            Bool
	SetFunc               ClassSet
	GetFunc               ClassGet
	GetPropertyListFunc   ClassGetPropertyList
	FreePropertyListFunc  ClassFreePropertyList
	PropertyCanRevertFunc ClassPropertyCanRevert
	PropertyGetRevertFunc ClassPropertyGetRevert
	NotificationFunc      ClassNotification
	ToStringFunc          ClassToString
	ReferenceFunc         ClassReference
	UnreferenceFunc       ClassUnreference
	CreateInstanceFunc    ClassCreateInstance
	FreeInstanceFunc      ClassFreeInstance
	GetVirtualFunc        ClassGetVirtual
	GetRidFunc            ClassGetRID
	ClassUserdata         c.Pointer
}

type ClassCreationInfo2 struct {
	IsVirtual               Bool
	IsAbstract              Bool
	IsExposed               Bool
	SetFunc                 ClassSet
	GetFunc                 ClassGet
	GetPropertyListFunc     ClassGetPropertyList
	FreePropertyListFunc    ClassFreePropertyList
	PropertyCanRevertFunc   ClassPropertyCanRevert
	PropertyGetRevertFunc   ClassPropertyGetRevert
	ValidatePropertyFunc    ClassValidateProperty
	NotificationFunc        ClassNotification2
	ToStringFunc            ClassToString
	ReferenceFunc           ClassReference
	UnreferenceFunc         ClassUnreference
	CreateInstanceFunc      ClassCreateInstance
	FreeInstanceFunc        ClassFreeInstance
	RecreateInstanceFunc    ClassRecreateInstance
	GetVirtualFunc          ClassGetVirtual
	GetVirtualCallDataFunc  ClassGetVirtualCallData
	CallVirtualWithDataFunc ClassCallVirtualWithData
	GetRidFunc              ClassGetRID
	ClassUserdata           c.Pointer
}
type ClassLibraryPtr c.Pointer
type ClassMethodFlags c.Int

const (
	METHOD_FLAG_NORMAL   ClassMethodFlags = 1
	METHOD_FLAG_EDITOR   ClassMethodFlags = 2
	METHOD_FLAG_CONST    ClassMethodFlags = 4
	METHOD_FLAG_VIRTUAL  ClassMethodFlags = 8
	METHOD_FLAG_VARARG   ClassMethodFlags = 16
	METHOD_FLAG_STATIC   ClassMethodFlags = 32
	METHOD_FLAGS_DEFAULT ClassMethodFlags = 1
)

type ClassMethodArgumentMetadata c.Int

const (
	METHOD_ARGUMENT_METADATA_NONE           ClassMethodArgumentMetadata = 0
	METHOD_ARGUMENT_METADATA_INT_IS_INT8    ClassMethodArgumentMetadata = 1
	METHOD_ARGUMENT_METADATA_INT_IS_INT16   ClassMethodArgumentMetadata = 2
	METHOD_ARGUMENT_METADATA_INT_IS_INT32   ClassMethodArgumentMetadata = 3
	METHOD_ARGUMENT_METADATA_INT_IS_INT64   ClassMethodArgumentMetadata = 4
	METHOD_ARGUMENT_METADATA_INT_IS_UINT8   ClassMethodArgumentMetadata = 5
	METHOD_ARGUMENT_METADATA_INT_IS_UINT16  ClassMethodArgumentMetadata = 6
	METHOD_ARGUMENT_METADATA_INT_IS_UINT32  ClassMethodArgumentMetadata = 7
	METHOD_ARGUMENT_METADATA_INT_IS_UINT64  ClassMethodArgumentMetadata = 8
	METHOD_ARGUMENT_METADATA_REAL_IS_FLOAT  ClassMethodArgumentMetadata = 9
	METHOD_ARGUMENT_METADATA_REAL_IS_DOUBLE ClassMethodArgumentMetadata = 10
)

// llgo:type C
type ClassMethodCall func(c.Pointer, ClassInstancePtr, *ConstVariantPtr, Int, VariantPtr, *CallError)

// llgo:type C
type ClassMethodValidatedCall func(c.Pointer, ClassInstancePtr, *ConstVariantPtr, VariantPtr)

// llgo:type C
type ClassMethodPtrCall func(c.Pointer, ClassInstancePtr, *ConstTypePtr, TypePtr)

type ClassMethodInfo struct {
	Name                 StringNamePtr
	MethodUserdata       c.Pointer
	CallFunc             ClassMethodCall
	PtrcallFunc          ClassMethodPtrCall
	MethodFlags          c.Uint32T
	HasReturnValue       Bool
	ReturnValueInfo      *PropertyInfo
	ReturnValueMetadata  ClassMethodArgumentMetadata
	ArgumentCount        c.Uint32T
	ArgumentsInfo        *PropertyInfo
	ArgumentsMetadata    *ClassMethodArgumentMetadata
	DefaultArgumentCount c.Uint32T
	DefaultArguments     *VariantPtr
}

// llgo:type C
type CallableCustomCall func(c.Pointer, *ConstVariantPtr, Int, VariantPtr, *CallError)

// llgo:type C
type CallableCustomIsValid func(c.Pointer) Bool

// llgo:type C
type CallableCustomFree func(c.Pointer)

// llgo:type C
type CallableCustomHash func(c.Pointer) c.Uint32T

// llgo:type C
type CallableCustomEqual func(c.Pointer, c.Pointer) Bool

// llgo:type C
type CallableCustomLessThan func(c.Pointer, c.Pointer) Bool

// llgo:type C
type CallableCustomToString func(c.Pointer, *Bool, StringPtr)

type CallableCustomInfo struct {
	CallableUserdata c.Pointer
	Token            c.Pointer
	ObjectId         GDObjectInstanceID
	CallFunc         CallableCustomCall
	IsValidFunc      CallableCustomIsValid
	FreeFunc         CallableCustomFree
	HashFunc         CallableCustomHash
	EqualFunc        CallableCustomEqual
	LessThanFunc     CallableCustomLessThan
	ToStringFunc     CallableCustomToString
}
type ScriptInstanceDataPtr c.Pointer

// llgo:type C
type ScriptInstanceSet func(ScriptInstanceDataPtr, ConstStringNamePtr, ConstVariantPtr) Bool

// llgo:type C
type ScriptInstanceGet func(ScriptInstanceDataPtr, ConstStringNamePtr, VariantPtr) Bool

// llgo:type C
type ScriptInstanceGetPropertyList func(ScriptInstanceDataPtr, *c.Uint32T) *PropertyInfo

// llgo:type C
type ScriptInstanceFreePropertyList func(ScriptInstanceDataPtr, *PropertyInfo)

// llgo:type C
type ScriptInstanceGetClassCategory func(ScriptInstanceDataPtr, *PropertyInfo) Bool

// llgo:type C
type ScriptInstanceGetPropertyType func(ScriptInstanceDataPtr, ConstStringNamePtr, *Bool) VariantType

// llgo:type C
type ScriptInstanceValidateProperty func(ScriptInstanceDataPtr, *PropertyInfo) Bool

// llgo:type C
type ScriptInstancePropertyCanRevert func(ScriptInstanceDataPtr, ConstStringNamePtr) Bool

// llgo:type C
type ScriptInstancePropertyGetRevert func(ScriptInstanceDataPtr, ConstStringNamePtr, VariantPtr) Bool

// llgo:type C
type ScriptInstanceGetOwner func(ScriptInstanceDataPtr) ObjectPtr

// llgo:type C
type ScriptInstancePropertyStateAdd func(ConstStringNamePtr, ConstVariantPtr, c.Pointer)

// llgo:type C
type ScriptInstanceGetPropertyState func(ScriptInstanceDataPtr, ScriptInstancePropertyStateAdd, c.Pointer)

// llgo:type C
type ScriptInstanceGetMethodList func(ScriptInstanceDataPtr, *c.Uint32T) *MethodInfo

// llgo:type C
type ScriptInstanceFreeMethodList func(ScriptInstanceDataPtr, *MethodInfo)

// llgo:type C
type ScriptInstanceHasMethod func(ScriptInstanceDataPtr, ConstStringNamePtr) Bool

// llgo:type C
type ScriptInstanceCall func(ScriptInstanceDataPtr, ConstStringNamePtr, *ConstVariantPtr, Int, VariantPtr, *CallError)

// llgo:type C
type ScriptInstanceNotification func(ScriptInstanceDataPtr, c.Int32T)

// llgo:type C
type ScriptInstanceNotification2 func(ScriptInstanceDataPtr, c.Int32T, Bool)

// llgo:type C
type ScriptInstanceToString func(ScriptInstanceDataPtr, *Bool, StringPtr)

// llgo:type C
type ScriptInstanceRefCountIncremented func(ScriptInstanceDataPtr)

// llgo:type C
type ScriptInstanceRefCountDecremented func(ScriptInstanceDataPtr) Bool

// llgo:type C
type ScriptInstanceGetScript func(ScriptInstanceDataPtr) ObjectPtr

// llgo:type C
type ScriptInstanceIsPlaceholder func(ScriptInstanceDataPtr) Bool
type ScriptLanguagePtr c.Pointer

// llgo:type C
type ScriptInstanceGetLanguage func(ScriptInstanceDataPtr) ScriptLanguagePtr

// llgo:type C
type ScriptInstanceFree func(ScriptInstanceDataPtr)
type ScriptInstancePtr c.Pointer

type ScriptInstanceInfo struct {
	SetFunc                 ScriptInstanceSet
	GetFunc                 ScriptInstanceGet
	GetPropertyListFunc     ScriptInstanceGetPropertyList
	FreePropertyListFunc    ScriptInstanceFreePropertyList
	PropertyCanRevertFunc   ScriptInstancePropertyCanRevert
	PropertyGetRevertFunc   ScriptInstancePropertyGetRevert
	GetOwnerFunc            ScriptInstanceGetOwner
	GetPropertyStateFunc    ScriptInstanceGetPropertyState
	GetMethodListFunc       ScriptInstanceGetMethodList
	FreeMethodListFunc      ScriptInstanceFreeMethodList
	GetPropertyTypeFunc     ScriptInstanceGetPropertyType
	HasMethodFunc           ScriptInstanceHasMethod
	CallFunc                ScriptInstanceCall
	NotificationFunc        ScriptInstanceNotification
	ToStringFunc            ScriptInstanceToString
	RefcountIncrementedFunc ScriptInstanceRefCountIncremented
	RefcountDecrementedFunc ScriptInstanceRefCountDecremented
	GetScriptFunc           ScriptInstanceGetScript
	IsPlaceholderFunc       ScriptInstanceIsPlaceholder
	SetFallbackFunc         ScriptInstanceSet
	GetFallbackFunc         ScriptInstanceGet
	GetLanguageFunc         ScriptInstanceGetLanguage
	FreeFunc                ScriptInstanceFree
}

type ScriptInstanceInfo2 struct {
	SetFunc                 ScriptInstanceSet
	GetFunc                 ScriptInstanceGet
	GetPropertyListFunc     ScriptInstanceGetPropertyList
	FreePropertyListFunc    ScriptInstanceFreePropertyList
	GetClassCategoryFunc    ScriptInstanceGetClassCategory
	PropertyCanRevertFunc   ScriptInstancePropertyCanRevert
	PropertyGetRevertFunc   ScriptInstancePropertyGetRevert
	GetOwnerFunc            ScriptInstanceGetOwner
	GetPropertyStateFunc    ScriptInstanceGetPropertyState
	GetMethodListFunc       ScriptInstanceGetMethodList
	FreeMethodListFunc      ScriptInstanceFreeMethodList
	GetPropertyTypeFunc     ScriptInstanceGetPropertyType
	ValidatePropertyFunc    ScriptInstanceValidateProperty
	HasMethodFunc           ScriptInstanceHasMethod
	CallFunc                ScriptInstanceCall
	NotificationFunc        ScriptInstanceNotification2
	ToStringFunc            ScriptInstanceToString
	RefcountIncrementedFunc ScriptInstanceRefCountIncremented
	RefcountDecrementedFunc ScriptInstanceRefCountDecremented
	GetScriptFunc           ScriptInstanceGetScript
	IsPlaceholderFunc       ScriptInstanceIsPlaceholder
	SetFallbackFunc         ScriptInstanceSet
	GetFallbackFunc         ScriptInstanceGet
	GetLanguageFunc         ScriptInstanceGetLanguage
	FreeFunc                ScriptInstanceFree
}
type InitializationLevel c.Int

const (
	INITIALIZATION_CORE      InitializationLevel = 0
	INITIALIZATION_SERVERS   InitializationLevel = 1
	INITIALIZATION_SCENE     InitializationLevel = 2
	INITIALIZATION_EDITOR    InitializationLevel = 3
	MAX_INITIALIZATION_LEVEL InitializationLevel = 4
)

type Initialization struct {
	MinimumInitializationLevel InitializationLevel
	Userdata                   c.Pointer
	Initialize                 c.Pointer
	Deinitialize               c.Pointer
}

// llgo:type C
type InterfaceFunctionPtr func(__llgo_va_list ...interface{})

// llgo:type C
type InterfaceGetProcAddress func(*c.Char) InterfaceFunctionPtr

// llgo:type C
type InitializationFunction func(InterfaceGetProcAddress, ClassLibraryPtr, *Initialization) Bool

/* INTERFACE */

type GodotVersion struct {
	Major  c.Uint32T
	Minor  c.Uint32T
	Patch  c.Uint32T
	String *c.Char
}

// llgo:type C
type InterfaceGetGodotVersion func(*GodotVersion)

// llgo:type C
type InterfaceMemAlloc func(c.SizeT) c.Pointer

// llgo:type C
type InterfaceMemRealloc func(c.Pointer, c.SizeT) c.Pointer

// llgo:type C
type InterfaceMemFree func(c.Pointer)

// llgo:type C
type InterfacePrintError func(*c.Char, *c.Char, *c.Char, c.Int32T, Bool)

// llgo:type C
type InterfacePrintErrorWithMessage func(*c.Char, *c.Char, *c.Char, *c.Char, c.Int32T, Bool)

// llgo:type C
type InterfacePrintWarning func(*c.Char, *c.Char, *c.Char, c.Int32T, Bool)

// llgo:type C
type InterfacePrintWarningWithMessage func(*c.Char, *c.Char, *c.Char, *c.Char, c.Int32T, Bool)

// llgo:type C
type InterfacePrintScriptError func(*c.Char, *c.Char, *c.Char, c.Int32T, Bool)

// llgo:type C
type InterfacePrintScriptErrorWithMessage func(*c.Char, *c.Char, *c.Char, *c.Char, c.Int32T, Bool)

// llgo:type C
type InterfaceGetNativeStructSize func(ConstStringNamePtr) c.Uint64T

// llgo:type C
type InterfaceVariantNewCopy func(UninitializedVariantPtr, ConstVariantPtr)

// llgo:type C
type InterfaceVariantNewNil func(UninitializedVariantPtr)

// llgo:type C
type InterfaceVariantDestroy func(VariantPtr)

// llgo:type C
type InterfaceVariantCall func(VariantPtr, ConstStringNamePtr, *ConstVariantPtr, Int, UninitializedVariantPtr, *CallError)

// llgo:type C
type InterfaceVariantCallStatic func(VariantType, ConstStringNamePtr, *ConstVariantPtr, Int, UninitializedVariantPtr, *CallError)

// llgo:type C
type InterfaceVariantEvaluate func(VariantOperator, ConstVariantPtr, ConstVariantPtr, UninitializedVariantPtr, *Bool)

// llgo:type C
type InterfaceVariantSet func(VariantPtr, ConstVariantPtr, ConstVariantPtr, *Bool)

// llgo:type C
type InterfaceVariantSetNamed func(VariantPtr, ConstStringNamePtr, ConstVariantPtr, *Bool)

// llgo:type C
type InterfaceVariantSetKeyed func(VariantPtr, ConstVariantPtr, ConstVariantPtr, *Bool)

// llgo:type C
type InterfaceVariantSetIndexed func(VariantPtr, Int, ConstVariantPtr, *Bool, *Bool)

// llgo:type C
type InterfaceVariantGet func(ConstVariantPtr, ConstVariantPtr, UninitializedVariantPtr, *Bool)

// llgo:type C
type InterfaceVariantGetNamed func(ConstVariantPtr, ConstStringNamePtr, UninitializedVariantPtr, *Bool)

// llgo:type C
type InterfaceVariantGetKeyed func(ConstVariantPtr, ConstVariantPtr, UninitializedVariantPtr, *Bool)

// llgo:type C
type InterfaceVariantGetIndexed func(ConstVariantPtr, Int, UninitializedVariantPtr, *Bool, *Bool)

// llgo:type C
type InterfaceVariantIterInit func(ConstVariantPtr, UninitializedVariantPtr, *Bool) Bool

// llgo:type C
type InterfaceVariantIterNext func(ConstVariantPtr, VariantPtr, *Bool) Bool

// llgo:type C
type InterfaceVariantIterGet func(ConstVariantPtr, VariantPtr, UninitializedVariantPtr, *Bool)

// llgo:type C
type InterfaceVariantHash func(ConstVariantPtr) Int

// llgo:type C
type InterfaceVariantRecursiveHash func(ConstVariantPtr, Int) Int

// llgo:type C
type InterfaceVariantHashCompare func(ConstVariantPtr, ConstVariantPtr) Bool

// llgo:type C
type InterfaceVariantBooleanize func(ConstVariantPtr) Bool

// llgo:type C
type InterfaceVariantDuplicate func(ConstVariantPtr, VariantPtr, Bool)

// llgo:type C
type InterfaceVariantStringify func(ConstVariantPtr, StringPtr)

// llgo:type C
type InterfaceVariantGetType func(ConstVariantPtr) VariantType

// llgo:type C
type InterfaceVariantHasMethod func(ConstVariantPtr, ConstStringNamePtr) Bool

// llgo:type C
type InterfaceVariantHasMember func(VariantType, ConstStringNamePtr) Bool

// llgo:type C
type InterfaceVariantHasKey func(ConstVariantPtr, ConstVariantPtr, *Bool) Bool

// llgo:type C
type InterfaceVariantGetTypeName func(VariantType, UninitializedStringPtr)

// llgo:type C
type InterfaceVariantCanConvert func(VariantType, VariantType) Bool

// llgo:type C
type InterfaceVariantCanConvertStrict func(VariantType, VariantType) Bool

// llgo:type C
type InterfaceGetVariantFromTypeConstructor func(VariantType) VariantFromTypeConstructorFunc

// llgo:type C
type InterfaceGetVariantToTypeConstructor func(VariantType) TypeFromVariantConstructorFunc

// llgo:type C
type InterfaceVariantGetPtrOperatorEvaluator func(VariantOperator, VariantType, VariantType) PtrOperatorEvaluator

// llgo:type C
type InterfaceVariantGetPtrBuiltinMethod func(VariantType, ConstStringNamePtr, Int) PtrBuiltInMethod

// llgo:type C
type InterfaceVariantGetPtrConstructor func(VariantType, c.Int32T) PtrConstructor

// llgo:type C
type InterfaceVariantGetPtrDestructor func(VariantType) PtrDestructor

// llgo:type C
type InterfaceVariantConstruct func(VariantType, UninitializedVariantPtr, *ConstVariantPtr, c.Int32T, *CallError)

// llgo:type C
type InterfaceVariantGetPtrSetter func(VariantType, ConstStringNamePtr) PtrSetter

// llgo:type C
type InterfaceVariantGetPtrGetter func(VariantType, ConstStringNamePtr) PtrGetter

// llgo:type C
type InterfaceVariantGetPtrIndexedSetter func(VariantType) PtrIndexedSetter

// llgo:type C
type InterfaceVariantGetPtrIndexedGetter func(VariantType) PtrIndexedGetter

// llgo:type C
type InterfaceVariantGetPtrKeyedSetter func(VariantType) PtrKeyedSetter

// llgo:type C
type InterfaceVariantGetPtrKeyedGetter func(VariantType) PtrKeyedGetter

// llgo:type C
type InterfaceVariantGetPtrKeyedChecker func(VariantType) PtrKeyedChecker

// llgo:type C
type InterfaceVariantGetConstantValue func(VariantType, ConstStringNamePtr, UninitializedVariantPtr)

// llgo:type C
type InterfaceVariantGetPtrUtilityFunction func(ConstStringNamePtr, Int) PtrUtilityFunction

// llgo:type C
type InterfaceStringNewWithLatin1Chars func(UninitializedStringPtr, *c.Char)

// llgo:type C
type InterfaceStringNewWithUtf8Chars func(UninitializedStringPtr, *c.Char)

// llgo:type C
type InterfaceStringNewWithUtf16Chars func(UninitializedStringPtr, *Char16T)

// llgo:type C
type InterfaceStringNewWithUtf32Chars func(UninitializedStringPtr, *Char32T)

// llgo:type C
type InterfaceStringNewWithWideChars func(UninitializedStringPtr, *c.Int32T)

// llgo:type C
type InterfaceStringNewWithLatin1CharsAndLen func(UninitializedStringPtr, *c.Char, Int)

// llgo:type C
type InterfaceStringNewWithUtf8CharsAndLen func(UninitializedStringPtr, *c.Char, Int)

// llgo:type C
type InterfaceStringNewWithUtf16CharsAndLen func(UninitializedStringPtr, *Char16T, Int)

// llgo:type C
type InterfaceStringNewWithUtf32CharsAndLen func(UninitializedStringPtr, *Char32T, Int)

// llgo:type C
type InterfaceStringNewWithWideCharsAndLen func(UninitializedStringPtr, *c.Int32T, Int)

// llgo:type C
type InterfaceStringToLatin1Chars func(ConstStringPtr, *c.Char, Int) Int

// llgo:type C
type InterfaceStringToUtf8Chars func(ConstStringPtr, *c.Char, Int) Int

// llgo:type C
type InterfaceStringToUtf16Chars func(ConstStringPtr, *Char16T, Int) Int

// llgo:type C
type InterfaceStringToUtf32Chars func(ConstStringPtr, *Char32T, Int) Int

// llgo:type C
type InterfaceStringToWideChars func(ConstStringPtr, *c.Int32T, Int) Int

// llgo:type C
type InterfaceStringOperatorIndex func(StringPtr, Int) *Char32T

// llgo:type C
type InterfaceStringOperatorIndexConst func(ConstStringPtr, Int) *Char32T

// llgo:type C
type InterfaceStringOperatorPlusEqString func(StringPtr, ConstStringPtr)

// llgo:type C
type InterfaceStringOperatorPlusEqChar func(StringPtr, Char32T)

// llgo:type C
type InterfaceStringOperatorPlusEqCstr func(StringPtr, *c.Char)

// llgo:type C
type InterfaceStringOperatorPlusEqWcstr func(StringPtr, *c.Int32T)

// llgo:type C
type InterfaceStringOperatorPlusEqC32str func(StringPtr, *Char32T)

// llgo:type C
type InterfaceStringResize func(StringPtr, Int) Int

// llgo:type C
type InterfaceStringNameNewWithLatin1Chars func(UninitializedStringNamePtr, *c.Char, Bool)

// llgo:type C
type InterfaceStringNameNewWithUtf8Chars func(UninitializedStringNamePtr, *c.Char)

// llgo:type C
type InterfaceStringNameNewWithUtf8CharsAndLen func(UninitializedStringNamePtr, *c.Char, Int)

// llgo:type C
type InterfaceXmlParserOpenBuffer func(ObjectPtr, *c.Uint8T, c.SizeT) Int

// llgo:type C
type InterfaceFileAccessStoreBuffer func(ObjectPtr, *c.Uint8T, c.Uint64T)

// llgo:type C
type InterfaceFileAccessGetBuffer func(ConstObjectPtr, *c.Uint8T, c.Uint64T) c.Uint64T

// llgo:type C
type InterfaceWorkerThreadPoolAddNativeGroupTask func(ObjectPtr, func(c.Pointer, c.Uint32T), c.Pointer, c.Int, c.Int, Bool, ConstStringPtr) c.Int64T

// llgo:type C
type InterfaceWorkerThreadPoolAddNativeTask func(ObjectPtr, func(c.Pointer), c.Pointer, Bool, ConstStringPtr) c.Int64T

// llgo:type C
type InterfacePackedByteArrayOperatorIndex func(TypePtr, Int) *c.Uint8T

// llgo:type C
type InterfacePackedByteArrayOperatorIndexConst func(ConstTypePtr, Int) *c.Uint8T

// llgo:type C
type InterfacePackedColorArrayOperatorIndex func(TypePtr, Int) TypePtr

// llgo:type C
type InterfacePackedColorArrayOperatorIndexConst func(ConstTypePtr, Int) TypePtr

// llgo:type C
type InterfacePackedFloat32ArrayOperatorIndex func(TypePtr, Int) *c.Float

// llgo:type C
type InterfacePackedFloat32ArrayOperatorIndexConst func(ConstTypePtr, Int) *c.Float

// llgo:type C
type InterfacePackedFloat64ArrayOperatorIndex func(TypePtr, Int) *c.Double

// llgo:type C
type InterfacePackedFloat64ArrayOperatorIndexConst func(ConstTypePtr, Int) *c.Double

// llgo:type C
type InterfacePackedInt32ArrayOperatorIndex func(TypePtr, Int) *c.Int32T

// llgo:type C
type InterfacePackedInt32ArrayOperatorIndexConst func(ConstTypePtr, Int) *c.Int32T

// llgo:type C
type InterfacePackedInt64ArrayOperatorIndex func(TypePtr, Int) *c.Int64T

// llgo:type C
type InterfacePackedInt64ArrayOperatorIndexConst func(ConstTypePtr, Int) *c.Int64T

// llgo:type C
type InterfacePackedStringArrayOperatorIndex func(TypePtr, Int) StringPtr

// llgo:type C
type InterfacePackedStringArrayOperatorIndexConst func(ConstTypePtr, Int) StringPtr

// llgo:type C
type InterfacePackedVector2ArrayOperatorIndex func(TypePtr, Int) TypePtr

// llgo:type C
type InterfacePackedVector2ArrayOperatorIndexConst func(ConstTypePtr, Int) TypePtr

// llgo:type C
type InterfacePackedVector3ArrayOperatorIndex func(TypePtr, Int) TypePtr

// llgo:type C
type InterfacePackedVector3ArrayOperatorIndexConst func(ConstTypePtr, Int) TypePtr

// llgo:type C
type InterfaceArrayOperatorIndex func(TypePtr, Int) VariantPtr

// llgo:type C
type InterfaceArrayOperatorIndexConst func(ConstTypePtr, Int) VariantPtr

// llgo:type C
type InterfaceArrayRef func(TypePtr, ConstTypePtr)

// llgo:type C
type InterfaceArraySetTyped func(TypePtr, VariantType, ConstStringNamePtr, ConstVariantPtr)

// llgo:type C
type InterfaceDictionaryOperatorIndex func(TypePtr, ConstVariantPtr) VariantPtr

// llgo:type C
type InterfaceDictionaryOperatorIndexConst func(ConstTypePtr, ConstVariantPtr) VariantPtr

// llgo:type C
type InterfaceObjectMethodBindCall func(MethodBindPtr, ObjectPtr, *ConstVariantPtr, Int, UninitializedVariantPtr, *CallError)

// llgo:type C
type InterfaceObjectMethodBindPtrcall func(MethodBindPtr, ObjectPtr, *ConstTypePtr, TypePtr)

// llgo:type C
type InterfaceObjectDestroy func(ObjectPtr)

// llgo:type C
type InterfaceGlobalGetSingleton func(ConstStringNamePtr) ObjectPtr

// llgo:type C
type InterfaceObjectGetInstanceBinding func(ObjectPtr, c.Pointer, *InstanceBindingCallbacks) c.Pointer

// llgo:type C
type InterfaceObjectSetInstanceBinding func(ObjectPtr, c.Pointer, c.Pointer, *InstanceBindingCallbacks)

// llgo:type C
type InterfaceObjectFreeInstanceBinding func(ObjectPtr, c.Pointer)

// llgo:type C
type InterfaceObjectSetInstance func(ObjectPtr, ConstStringNamePtr, ClassInstancePtr)

// llgo:type C
type InterfaceObjectGetClassName func(ConstObjectPtr, ClassLibraryPtr, UninitializedStringNamePtr) Bool

// llgo:type C
type InterfaceObjectCastTo func(ConstObjectPtr, c.Pointer) ObjectPtr

// llgo:type C
type InterfaceObjectGetInstanceFromId func(GDObjectInstanceID) ObjectPtr

// llgo:type C
type InterfaceObjectGetInstanceId func(ConstObjectPtr) GDObjectInstanceID

// llgo:type C
type InterfaceRefGetObject func(ConstRefPtr) ObjectPtr

// llgo:type C
type InterfaceRefSetObject func(RefPtr, ObjectPtr)

// llgo:type C
type InterfaceScriptInstanceCreate func(*ScriptInstanceInfo, ScriptInstanceDataPtr) ScriptInstancePtr

// llgo:type C
type InterfaceScriptInstanceCreate2 func(*ScriptInstanceInfo2, ScriptInstanceDataPtr) ScriptInstancePtr

// llgo:type C
type InterfacePlaceHolderScriptInstanceCreate func(ObjectPtr, ObjectPtr, ObjectPtr) ScriptInstancePtr

// llgo:type C
type InterfacePlaceHolderScriptInstanceUpdate func(ScriptInstancePtr, ConstTypePtr, ConstTypePtr)

// llgo:type C
type InterfaceObjectGetScriptInstance func(ConstObjectPtr, ObjectPtr) ScriptInstanceDataPtr

// llgo:type C
type InterfaceCallableCustomCreate func(UninitializedTypePtr, *CallableCustomInfo)

// llgo:type C
type InterfaceCallableCustomGetUserData func(ConstTypePtr, c.Pointer) c.Pointer

// llgo:type C
type InterfaceClassdbConstructObject func(ConstStringNamePtr) ObjectPtr

// llgo:type C
type InterfaceClassdbGetMethodBind func(ConstStringNamePtr, ConstStringNamePtr, Int) MethodBindPtr

// llgo:type C
type InterfaceClassdbGetClassTag func(ConstStringNamePtr) c.Pointer

// llgo:type C
type InterfaceClassdbRegisterExtensionClass func(ClassLibraryPtr, ConstStringNamePtr, ConstStringNamePtr, *ClassCreationInfo)

// llgo:type C
type InterfaceClassdbRegisterExtensionClass2 func(ClassLibraryPtr, ConstStringNamePtr, ConstStringNamePtr, *ClassCreationInfo2)

// llgo:type C
type InterfaceClassdbRegisterExtensionClassMethod func(ClassLibraryPtr, ConstStringNamePtr, *ClassMethodInfo)

// llgo:type C
type InterfaceClassdbRegisterExtensionClassIntegerConstant func(ClassLibraryPtr, ConstStringNamePtr, ConstStringNamePtr, ConstStringNamePtr, Int, Bool)

// llgo:type C
type InterfaceClassdbRegisterExtensionClassProperty func(ClassLibraryPtr, ConstStringNamePtr, *PropertyInfo, ConstStringNamePtr, ConstStringNamePtr)

// llgo:type C
type InterfaceClassdbRegisterExtensionClassPropertyIndexed func(ClassLibraryPtr, ConstStringNamePtr, *PropertyInfo, ConstStringNamePtr, ConstStringNamePtr, Int)

// llgo:type C
type InterfaceClassdbRegisterExtensionClassPropertyGroup func(ClassLibraryPtr, ConstStringNamePtr, ConstStringPtr, ConstStringPtr)

// llgo:type C
type InterfaceClassdbRegisterExtensionClassPropertySubgroup func(ClassLibraryPtr, ConstStringNamePtr, ConstStringPtr, ConstStringPtr)

// llgo:type C
type InterfaceClassdbRegisterExtensionClassSignal func(ClassLibraryPtr, ConstStringNamePtr, ConstStringNamePtr, *PropertyInfo, Int)

// llgo:type C
type InterfaceClassdbUnregisterExtensionClass func(ClassLibraryPtr, ConstStringNamePtr)

// llgo:type C
type InterfaceGetLibraryPath func(ClassLibraryPtr, UninitializedStringPtr)

// llgo:type C
type InterfaceEditorAddPlugin func(ConstStringNamePtr)

// llgo:type C
type InterfaceEditorRemovePlugin func(ConstStringNamePtr)
