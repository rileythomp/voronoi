package voronoi

// import (
// 	"fmt"
// 	"image"
// 	"image/color"
// 	"math/rand"
// 	"os"
// 	"strconv"
// 	"time"

// 	"gitlab.com/rileythomp14/voronoi/src/utils"
// )

// type ProgramArguments struct {
// 	Image        string
// 	Sites        int
// 	DistFunc     string
// 	AnimFunc     string
// 	SitesVisible bool
// }

// var (
// 	animationFuncs = map[string]AnimationFunc{
// 		"a": AddingSites,
// 		"m": MovingSites,
// 		"s": StainGlass0,
// 		"q": QuadImage0,
// 	}

// 	distanceFuncs = map[string]DistanceFunc{
// 		"e": EuclideanDistance,
// 		"m": ManhattanDistance,
// 		"c": ChebyshevDistance,
// 	}
// )

// func printUsage() {
// 	fmt.Println("Usage:")
// 	fmt.Println("./voronoi <image> <sites> <dist-func> <animate-func> [sites-visible]")
// 	fmt.Println("<image>: the image name")
// 	fmt.Println("<sites>: 1-2000, number of sites in animation")
// 	fmt.Println("<dist-func>: e|m|c, euclidean(e), manhattan(m) or chebyshev(c) distance function used in animation")
// 	fmt.Println("<animate-func>: a|m|s|q, adding(a) sites, moving(m) sites, stained glass(s) or quadimage(q) animation created")
// 	fmt.Println("[sites-visible]: t|f, makes sites visible in animation. Default is false.")
// 	fmt.Println("<arg> are required, [arg] are optional")
// }

// func main() {
// 	rand.Seed(time.Now().UTC().UnixNano())
// 	fmt.Println("Image Manipulation Art")

// 	args, err := parseProgramArgs()
// 	if err != nil {
// 		printUsage()
// 		if err.Error() != "help" {
// 			fmt.Printf("There was an error parsing the program arguments: %s\n", err.Error())
// 		}
// 		return
// 	}

// 	err = Run(args.Image, args.DistFunc, args.AnimFunc, args.Sites, args.SitesVisible)
// 	if err != nil {
// 		fmt.Printf("There was an error running the program: %s\n", err)
// 		return
// 	}
// }

// func Run(image, distFunc, animFunc string, numSites int, sitesVisible bool) error {
// 	var (
// 		distance  DistanceFunc
// 		animation AnimationFunc
// 		ok        bool
// 	)

// 	if distance, ok = distanceFuncs[distFunc]; !ok {
// 		errMsg := "Distance function should be one of:\n"
// 		errMsg += "e for Euclidean distance\n"
// 		errMsg += "m for Manhattan distance\n"
// 		errMsg += "c for Chebyshev distance"
// 		return fmt.Errorf(errMsg)
// 	}

// 	if animation, ok = animationFuncs[animFunc]; !ok {
// 		errMsg := "Distance function should be one of:\n"
// 		errMsg += "a for adding sites\n"
// 		errMsg += "m for moving sites\n"
// 		errMsg += "s for stained glass\n"
// 		errMsg += "q for quadtree images"
// 		return fmt.Errorf(errMsg)
// 	}

// 	animation(image, numSites, sitesVisible, distance)

// 	return nil
// }

// func StainGlass0(imgName string, numSites int, sitesVisible bool, distFunc DistanceFunc) error {
// 	inputImg, err := utils.GetImageFromPath(imgName)
// 	if err != nil {
// 		return fmt.Errorf("image %s was not found", imgName)
// 	}
// 	tl, br := inputImg.Bounds().Min, inputImg.Bounds().Max
// 	height, width := br.Y-tl.Y, br.X-tl.X
// 	img := utils.ImageToRGBA(inputImg)
// 	qt := NewQuadtree("root", tl, br, 0, img)
// 	sg := StainedGlass{img, []Site{}, qt}
// 	splits := utils.Min(MaxSites, numSites)
// 	for i := 0; i < splits; i++ {
// 		sg.sites = sg.qt.splitTree(img)
// 		// createVoronoiImg(width, height, sitesVisible, sg.sites, distFunc, splits)
// 	}
// 	return createVoronoiImg(splits, width, height, ".", sitesVisible, sg.sites, distFunc)
// }

// func QuadImage0(imgName string, numSites int, sitesVisible bool, distFunc DistanceFunc) error {
// 	inputImg, err := utils.GetImageFromPath(imgName)
// 	if err != nil {
// 		return fmt.Errorf("image %s was not found", imgName)
// 	}
// 	tl, br := inputImg.Bounds().Min, inputImg.Bounds().Max
// 	img := utils.ImageToRGBA(inputImg)
// 	qt := NewQuadtree("root", tl, br, 0, img)
// 	splits := utils.Min(MaxSites, numSites)
// 	for i := 0; i < splits; i++ {
// 		qt.splitTree(img)
// 		// qt.createQuadImg(i, sitesVisible)
// 	}
// 	return qt.createQuadImg(splits, sitesVisible)
// }

// func AddingSites(imgName string, numSites int, sitesVisible bool, distFunc DistanceFunc) error {
// 	var sites []Site
// 	width, height := DefaultWidth, DefaultHeight
// 	tl, br := image.Point{0, 0}, image.Point{width, height}
// 	c1, c2 := getGradientColors()
// 	// Generate gifs of adding sites
// 	numSites = utils.Min(DefaultFrames, numSites)
// 	for i := 0; i < numSites; i++ {
// 		sites = append(sites, NewSite(tl, br, c1, c2))
// 		createVoronoiImg(i, width, height, "frames", sitesVisible, sites, distFunc)
// 	}
// 	return nil
// 	// return createVoronoiImg(numSites, width, height, sitesVisible, sites, distFunc)
// }

// func MovingSites(imgName string, numSites int, sitesVisible bool, distFunc DistanceFunc) error {
// 	var sites []Site
// 	width, height := DefaultWidth, DefaultHeight
// 	tl, br := image.Point{0, 0}, image.Point{width, height}
// 	c1, c2 := getGradientColors()
// 	// Creates frames of sites moving
// 	for i := 0; i < numSites; i++ {
// 		sites = append(sites, NewSite(tl, br, c1, c2))
// 	}
// 	for j := 0; j < DefaultFrames; j++ {
// 		for i := range sites {
// 			sites[i].x += sites[i].slope.dx
// 			if sites[i].x < 0 {
// 				sites[i].x = width - 1
// 			} else if sites[i].x >= width {
// 				sites[i].x = 0
// 			}
// 			sites[i].y += sites[i].slope.dy
// 			if sites[i].y < 0 {
// 				sites[i].y = height - 1
// 			} else if sites[i].y >= height {
// 				sites[i].y = 0
// 			}
// 		}
// 		createVoronoiImg(j, width, height, "frames", sitesVisible, sites, distFunc)
// 	}
// 	return nil
// 	// return createVoronoiImg(numSites, width, height, sitesVisible, sites, distFunc)
// }

// func getGradientColors() (string, string) {
// 	fmt.Println("Enter gradient colors for voronoi diagram")
// 	fmt.Println("Use 'red', 'blue', or 'green' for default gradients")
// 	fmt.Println("or hit enter to just use random colors.")
// 	fmt.Println("Gradient colors must be in hex [0-9A-Fa-f]{6}")
// 	c1 := utils.GetUserInput("Color 1: ")
// 	c2 := utils.GetUserInput("Color 2: ")
// 	return c1, c2
// }

// func parseProgramArgs() (ProgramArguments, error) {
// 	args := ProgramArguments{}

// 	if len(os.Args) == 1 || utils.Contains(os.Args, "-h") || utils.Contains(os.Args, "--help") {
// 		return args, fmt.Errorf("help")
// 	} else if len(os.Args) < 5 {
// 		return args, fmt.Errorf("insufficient arguments")
// 	}

// 	args.Image = os.Args[1]
// 	_, err := utils.GetImageFromPath(args.Image)
// 	if err != nil {
// 		return args, fmt.Errorf("image %s was not found", args.Image)
// 	}

// 	args.Sites, err = strconv.Atoi(os.Args[2])
// 	if err != nil || args.Sites < 1 {
// 		return args, fmt.Errorf("invalid number of sites")
// 	}

// 	args.DistFunc = os.Args[3]
// 	args.AnimFunc = os.Args[4]

// 	if len(os.Args) > 5 {
// 		if args.SitesVisible, err = strconv.ParseBool(os.Args[5]); err != nil {
// 			return args, fmt.Errorf("invalid boolean")
// 		}
// 	}

// 	return args, nil
// }

// func (qt *Quadtree) createQuadImg(n int, lines bool) error {
// 	Black := color.RGBA{0, 0, 0, 0xff}
// 	img := image.NewRGBA(image.Rectangle{qt.upLeft, qt.lowRight})
// 	var nodeStack Stack
// 	nodeStack.Push(qt)
// 	for !nodeStack.IsEmpty() {
// 		curNode, _ := nodeStack.Pop()
// 		// is leaf node
// 		if curNode.nw == nil {
// 			for y := curNode.upLeft.Y; y < curNode.lowRight.Y; y++ {
// 				for x := curNode.upLeft.X; x < curNode.lowRight.X; x++ {
// 					if lines && (y == curNode.upLeft.Y || y == curNode.lowRight.Y || x == curNode.upLeft.X || x == curNode.lowRight.X) {
// 						img.Set(x, y, Black)
// 					} else {
// 						img.Set(x, y, curNode.color)
// 					}
// 				}
// 			}
// 		} else {
// 			// not leaf, add children to stack
// 			nodeStack.Push(curNode.nw)
// 			nodeStack.Push(curNode.ne)
// 			nodeStack.Push(curNode.se)
// 			nodeStack.Push(curNode.sw)
// 		}
// 	}
// 	return utils.OutputPngFrame(".", n, img)
// }
