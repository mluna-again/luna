#! /usr/bin/env bash

usage() {
  cat <<EOF
Note:
  this script only works only works if *all* sprites have the same size, if you have some exotic sheets
  you are going to have to kick some rocks.

Usage:

# Just put your sprites in luna/assets and run (from the project root)
$ ./luna/scripts/cut_sprites.sh
//go:embed sprites/cat/sleeping0_sm.ascii
var catSleeping0SM string
...
...
...

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
  pet="$4"

  magick "$file" -crop "$size" "tile.png" || return 1

  frame=0
  for tile in tile*.png; do
    small="${name}${frame}_sm.ascii"
    medium="${name}${frame}_md.ascii"
    large="${name}${frame}_xl.ascii"

    chafa --scale max --size 44x44 "$tile" > "$large" || return 1
    chafa --scale max --size 32x32 "$tile" > "$medium" || return 1
    chafa --scale max --size 20x20 "$tile" > "$small" || return 1
    cat <<EOF
//go:embed sprites/${pet}/${small}
var ${pet}${name^}${frame}SM string
//go:embed sprites/${pet}/${medium}
var ${pet}${name^}${frame}MD string
//go:embed sprites/${pet}/${large}
var ${pet}${name^}${frame}XL string
EOF

    frame=$(( frame + 1 ))
  done

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
