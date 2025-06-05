#! /bin/bash
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd $SCRIPT_DIR

cd wasm-cpp
make wasm
cd ..

cp -r wasm-cpp/libgdsummator.web.template_debug.wasm32.wasm game/builds/libgdsummator.web.template_debug.wasm32.wasm