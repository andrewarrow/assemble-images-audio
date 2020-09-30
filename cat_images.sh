#!/bin/sh
rm video/*
ffmpeg -framerate 1 -pattern_type glob -i '../aae1/*.png' -c:v libx264 -r 30 -pix_fmt yuv420p video/out.mp4

ffmpeg -i video/out.mp4 -i ../aae1/audio/001.m4a -c copy -map 0:v:0 -map 1:a:0 video/final.mov