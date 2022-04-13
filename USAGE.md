# Usage

## API Usage

Must have [Go](https://golang.org/doc/install) installed.

`$ git clone https://gitlab.com/rileythomp14/voronoi`

`$ cd voronoi`

`$ ./build.sh`

`$ ./voronoi`

Then use Postman or some other API client with the following endpoints.

---

`GET /voronoi`: Creates a 512x512 png image of a Voronoi diagram.

`GET /moving`: Creates a 10 second mp4 video of a moving Voronoi diagram.

Query params:

`lines`: A boolean value to display black lines on region borders (default false).

`sites`: The number of regions to divide the image into (1-500, default 100).

`distance`: The distance algorithm to be used (one of `euclidean`, `manhattan`, `chebyshev`, default is `euclidean`).

`color1`: The first color to be used in the color gradient (default is `#123456`),

`color2`: The second color to be used in the color gradient (default is `#ffffff`).

---

`GET /fractal`: Creates a 512x512 png image of a fractal based on a T Square.

`GET /htree`: Creates a 512x512 png image of an H Tree.

`GET /tsquare`: Creates a 512x512 png image of a T Square.

Query params:

`splits`: The number of splits to perform (0-7, default 3).

---

`POST /stainedglass`: Creates a png stained glass image of the given image.

`POST /stainedglass2`: Creates a png stained glass image of the given image, with a better colouring algorithm.

`POST /quadimage`: Creates a png quadtree image of the given image.

`POST /circles`: Creates a png quadtree circle image of the given image.

Body:

A `jpg` or `png` image.

Query params:

`lines`: A boolean value to display black lines on region borders (default false).

`sites`: The number of regions to divide the image into (1-500, default 100).

`distance`: The distance algorithm to be used (one of `euclidean`, `manhattan`, `chebyshev`, default is `euclidean`).

---

## Old CLI Usage

Must have [Go](https://golang.org/doc/install) installed.

`$ git clone https://gitlab.com/rileythomp14/voronoi`

`$ cd voronoi`

`$ ./build.sh`

`$ ./voronoi <image> <sites> <dist-func> <animate-func> [sites-visible]`

`<image>`: the image name\
`<sites>`: 1-2000, number of sites in animation\
`<dist-func>`: e|m|c, euclidean(e), manhattan(m) or chebyshev(c) distance function used in animation\
`<animate-func>`: a|m|s|q, adding(a) sites, moving(m) sites, stained glass(s) or quadimage(q) animation created\
`[sites-visible]`: t|f, makes sites visible in animation. Default is false.\
`<arg>` are required, `[arg]` are optional

Stained glass and QuadImage animations will create a png image.

Adding and moving sites animations will creates 200 images in\
`frames/[0-1][0-9]{2}.png`.\
They can then be used with \
`$ ./mp4.sh <file-name>`\
to create `<file-name>`. Must have `ffmpeg` installed. Tested on a Mac.\
`$ ./gif.sh`\
to create `voronoi.gif`. Must have `imagemagick` installed. Tested on Linux.
