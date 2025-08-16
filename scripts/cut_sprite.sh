#! /usr/bin/env bash

name="$1"
name_ext="${name##*.}"
size="${2:-32x32}"

usage() {
  cat <<EOF
Usage:

# Put your raw sprite in assets/ and run
$ cut_sprite.sh <sprite_name> [<size>]

# With a specific size (default: 32x32)
$ cut_sprite.sh cute_and_big_bunny.png 64x64
EOF
  exit 1
}

[ -z "$name" ] && usage

if [ -d "./assets" ]; then
  cd assets || exit
else
  cd ../assets || exit
fi

if [ ! -r "$name" ]; then
  echo "$name sprite not found"
  exit 1
fi

magick "$name" -crop "$size" tile.png
