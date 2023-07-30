package internal

//go:generate echo proto generattion
//go:generate go generate ./pb

//go:generate echo swagger copy
//go:generate ./build.dev.sh

//go:generate echo client generation
//go:generate node server/ui/vue/gen.js

//go:generate echo ui build
//go:generate yarn   --cwd ./server/ui/vue/ build
//go:generate rm -rf ./server/ui/build
//go:generate mv ./server/ui/vue/dist ./server/ui/build
