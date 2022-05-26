# voronoi

Computer art and animations created using [Voronoi diagrams](https://en.wikipedia.org/wiki/Voronoi_diagram) and [Quadtrees](https://en.wikipedia.org/wiki/Quadtree).

See [here](https://github.com/rileythomp/voronoi/blob/master/USAGE.md) for how to use.

Voronoi diagrams and quadtrees can be used together as a partitioning and colouring algorithm to give images a mosaic look.

| Original                        | Mosaic Image                       | Mosaic Gif              |                 
| ---                             | ---                                       | ---                            |     
| ![parrot](images/parrots.png)   | ![parrotg](images/stainglass/parrots.png) | ![parrotgif](clips/stainglass/parrots.gif) |
| ![treeimg](images/tree.png)     | ![treesg](images/stainglass/tree.png)     | ![treegif](clips/stainglass/tree.gif)      |
| ![chimpimg](images/chimp.png)   | ![chimpsg](images/stainglass/chimp.png)   | ![chimpgif](clips/stainglass/chimp.gif)   | 
| ![mbimg](images/mandelbrot.png) | ![mbsg](images/stainglass/mandelbrot.png) | ![mbgif](clips/stainglass/mandelbrot.gif)  |

See [Examples.md](https://gitlab.com/rileythomp14/voronoi/-/blob/master/Examples.md) for more examples.

A Voronoi diagram is an image divided into a given number of regions, where each region is defined by a single point (a "site") and all the points that are closer to that site than any other. Here is an example of adding sites (shown by a black dot) 1 by 1 to an image and giving each region a random colour.

<img src="clips/voronoi/adding.gif" width="420"/>

Here's 2 examples of 20 or so sites all moving in random directions. 

<img src="clips/voronoi/moving2.gif" width="420"/>
<img src="clips/voronoi/moving.gif" width="420"/>

[Here](https://gitlab.com/rileythomp14/voronoi/-/blob/master/clips/addmove.mp4) are another [two](https://gitlab.com/rileythomp14/voronoi/-/blob/master/clips/voronoi.mp4) interesting examples.


In the examples above, distance to a site was determined using [Euclidean distance](https://en.wikipedia.org/wiki/Euclidean_distance), but other methods can be used, such as the [Manhattan distance](https://en.wikipedia.org/wiki/Taxicab_geometry) or the [Chebyshev distance](https://en.wikipedia.org/wiki/Chebyshev_distance).

| Euclidean                               | Manhattan                               | Chebyshev                               |                 
| ---                                     | ---                                     | ---                                     |    
| ![euclidean2](images/voronoi/euclidean2.png)    | ![manhattan2](images/voronoi/manhattan2.png)    | ![chebyshev2](images/voronoi/chebyshev2.png)    | 
| ![euclidean](images/voronoi/euclidean.png)      | ![manhattan](images/voronoi/manhattan.png)      | ![chebyshev](images/voronoi/chebyshev.png)      |
| ![chimp](images/stainglass/chimp.png)   | ![chimp3](images/stainglass/chimp3.png) | ![chimp5](images/stainglass/chimp5.png) |
| ![chimp2](images/stainglass/chimp2.png) | ![chimp4](images/stainglass/chimp4.png) | ![chimp6](images/stainglass/chimp6.png) |
| ![tree](images/stainglass/tree.png)     | ![tree3](images/stainglass/tree3.png)   | ![tree5](images/stainglass/tree5.png)   |
| ![tree2](images/stainglass/tree2.png)   | ![tree4](images/stainglass/tree4.png)   | ![tree6](images/stainglass/tree6.png)   |

A quadtree is a tree data structure where each node has 4 children. It is used to represent the image by having each quadrant be the average colour of the image in that region. A metric similar to standard deviation for colour is used split the quadrant with the largest colour error. This has the effect of gradually increasing resolution. 

| Original                       | Quadtree Image                      | Quadtree Gif                         |                 
| ---                            | ---                                 | ---                                  |    
| ![tree](images/tree.png)       | ![qtree](images/quads/tree.png)     | ![qtreeg](clips/quads/tree.gif)       |
| ![chimp](images/chimp.png)     | ![qchimp](images/quads/chimp.png)   | ![qchimpg](clips/quads/chimp.gif)     | 
| ![parrots](images/parrots.png) | ![qparrot](images/quads/parrots.png) | ![qparrotg](clips/quads/parrots.gif) |

This algorithm can also be used with other shapes like circles.

| Original                       | Circle Image                      |         
| ---                            | ---                                 | 
| ![monalisa](images/monalisa.png)       | ![monalisacircle](images/circles/monalisa.png)     |
| ![lenna](images/lenna.png)       | ![lennacircle](images/circles/lenna.png)     |

The quadtree data structure can also be used to create other fractal animations

| Fractal                       | H Tree                      | T Square                          |                 
| ---                            | ---                         | ---                              |    
| ![fractal](clips/fractal.gif)  | ![htree](clips/htree.gif)   | ![tsquare](clips/tsquare.gif)    |


See [QuadImage](https://gitlab.com/rileythomp14/QuadImage) for more info.

