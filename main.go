package main

import (
	"bytes"
	"image"
	"log"
	"mime"
	"strconv"
	"net/http"
	"image/jpeg"
	"path/filepath"
	"gopkg.in/mgo.v2"
	"github.com/gorilla/mux"
	"github.com/nfnt/resize"
	"github.com/alsey/mongo-image-server/config"
	"github.com/alsey/mongo-image-server/health"
	"github.com/alsey/mongo-image-server/logger"
)

func main() {

	var (
		info    *mgo.DialInfo
		session *mgo.Session
		gfs     *mgo.GridFS
		err     error
	)

	info = &mgo.DialInfo{
		Addrs:       []string{config.GetMongoAddr()},
		Database:    config.GetMongoDB(),
		Source:      config.GetMongoDB(),
		ServiceHost: config.GetMongoAddr(),
		Username:    config.GetMongoUser(),
		Password:    config.GetMongoPassword(),
	}

	logger.Info("connection info is %v", info)

	if session, err = mgo.DialWithInfo(info); nil != err {
		logger.Fatal("failed to connect mongodb", err)
		return
	}

	defer session.Close()

	gfs = session.DB(config.GetMongoDB()).GridFS("fs")

	r := mux.NewRouter()

	r.HandleFunc("/health", health.Health)
	r.HandleFunc("/env", health.Env)
	r.HandleFunc("/favicon.ico", health.Favicon)

	r.HandleFunc("/images/{filename}", func(w http.ResponseWriter, r *http.Request) {

		var (
			file    *mgo.GridFile
			content []byte
			status  = 200
			errmsg  string
			err     error
		)

		defer func() {
			if nil != err {
				http.Error(w, err.Error(), status)
			}
			if status != 200 {
				http.Error(w, errmsg, status)
			}
		}()

		vars := mux.Vars(r)
		filename := vars["filename"]

		if len(filename) == 0 {
			status = http.StatusBadRequest
			errmsg = "no filename"
			return
		}

		logger.Info("filename is %s", filename)

		if file, err = gfs.Open(filename); nil != err {
			status = http.StatusNotFound
			errmsg = "file not exist"
			return
		}

		content = make([]byte, file.Size())
		if _, err = file.Read(content); nil != err {
			status = http.StatusInternalServerError
			errmsg = "error! while reading file"
			return
		}

		if err = file.Close(); nil != err {
			status = http.StatusInternalServerError
			errmsg = "i/o error"
			return
		}

		file_ext := filepath.Ext(filename)

		logger.Info("file extension is %s", file_ext)

		content_type := mime.TypeByExtension(file_ext)

		logger.Info("content type is %s", content_type)

		if len(content_type) > 0 {
			r.Header.Set("Content-Type", content_type)
		} else {
			r.Header.Set("Content-Type", "image/*")
		}

		width_str := r.URL.Query().Get("w")

		var (
			width        uint64
			is_width_set = false
		)
		if len(width_str) > 0 {

			if width, err = strconv.ParseUint(width_str, 10, 32); nil != err {
				status = http.StatusBadRequest
				return
			}

			is_width_set = true
		}

		height_str := r.URL.Query().Get("h")

		var (
			height        uint64
			is_height_set = false
		)
		if len(height_str) > 0 {

			if height, err = strconv.ParseUint(height_str, 10, 32); nil != err {
				status = http.StatusBadRequest
				return
			}

			is_height_set = true
		}

		logger.Info("width and height is [%d, %d], status [%t, %t]", width, height, is_width_set, is_height_set)

		if is_width_set || is_height_set {
			var (
				original_image image.Image
				new_image      image.Image
			)
			if original_image, _, err = image.Decode(bytes.NewReader(content)); nil != err {
				logger.Error("image decode error! %v", err)
				goto LABEL_IMAGE_HANDLE_FINISHED
			}

			new_image = resize.Resize(uint(width), uint(height), original_image, resize.Lanczos3)
			buf := new(bytes.Buffer)
			if err := jpeg.Encode(buf, new_image, nil); nil != err {
				logger.Error("image encode error! %v", err)
				goto LABEL_IMAGE_HANDLE_FINISHED
			}
			content = buf.Bytes()

			r.Header.Set("Content-Type", "image/jpeg")
		}

	LABEL_IMAGE_HANDLE_FINISHED:

		w.Write(content)
	})

	port := config.GetServPort()
	logger.Info("listen on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
