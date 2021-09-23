#include "lua.h"
#include "lauxlib.h"

#include "srp.h"

#define MESSAGE "hello world!"

int say_hello(lua_State *L) {
		srp_verifier_delete(NULL);


    lua_pushstring(L, MESSAGE);
    return 1;
}

static const struct luaL_Reg functions [] = {
    {"say_hello", say_hello},
    {NULL, NULL}
};

int luaopen_mymod(lua_State *L) {
    luaL_register(L, "hello", functions);
    return 1;
}
