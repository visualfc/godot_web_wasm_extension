#! /bin/bash
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd $SCRIPT_DIR
# Build godot-cpp

if [ -d "game/builds" ]; then
    rm -rf game/builds
fi
unzip -o engine.zip -d game/builds

cd godot-cpp
scons platform=web 
cd ..
scons platform=web 

cp -r game/bin/libgdsummator.web.template_debug.wasm32.wasm game/builds/libgdsummator.web.template_debug.wasm32.wasm