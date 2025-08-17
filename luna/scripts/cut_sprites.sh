#! /usr/bin/env bash

usage() {
  cat <<EOF
Note:
  this script only works only works if *all* sprites have the same size, if you have some exotic sheets
  you are going to have to kick some rocks.

Usage:

# Just put your sprites in luna/assets and run (from the project root)
$ ./luna/scripts/cut_sprites.sh

# If your sprites are not standard size (32x32) you can specify the size
$ ./luna/scripts/cut_sprites.sh --size 64x64
EOF
  exit 1
}

_SIZE="32x32"

while [ -n "$1" ]; do
  case "${1,,}" in
    --size)
      shift
      _SIZE="$1"
      ;;

    --help|-help|-h|help)
      usage
      ;;
  esac
  shift
done

cd luna/assets || exit

assets_dir=$(pwd)

make_sprites() {
  local file size name action
  file="$1"
  name="$2"
  size="$3"
  action="$4"

  printf "Making $name $action... "

  magick "$file" -crop "$size" "tile.png" || return 1

  frame=0
  for tile in tile*.png; do
    chafa --size 32x32 "$tile" > "${name}${frame}_xl.ascii" || return 1
    chafa --size 16x16 "$tile" > "${name}${frame}_md.ascii" || return 1
    chafa --size 12x12 "$tile" > "${name}${frame}_sm.ascii" || return 1

    frame=$(( frame + 1 ))
  done

  echo "done."

  rm *.png
}

for action in idle sleeping attacking; do
  for pet in $assets_dir/*_${action}.png; do
    name=$(basename "$pet" _${action}.png) || exit
    cd ../sprites || exit
    if [ ! -d "$name" ]; then
      mkdir "$name" || exit
    fi
    cd "$name" || exit

    no_ext=$(basename "$file" .png) || exit
    make_sprites "$pet" "$action" "$_SIZE" "$name" || exit

    cd "$assets_dir" || exit
  done
done
