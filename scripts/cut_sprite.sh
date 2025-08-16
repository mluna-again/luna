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

if [ -d "./sprites" ]; then
  cd sprites || exit
else
  cd ../sprites || exit
fi

if [ ! -r "../assets/$name" ]; then
  echo "$name sprite not found"
  exit 1
fi

sprite_dir="$(basename "$name" ".$name_ext")"
mkdir -p "$sprite_dir" || exit
cd "$sprite_dir" || exit

magick "../../assets/$name" -crop "$size" "tile.png"

for tile in tile*.png; do
  tile_no_ext=$(basename $tile .png)
  chafa --size "$size" "$tile" | tee "${tile_no_ext}.ascii"
done

rm tile*.png
