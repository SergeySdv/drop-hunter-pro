#!/bin/bash

for i in $(find pb/gen -type f -name '*.swagger.json')
do

  cp "$i" ui/vue/swagger/"$(basename -- "$i")"
  echo "$i"
done
