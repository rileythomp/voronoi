package handlers

import (
	"bytes"
	b64 "encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"net/http"
	"strconv"

	"strings"

	"gitlab.com/rileythomp14/voronoi/src/utils"
)

func RespondWithVideo(w http.ResponseWriter, r *http.Request, vidPath, format string, respCode int) error {
	w.Header().Set("Content-Type", "video/"+format)
	http.ServeFile(w, r, vidPath)
	return nil
}

func RespondWithImage(w http.ResponseWriter, img image.Image, format string, respCode int) error {
	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, img, nil); err != nil {
		return err
	}
	data := buf.Bytes()
	w.Header().Set("Content-Type", "image/"+format)
	w.WriteHeader(respCode)
	w.Write(data)
	return nil
}

func RespondWithError(w http.ResponseWriter, err error, respCode int) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(respCode)
	w.Write([]byte(err.Error()))
}

func getVoronoiParams(query string) (VoronoiParams, error) {
	var err error
	vp := VoronoiParams{
		lines:    false,
		sites:    100,
		distance: "euclidean",
		color1:   "123456",
		color2:   "ffffff",
		splits:   3,
	}
	params := strings.Split(query, "&")
	for _, param := range params {
		kv := strings.Split(param, "=")
		switch kv[0] {
		case "lines":
			if vp.lines, err = strconv.ParseBool(kv[1]); err != nil {
				return vp, err
			}
		case "sites":
			if vp.sites, err = strconv.Atoi(kv[1]); err != nil {
				return vp, err
			}
		case "splits":
			if vp.splits, err = strconv.Atoi(kv[1]); err != nil {
				return vp, err
			}
		case "distance":
			if utils.Contains([]string{"euclidean", "manhattan", "chebyshev"}, kv[1]) {
				vp.distance = kv[1]
			}
		case "color1":
			vp.color1 = kv[1]
		case "color2":
			vp.color2 = kv[1]
		}
	}
	return vp, nil
}

func getJWT(r *http.Request, schema string) (string, error) {
	auth := r.Header.Values("Authorization")
	if len(auth) < 1 {
		return "", fmt.Errorf("no authorization header found")
	}
	token := strings.Split(auth[0], schema+" ")
	if len(token) < 2 {
		return "", fmt.Errorf("no token found")
	}
	return token[1], nil
}

func getJWTPayload(jwt string) (string, error) {
	encParts := strings.Split(jwt, ".")
	if len(encParts) < 2 {
		return "", fmt.Errorf("no payload in jwt")
	}
	decoded, err := b64.RawURLEncoding.DecodeString(encParts[1])
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}
