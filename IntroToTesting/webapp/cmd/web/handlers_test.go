package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"image"
	"image/png"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path"
	"strings"
	"sync"
	"testing"
	"webapp/pkg/data"
)

func Test_application_handlers(t *testing.T) {
	var tests = []struct {
		name                    string
		url                     string
		expectedStatusCode      int
		expectedURL             string
		expectedFirstStatusCode int
	}{
		{"home", "/", http.StatusOK, "/", http.StatusOK},
		{"404", "/fish", http.StatusNotFound, "/fish", http.StatusNotFound},
		{"profile", "/user/profile", http.StatusOK, "/", http.StatusTemporaryRedirect},
	}

	routes := app.routes()

	// Create a test server
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: tr,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	// Range through test data
	for _, e := range tests {
		resp, err := ts.Client().Get(ts.URL + e.url)
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}

		if resp.StatusCode != e.expectedStatusCode {
			t.Errorf("%s: expected status: %d, but got: %d", e.name, e.expectedStatusCode, resp.StatusCode)
		}

		if resp.Request.URL.Path != e.expectedURL {
			t.Errorf("%s: expected final url of %s, but got %s", e.name, e.expectedURL, resp.Request.URL.Path)
		}

		resp2, _ := client.Get(ts.URL + e.url)
		if resp2.StatusCode != e.expectedFirstStatusCode {
			t.Errorf("%s: expected first returned status code to be %d, but got %d", e.name, e.expectedFirstStatusCode, resp2.StatusCode)
		}
	}
}

func TestAppHome(t *testing.T) {
	tests := []struct {
		name         string
		putInSession string
		expectedHTML string
	}{
		{"first visit", "", "<small>From Session:"},
		{"second visit", "hello, world!", "<small>From Session: hello, world!"},
	}

	for _, e := range tests {
		// Create a request
		req, _ := http.NewRequest("GET", "/", nil)

		req = addContextAndSessionToRequest(req, app)
		_ = app.Session.Destroy(req.Context())

		if e.putInSession != "" {
			app.Session.Put(req.Context(), "test", e.putInSession)
		}

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(app.Home)

		handler.ServeHTTP(rr, req)

		// Check status code
		if rr.Code != http.StatusOK {
			t.Errorf("TestAppHome returned wrong status code; expected 200, but got: %d", rr.Code)
		}

		body, _ := io.ReadAll(rr.Body)
		if !strings.Contains(string(body), e.expectedHTML) {
			t.Errorf("%s: did not find %s in response body", e.name, e.expectedHTML)
		}
	}
}

func TestApp_renderWithBadTemplate(t *testing.T) {
	// Set template path to a location with a bad template
	pathToTemplates = "./testdata/"

	req, _ := http.NewRequest("GET", "/", nil)
	req = addContextAndSessionToRequest(req, app)
	rr := httptest.NewRecorder()

	err := app.render(rr, req, "bad.page.gohtml", &TemplateData{})
	if err == nil {
		t.Error("Expected error from bad template, but did not get one")
	}

	pathToTemplates = "./../../templates/"
}

func getCtx(req *http.Request) context.Context {
	ctx := context.WithValue(req.Context(), contextUserKey, "unknown")
	return ctx
}

func addContextAndSessionToRequest(req *http.Request, app application) *http.Request {
	req = req.WithContext(getCtx(req))

	ctx, _ := app.Session.Load(req.Context(), req.Header.Get("X-Session"))

	return req.WithContext(ctx)
}

func Test_app_Login(t *testing.T) {
	var tests = []struct {
		name               string
		postedData         url.Values
		expectedStatusCode int
		expectedLoc        string
	}{
		{
			name: "valid login",
			postedData: url.Values{
				"email":    {"admin@example.com"},
				"password": {"secret"},
			},
			expectedStatusCode: http.StatusSeeOther,
			expectedLoc:        "/user/profile",
		},
		{
			name: "missing form data",
			postedData: url.Values{
				"email":    {""},
				"password": {""},
			},
			expectedStatusCode: http.StatusSeeOther,
			expectedLoc:        "/",
		},
		{
			name: "user not found",
			postedData: url.Values{
				"email":    {"you@error.com"},
				"password": {"password"},
			},
			expectedStatusCode: http.StatusSeeOther,
			expectedLoc:        "/",
		},
		{
			name: "bad credentials",
			postedData: url.Values{
				"email":    {"admin@example.com"},
				"password": {"secretttt"},
			},
			expectedStatusCode: http.StatusSeeOther,
			expectedLoc:        "/",
		},
	}

	for _, e := range tests {
		req, _ := http.NewRequest("POST", "/login", strings.NewReader(e.postedData.Encode()))
		req = addContextAndSessionToRequest(req, app)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(app.Login)
		handler.ServeHTTP(rr, req)

		if rr.Code != e.expectedStatusCode {
			t.Errorf("%s: returned wrong status code; expected %d, but got %d", e.name, e.expectedStatusCode, rr.Code)
		}

		actualLoc, err := rr.Result().Location()
		if err == nil {
			if actualLoc.String() != e.expectedLoc {
				t.Errorf("%s: expected location %s, but got %s", e.name, e.expectedLoc, actualLoc.String())
			}
		} else {
			t.Errorf("%s: no location header set", e.name)
		}
	}
}

func Test_app_UploadFiles(t *testing.T) {
	// Set up pipes
	pr, pw := io.Pipe()

	// Create a new writer, of type *io.Writer
	writer := multipart.NewWriter(pw)

	// Create a waitgroup, and add 1 to it
	wg := &sync.WaitGroup{}
	wg.Add(1)

	// Simulate uploading a file using a goroutine and our writer
	go simulatePNGUpload("./testdata/img.png", writer, t, wg)

	// Read from the pipe which receives data
	request := httptest.NewRequest("POST", "/", pr)
	request.Header.Add("Content-Type", writer.FormDataContentType())

	// Call app.UploadFiles
	uploadedFiles, err := app.UploadFiles(request, "./testdata/uploads/")
	if err != nil {
		t.Error(err)
	}

	// Preform our tests
	if _, err := os.Stat(fmt.Sprintf("./testdata/uploads/%s", uploadedFiles[0].OriginalFileName)); os.IsNotExist(err) {
		t.Errorf("Expected file to exist: %s", err.Error())
	}

	// Clean up
	_ = os.Remove(fmt.Sprintf("./testdata/uploads/%s", uploadedFiles[0].OriginalFileName))

	wg.Wait()
}

func simulatePNGUpload(fileToUpload string, writer *multipart.Writer, t *testing.T, wg *sync.WaitGroup) {
	defer writer.Close()
	defer wg.Done()

	// Create the form data filed 'file' with value being filename
	part, err := writer.CreateFormFile("file", path.Base(fileToUpload))
	if err != nil {
		t.Error(err)
	}

	// Open the actual file
	f, err := os.Open(fileToUpload)
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	// Decode the image
	img, _, err := image.Decode(f)
	if err != nil {
		t.Error("Error decoding image:", err)
	}

	// Write the PNG to our io.Writer
	err = png.Encode(part, img)
	if err != nil {
		t.Error(err)
	}
}

func Test_app_UploadProfilePic(t *testing.T) {
	uploadPath = "./testdata/uploads"
	filePath := "./testdata/img.png"

	// Specify a field name for the form
	fieldName := "file"

	// Create a bytes.Buffer to act as the request body
	body := new(bytes.Buffer)

	// Create a new writer
	mw := multipart.NewWriter(body)

	file, err := os.Open(filePath)
	if err != nil {
		t.Fatal(err)
	}

	w, err := mw.CreateFormFile(fieldName, filePath)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := io.Copy(w, file); err != nil {
		t.Fatal(err)
	}

	mw.Close()
	req := httptest.NewRequest(http.MethodPost, "/upload", body)
	req = addContextAndSessionToRequest(req, app)
	app.Session.Put(req.Context(), "user", data.User{ID: 1})
	req.Header.Add("Content-Type", mw.FormDataContentType())

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(app.UploadProfilePic)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusSeeOther {
		t.Error("Wrong status code")
	}

	_ = os.Remove("./testdata/uploads/img.png")
}
