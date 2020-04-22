package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"fractal/cmdflags"
	"fractal/types"
	"fractal/util"
	"html/template"
	"image"
	"image/png"
	"log"
	"net/http"
	"time"

	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func handleRequests() {

	router := mux.NewRouter()
	// CRUD
	router.HandleFunc("/image", getImage).Methods("get")

	router.HandleFunc("/tile", getTile).Methods("get")

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	router.HandleFunc("/", siteRootHandler).Methods("get")

	port := fmt.Sprintf(":%d", cmdflags.HttpPort)

	log.Println("Listenning on port: " + port)

	log.Fatal(http.ListenAndServe(port, router))
}

func getImage(w http.ResponseWriter, r *http.Request) {

	util.SetContentType(w, util.JSON{})
	json.NewEncoder(w).Encode(Character{Name: "Stas"})
	log.Printf("Successfully getImage with ")

}

func getTile(w http.ResponseWriter, r *http.Request) {
	log.Printf("getTile start")

	// for debug React.lazy
	time.Sleep(2 * time.Second)

	width, ok := r.URL.Query()["width"]
	if !ok || len(width[0]) < 1 {
		util.HandleError(w, "Param is not provided")
	}

	height, ok := r.URL.Query()["height"]
	if !ok || len(height[0]) < 1 {
		util.HandleError(w, "Param is not provided")
	}

	imageMaxX, _ := strconv.Atoi(width[0])
	imageMaxY, _ := strconv.Atoi(height[0])

	b := Borders{X1: -2.5, X2: 2.5, Y1: -2.5, Y2: 2.5}

	img := GenerateFractalImage(imageMaxX, imageMaxY, b)

	writeImageToResponse(w, &img)
	log.Printf("Successfully getTile")

}

func writeImageToResponse(w http.ResponseWriter, img *image.Image) {

	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, *img); err != nil {
		log.Println("unable to encode image.")
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}

func siteRootHandler(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	protoValue, ok := query["p"]

	// generate some QR code look a like image
	const imageMaxX = 80
	const imageMaxY = 80

	b := Borders{X1: -2.5, X2: 2.5, Y1: -2.5, Y2: 2.5}

	if ok {
		b = parseProto(protoValue[0])
	}

	log.Printf("Showing rectangle: b=%v", b)

	const plates = 10

	var imagesInfo types.ImagesInfo
	var imageRows []types.ImageRow

	for j := 0; j < plates; j++ {
		images := []types.Image{}
		for i := 0; i < plates; i++ {
			fragmentBorder := getFragmet(b, i, j, plates, plates)

			img := GenerateFractalImage(imageMaxX, imageMaxY, fragmentBorder)

			buffer := new(bytes.Buffer)
			if err := png.Encode(buffer, img); err != nil {
				log.Fatalln("unable to encode image.")
			}
			str := base64.StdEncoding.EncodeToString(buffer.Bytes())
			images = append(images, types.Image{Base64Coded: str, AltText: "New image", Params: "?" + getParams(fragmentBorder)})

		}
		imageRows = append(imageRows, types.ImageRow{Images: images})
	}

	imagesInfo.ImageRows = imageRows
	imagesInfo.ColumnsCount = plates
	imagesInfo.RowsCount = plates

	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Printf("Error while parsing template: %s\n", err.Error())
		return
	}
	err = t.Execute(w, imagesInfo)
	if err != nil {
		fmt.Printf("Error while processing template: %s\n", err.Error())
	}
}

type FileSystem struct {
	fs http.FileSystem
}

// Open opens file
func (fs FileSystem) Open(path string) (http.File, error) {
	f, err := fs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := strings.TrimSuffix(path, "/") + "/index.html"
		if _, err := fs.fs.Open(index); err != nil {
			return nil, err
		}
	}

	return f, nil
}
