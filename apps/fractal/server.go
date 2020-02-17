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
	"image/png"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func handleRequests() {

	router := mux.NewRouter()
	// CRUD
	router.HandleFunc("/image", getImage).Methods("get")

	// router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

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

func siteRootHandler(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()

	x1Value, ok := query["x1"]
	x2Value, _ := query["x2"]

	y1Value, _ := query["y1"]
	y2Value, _ := query["y2"]

	log.Printf("Processing params: x1=%v", x1Value)

	// generate some QR code look a like image
	const imageMaxX = 80
	const imageMaxY = 80

	b := Borders{X1: -2.5, X2: 2.5, Y1: -2.5, Y2: 2.5}
	if ok {
		x1, _ := strconv.ParseFloat(x1Value[0], 64)
		x2, _ := strconv.ParseFloat(x2Value[0], 64)
		y1, _ := strconv.ParseFloat(y1Value[0], 64)
		y2, _ := strconv.ParseFloat(y2Value[0], 64)
		b = Borders{X1: x1, X2: x2, Y1: y1, Y2: y2}
	}

	const plates = 10

	var imagesInfo types.ImagesInfo
	var imageRows []types.ImageRow

	for j := 0; j < plates; j++ {
		images := []types.Image{}
		for i := 0; i < plates; i++ {
			fragmentBorder := getFragmet(b, i, j, plates, plates)

			// log.Printf("Successfully getImage with %v", fragmentBorder)

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
