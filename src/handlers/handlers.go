package handlers

import (
	"fmt"
	"image"
	"net/http"

	"gitlab.com/rileythomp14/voronoi/src/utils"
	"gitlab.com/rileythomp14/voronoi/src/voronoi"
)

type (
	VoronoiParams struct {
		lines    bool
		sites    int
		splits   int
		distance string
		color1   string
		color2   string
	}

	Route struct {
		Name        string
		Method      string
		Pattern     string
		HandlerFunc http.HandlerFunc
	}
)

const (
	QueryParamErr  = "error getting query parameters: %s"
	ImgDecodingErr = "error decoding request image: %s"
	ImgResponseErr = "error responding with image: %s"
	VidResponseErr = "error responding with video: %s"
	ImgTooLargeErr = "image must be less than 1000px in height and width"
	Bearer         = "Bearer"
)

var logger = NewLogger()

func GetRoutes() []Route {
	return []Route{
		{
			Name:        "Voronoi Diagram",
			Method:      http.MethodGet,
			Pattern:     "/voronoi",
			HandlerFunc: VoronoiDiagram,
		},
		{
			Name:        "Moving Voronoi",
			Method:      http.MethodGet,
			Pattern:     "/moving",
			HandlerFunc: MovingVoronoi,
		},
		{
			Name:        "Stained Glass - Region is average color of site quadrant",
			Method:      http.MethodPost,
			Pattern:     "/stainedglass",
			HandlerFunc: StainedGlass,
		},
		{
			Name:        "Stained Glass 2 - Region is average color of its pixels",
			Method:      http.MethodPost,
			Pattern:     "/stainedglass2",
			HandlerFunc: StainedGlass2,
		},
		{
			Name:        "Quadtree Image",
			Method:      http.MethodPost,
			Pattern:     "/quadimage",
			HandlerFunc: QuadImage,
		},
		{
			Name:        "Circles",
			Method:      http.MethodPost,
			Pattern:     "/circles",
			HandlerFunc: DrawCircles,
		},
		{
			Name:        "H-Tree",
			Method:      http.MethodGet,
			Pattern:     "/htree",
			HandlerFunc: DrawHTree,
		},
		{
			Name:        "T-Square",
			Method:      http.MethodGet,
			Pattern:     "/tsquare",
			HandlerFunc: DrawTSquare,
		},
		{
			Name:        "Fractal",
			Method:      http.MethodGet,
			Pattern:     "/fractal",
			HandlerFunc: DrawFractal,
		},
	}
}

func VoronoiDiagram(w http.ResponseWriter, r *http.Request) {
	logger.Infof("received request to create voronoi diagram")
	vp, err := getVoronoiParams(r.URL.RawQuery)
	if err != nil {
		logger.Errorf(QueryParamErr, err)
		RespondWithError(w, fmt.Errorf(QueryParamErr, err), http.StatusBadRequest)
		return
	}
	logger.Infof("creating %s voronoi diagram with %d sites", vp.distance, vp.sites)
	vImg := voronoi.VoronoiDiagram(vp.sites, vp.lines, vp.color1, vp.color2, vp.distance)
	if err := RespondWithImage(w, vImg, "png", http.StatusCreated); err != nil {
		logger.Errorf(ImgResponseErr, err)
		RespondWithError(w, fmt.Errorf(ImgResponseErr, err), http.StatusInternalServerError)
		return
	}
	logger.Infof("created %s voronoi diagram with %d sites", vp.distance, vp.sites)
}

func DrawHTree(w http.ResponseWriter, r *http.Request) {
	logger.Infof("received request to create htree image")
	vp, err := getVoronoiParams(r.URL.RawQuery)
	if err != nil {
		logger.Errorf(QueryParamErr, err)
		RespondWithError(w, fmt.Errorf(QueryParamErr, err), http.StatusBadRequest)
		return
	}
	logger.Infof("creating htree image")
	vp.splits = utils.Min(utils.Max(vp.splits, 0), 7)
	vImg := voronoi.HTreeImage(vp.splits)
	if err := RespondWithImage(w, vImg, "png", http.StatusCreated); err != nil {
		logger.Errorf(ImgResponseErr, err)
		RespondWithError(w, fmt.Errorf(ImgResponseErr, err), http.StatusInternalServerError)
		return
	}
	logger.Infof("created htree image")
}

func DrawTSquare(w http.ResponseWriter, r *http.Request) {
	logger.Infof("received request to create tsquare image")
	vp, err := getVoronoiParams(r.URL.RawQuery)
	if err != nil {
		logger.Errorf(QueryParamErr, err)
		RespondWithError(w, fmt.Errorf(QueryParamErr, err), http.StatusBadRequest)
		return
	}
	logger.Infof("creating tsquare image")
	vp.splits = utils.Min(utils.Max(vp.splits, 0), 7)
	vImg := voronoi.TSquareImage(vp.splits)
	if err := RespondWithImage(w, vImg, "png", http.StatusCreated); err != nil {
		logger.Errorf(ImgResponseErr, err)
		RespondWithError(w, fmt.Errorf(ImgResponseErr, err), http.StatusInternalServerError)
		return
	}
	logger.Infof("created tsquare image")
}

func DrawFractal(w http.ResponseWriter, r *http.Request) {
	logger.Infof("received request to create fractal image")
	vp, err := getVoronoiParams(r.URL.RawQuery)
	if err != nil {
		logger.Errorf(QueryParamErr, err)
		RespondWithError(w, fmt.Errorf(QueryParamErr, err), http.StatusBadRequest)
		return
	}
	logger.Infof("creating fractal image")
	vp.splits = utils.Min(utils.Max(vp.splits, 0), 7)
	vImg := voronoi.FractalImage(vp.splits)
	if err := RespondWithImage(w, vImg, "png", http.StatusCreated); err != nil {
		logger.Errorf(ImgResponseErr, err)
		RespondWithError(w, fmt.Errorf(ImgResponseErr, err), http.StatusInternalServerError)
		return
	}
	logger.Infof("created fractal image")
}

func MovingVoronoi(w http.ResponseWriter, r *http.Request) {
	logger.Infof("received request to create moving voronoi")
	vp, err := getVoronoiParams(r.URL.RawQuery)
	if err != nil {
		logger.Errorf(QueryParamErr, err)
		RespondWithError(w, fmt.Errorf(QueryParamErr, err), http.StatusBadRequest)
		return
	}
	logger.Infof("creating %s moving voronoi with %d sites", vp.distance, vp.sites)
	vidPath, format := voronoi.MovingVoronoi(vp.sites, vp.lines, vp.color1, vp.color2, vp.distance)
	if err := RespondWithVideo(w, r, vidPath, format, http.StatusCreated); err != nil {
		logger.Errorf(VidResponseErr, err)
		RespondWithError(w, fmt.Errorf(VidResponseErr, err), http.StatusInternalServerError)
		return
	}
	logger.Infof("created %s moving voronoi with %d sites", vp.distance, vp.sites)
}

func StainedGlass(w http.ResponseWriter, r *http.Request) {
	logger.Infof("received request to create stained glass")
	img, _, err := image.Decode(r.Body)
	if err != nil {
		logger.Errorf(ImgDecodingErr, err)
		RespondWithError(w, fmt.Errorf(ImgDecodingErr, err), http.StatusBadRequest)
		return
	}
	vp, err := getVoronoiParams(r.URL.RawQuery)
	if err != nil {
		logger.Errorf(QueryParamErr, err)
		RespondWithError(w, fmt.Errorf(QueryParamErr, err), http.StatusBadRequest)
		return
	}
	logger.Infof("creating %s stained glass with %d sites", vp.distance, vp.sites)
	sg := voronoi.StainGlass(img, vp.sites, vp.lines, vp.distance)
	if err := RespondWithImage(w, sg, "png", http.StatusCreated); err != nil {
		logger.Errorf(ImgResponseErr, err)
		RespondWithError(w, fmt.Errorf(ImgResponseErr, err), http.StatusInternalServerError)
	}
	logger.Infof("created %s stained glasss with %d sites", vp.distance, vp.sites)
}

func StainedGlass2(w http.ResponseWriter, r *http.Request) {
	logger.Infof("received request to create stained glass 2")
	img, _, err := image.Decode(r.Body)
	if err != nil {
		logger.Errorf(ImgDecodingErr, err)
		RespondWithError(w, fmt.Errorf(ImgDecodingErr, err), http.StatusBadRequest)
		return
	}
	if img.Bounds().Dx() > 1000 || img.Bounds().Dy() > 1000 {
		logger.Errorf(ImgTooLargeErr)
		RespondWithError(w, fmt.Errorf(ImgTooLargeErr), http.StatusBadRequest)
		return
	}
	vp, err := getVoronoiParams(r.URL.RawQuery)
	if err != nil {
		logger.Errorf(QueryParamErr, err)
		RespondWithError(w, fmt.Errorf(QueryParamErr, err), http.StatusBadRequest)
		return
	}
	logger.Infof("creating %s stained glass 2 with %d sites", vp.distance, vp.sites)
	sg := voronoi.StainGlass2(img, vp.sites, vp.lines, vp.distance)
	if err := RespondWithImage(w, sg, "png", http.StatusCreated); err != nil {
		logger.Errorf(ImgResponseErr, err)
		RespondWithError(w, fmt.Errorf(ImgResponseErr, err), http.StatusInternalServerError)
	}
	logger.Infof("created %s stained glasss 2 with %d sites", vp.distance, vp.sites)
}

func QuadImage(w http.ResponseWriter, r *http.Request) {
	logger.Infof("received request to create quad image")
	img, _, err := image.Decode(r.Body)
	if err != nil {
		logger.Errorf(ImgDecodingErr, err)
		RespondWithError(w, fmt.Errorf(ImgDecodingErr, err), http.StatusBadRequest)
		return
	}
	vp, err := getVoronoiParams(r.URL.RawQuery)
	if err != nil {
		logger.Errorf(QueryParamErr, err)
		RespondWithError(w, fmt.Errorf(QueryParamErr, err), http.StatusBadRequest)
		return
	}
	logger.Infof("creating quad image with %d sites", vp.sites)
	qi := voronoi.QuadImage(img, vp.sites, vp.lines)
	if err := RespondWithImage(w, qi, "png", http.StatusCreated); err != nil {
		logger.Errorf(ImgResponseErr, err)
		RespondWithError(w, fmt.Errorf(ImgResponseErr, err), http.StatusInternalServerError)
	}
	logger.Infof("created quad image with %d sites", vp.sites)
}

func DrawCircles(w http.ResponseWriter, r *http.Request) {
	logger.Infof("received request to draw circles")
	img, _, err := image.Decode(r.Body)
	if err != nil {
		logger.Errorf(ImgDecodingErr, err)
		RespondWithError(w, fmt.Errorf(ImgDecodingErr, err), http.StatusBadRequest)
		return
	}
	vp, err := getVoronoiParams(r.URL.RawQuery)
	if err != nil {
		logger.Errorf(QueryParamErr, err)
		RespondWithError(w, fmt.Errorf(QueryParamErr, err), http.StatusBadRequest)
		return
	}
	logger.Infof("creating circles with %d sites", vp.sites)
	circlesImg := voronoi.DrawCircles(img, vp.sites, vp.lines)
	if err := RespondWithImage(w, circlesImg, "png", http.StatusCreated); err != nil {
		logger.Errorf(ImgResponseErr, err)
		RespondWithError(w, fmt.Errorf(ImgResponseErr, err), http.StatusInternalServerError)
		return
	}
	logger.Infof("created circles with %d sites", vp.sites)
}
