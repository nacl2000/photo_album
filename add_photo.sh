#!/bin/bash
location="$1"
rm -rf /tmp/img
mkdir /tmp/img
mv /home/nacl/work/*.JPG /tmp/img
if [[ ! -d /home/nacl/compress_photo_data/$location ]]; then
    mkdir /home/nacl/compress_photo_data/$location
fi

if [[ ! -d /home/nacl/photo_data/$location/ ]]; then
    mkdir /home/nacl/photo_data/$location/
fi

collie -r /tmp/img/ -o "/home/nacl/compress_photo_data/$location/" -q 30
mv /tmp/img/*.JPG "/home/nacl/photo_data/$location/"
