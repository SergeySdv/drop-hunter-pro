#!/bin/bash

rm -rf server/ui/vue/swagger
mkdir server/ui/vue/swagger

for i in $(find pb/gen -type f -name '*.swagger.json')
do

  cp "$i" server/ui/vue/swagger/"$(basename -- "$i")"
  echo "$i"
done
