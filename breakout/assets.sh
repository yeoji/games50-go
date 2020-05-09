#!/bin/bash

####
# Auto generate the go:generate comments for assets
# Format: //go:generate file2byteslice -package=fonts -input=./fonts/font.ttf -output=./fonts/font.generated.go -var=Font_ttf

ASSET_PACKAGES=(
    "fonts"
    "sounds"
    "graphics"
)

for pkg in "${ASSET_PACKAGES[@]}"
do
    files=$(find assets/$pkg/* \! -name "*.go" -execdir echo {} ';')
    for file in $files
    do
        uppercased=$(echo $file | awk '{print toupper(substr($0,0,1))tolower(substr($0,2))}')
        extensionChanged=$(echo $uppercased | sed 's/\./_/')
        echo "//go:generate file2byteslice -package=$pkg -input=./$pkg/$file -output=./$pkg/$file.generated.go -var=$extensionChanged"
    done
done