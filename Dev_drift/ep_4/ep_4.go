package ep4

var i8 int8   //-2**8 -> 2**8 -1
var i16 int16 //-2**16 -> 2**16 -1
var i32 int32 //-2**32 -> 2**32 -1
var i64 int64 //-2**64 -> 2**64 -1

var ui8 uint8   // 0 -> 2**8 - 1
var ui16 uint16 // 0 -> 2**16 - 1
var ui32 uint32 // 0 -> 2**32 - 1
var ui64 uint64 // 0 -> 2**64 - 1

var bt byte // Синоним uint8
var rn rune // Синоним int32
var i int   // Синоним int32
var ui uint // Синоним uint32/64

var f32 float32 // 1.4 * 10**-45 -> 3.4 * 10 **38
var f64 float64 // 4.9 * 10 ** -324 ->  1.8 * 10 ** 308

var comp64 complex64
var comp128 complex128 // 1.4 * 10**-45 -> 3.4 * 10 **38

var blt bool = true
var blf bool = false

var str1 string = `str 1`
var str2 string = "str 2"
