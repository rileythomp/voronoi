echo "Creating voronoi.gif..."
convert -delay 10 frames/{0..1}{0..9}{0..9}\.png voronoi.gif
echo "Finished creating voronoi.gif"
