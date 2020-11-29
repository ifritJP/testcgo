package testcgo

// #include <string.h>
// #include <stdlib.h>
// #cgo CFLAGS: -I/usr/include/lua5.3
// #cgo LDFLAGS: -ldl -lm -llua5.3
// #include <lauxlib.h>
// #include <lualib.h>
import "C"

import "fmt";

func Sub() {
    fmt.Println( "hogea:5" )
}

type lua_int = C.longlong
type lua_num = C.double
type lua_bool = C.int

func luaL_newstate() *C.lua_State {
    return C.luaL_newstate()
}
func lua_close(vm *C.lua_State) {
    C.lua_close( vm )
}
func lua_version(vm *C.lua_State) lua_num {
    return *C.lua_version( vm )
}

