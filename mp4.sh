!/bin/sh
# echo "Creating voronoi.mp4..."
# sleep 2
ffmpeg -framerate 10 -i frames/%03d.png -pix_fmt yuv420p $1
# sleep 1
# echo "Finished creating voronoi.mp4"
