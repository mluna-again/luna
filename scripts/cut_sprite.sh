#! /usr/bin/env bash

file="$1"
file_ext="${file##*.}"
name="$2"
animation="$3"
size="${4:-32x32}"

usage() {
  cat <<EOF
Usage:

# Put your raw sprite in assets/ and run
$ cut_sprite.sh <sprite_file> <sprite_name> <animation_name> [<size>]

# With a specific size (default: 32x32)
$ cut_sprite.sh cute_and_big_bunny.png big_bunny sleeping 64x64
EOF
  exit 1
}

[ -z "$file" ] && usage
[ -z "$name" ] && usage
[ -z "$animation" ] && usage

if [ -d "./sprites" ]; then
  cd sprites || exit
else
  cd ../sprites || exit
fi

if [ ! -r "../assets/$file" ]; then
  echo "$file sprite not found"
  exit 1
fi

sprite_dir="$name"
mkdir -p "$sprite_dir" || exit
cd "$sprite_dir" || exit

magick "../../assets/$file" -crop "$size" "tile.png"

for tile in tile*.png; do
  filename=$(sed "s|tile|$animation|" <<< "$tile")
  chafa --size "$size" "$tile" | tee "$filename"
done

rm tile*.png
