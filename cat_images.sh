#!/bin/sh

ffmpeg -framerate 1 -pattern_type glob -i 'video/*.png' -c:v libx264 -r 30 -pix_fmt yuv420p video/out.mp4

ffmpeg -i video/out.mp4 -i ../aae1/audio/long.wav -c copy -map 0:v:0 -map 1:a:0 video/final.mov
