#!/bin/sh
GOOS=js GOARCH=wasm llgo build -tags=nogc -o ../game/builds/libgdsummator.web.template_debug.wasm32.wasm