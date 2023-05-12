package diary

//go:generate echo proto generattion
//go:generate go generate ./pb

//go:generate echo swagger copy
//go:generate ./build.dev.sh

//go:generate echo client generation
//go:generate node ui/vue/gen.js

//go:generate echo ui build
//go:generate yarn   --cwd ./ui/vue/ build
//go:generate rm -rf ./ui/build
//go:generate mv ./ui/vue/dist ./ui/build
