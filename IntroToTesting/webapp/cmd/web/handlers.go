package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
	"webapp/pkg/data"
)

var pathToTemplates = "./templates/"
var uploadPath = "./static/img"

type TemplateData struct {
	IP    string
	Data  map[string]any
	Error string
	Flash string
	User  data.User
}

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var td = make(map[string]any)

	if app.Session.Exists(r.Context(), "test") {
		msg := app.Session.GetString(r.Context(), "test")
		td["test"] = msg
	} else {
		app.Session.Put(r.Context(), "test", "Hit this page at "+time.Now().UTC().String())
	}

	_ = app.render(w, r, "home.page.gohtml", &TemplateData{Data: td})
}

func (app *application) Profile(w http.ResponseWriter, r *http.Request) {
	_ = app.render(w, r, "profile.page.gohtml", &TemplateData{})
}

func (app *application) render(w http.ResponseWriter, r *http.Request, t string, td *TemplateData) error {
	// Parse the template from disk.
	parsedTemplate, err := template.ParseFiles(path.Join(pathToTemplates, t), path.Join(pathToTemplates, "base.layout.gohtml"))
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return err
	}

	td.IP = app.ipFromContext(r.Context())

	td.Error = app.Session.PopString(r.Context(), "error")
	td.Flash = app.Session.PopString(r.Context(), "flash")

	if app.Session.Exists(r.Context(), "user") {
		td.User = app.Session.Get(r.Context(), "user").(data.User)
	}

	// Execute the template, passing it data, if any
	err = parsedTemplate.Execute(w, td)
	if err != nil {
		return err
	}

	return nil
}

func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	// Validate data
	form := NewForm(r.PostForm)
	form.Required("email", "password")

	if !form.Valid() {
		// Redirect to the login page with error message
		app.Session.Put(r.Context(), "error", "Invalid login credentials!")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	user, err := app.DB.GetUserByEmail(email)
	if err != nil {
		// Redirect to the login page with error message
		app.Session.Put(r.Context(), "error", "Invalid login!")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Authenticate the user
	if !app.authenticate(r, user, password) {
		app.Session.Put(r.Context(), "error", "Invalid login!")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// If not authenticated then redirect with error

	// Prevent fixation attack
	_ = app.Session.RenewToken(r.Context())

	// Store success message in the session

	// Redirect to some other page
	app.Session.Put(r.Context(), "flash", "Successfully logged in!")
	http.Redirect(w, r, "/user/profile", http.StatusSeeOther)
}

func (app *application) authenticate(r *http.Request, user *data.User, password string) bool {
	if valid, err := user.PasswordMatches(password); err != nil || !valid {
		return false
	}

	app.Session.Put(r.Context(), "user", user)
	return true
}

func (app *application) UploadProfilePic(w http.ResponseWriter, r *http.Request) {
	// Call a function that extracts a file from an upload (request)
	files, err := app.UploadFiles(r, uploadPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get a user from the session
	user := app.Session.Get(r.Context(), "user").(data.User)

	// Create a var of type data.UserImage
	var i = data.UserImage{
		UserID:   user.ID,
		FileName: files[0].OriginalFileName,
	}

	// Insert the user image into user_images
	_, err = app.DB.InsertUserImage(i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Refresh the sessional variable "user"
	updatedUser, err := app.DB.GetUser(user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	app.Session.Put(r.Context(), "user", updatedUser)

	// Redirect back to profile page
	http.Redirect(w, r, "/user/profile", http.StatusSeeOther)
}

type UploadedFile struct {
	OriginalFileName string
	FileSize         int64
}

func (app *application) UploadFiles(r *http.Request, uploadDir string) ([]*UploadedFile, error) {
	var uploadedFiles []*UploadedFile

	err := r.ParseMultipartForm(int64(1024 * 1024 * 5))
	if err != nil {
		return nil, fmt.Errorf("The uploaded file is too big, and must be less than %d bytes", 1024*1024*5)
	}

	for _, fHeaders := range r.MultipartForm.File {
		for _, hdr := range fHeaders {
			uploadedFiles, err = func(uploadedFiles []*UploadedFile) ([]*UploadedFile, error) {
				var uploadedFile UploadedFile
				infile, err := hdr.Open()
				if err != nil {
					return nil, err
				}
				defer infile.Close()

				uploadedFile.OriginalFileName = hdr.Filename

				var outFile *os.File
				defer outFile.Close()

				if outFile, err = os.Create(filepath.Join(uploadDir, uploadedFile.OriginalFileName)); nil != err {
					return nil, err
				} else {
					fileSize, err := io.Copy(outFile, infile)
					if err != nil {
						return nil, err
					}
					uploadedFile.FileSize = fileSize
				}

				uploadedFiles = append(uploadedFiles, &uploadedFile)

				return uploadedFiles, nil
			}(uploadedFiles)
			if err != nil {
				return uploadedFiles, err
			}
		}
	}

	return uploadedFiles, nil
}
