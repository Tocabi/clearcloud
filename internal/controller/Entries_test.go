package controller

import (
	"clearcloud/internal/model"
	"clearcloud/internal/service"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"testing"
)

func TestListEntries(t *testing.T) {
	tempDir := t.TempDir()
	subDir := filepath.Join(tempDir, "sub")
	_ = os.WriteFile(filepath.Join(tempDir, "root.txt"), []byte("Root file\n"), 0777)
	_ = os.Mkdir(subDir, 0777)
	_ = os.WriteFile(filepath.Join(subDir, "foo.txt"), []byte("I am foo\n"), 0777)
	_ = os.WriteFile(filepath.Join(subDir, "bar.txt"), []byte("I am bar\n"), 0777)

	lib := model.Library{
		Type:       model.TypeGeneric,
		Name:       "Test Library",
		RootFolder: tempDir,
	}
	testApp.DB.Create(&lib)

	t.Run("it lists files when querying a parent", func(t *testing.T) {
		req := adminRequest("GET", fmt.Sprintf("/api/libraries/%d/entries?parent=%s", lib.ID, url.QueryEscape("/sub")), nil)
		recorder := httptest.NewRecorder()
		testApp.Router.ServeHTTP(recorder, req)

		assertStatus(t, recorder, 200)
		var body []service.FileInfo
		assertJsonUnmarshal(t, recorder, &body)
		if len(body) != 2 {
			t.Errorf("Expected 2 files, but got %d", len(body))
		}
		if body[0].Name != "bar.txt" || body[1].Name != "foo.txt" {
			t.Errorf("Expected bar.txt and foo.txt, but got: %s and %s", body[0].Name, body[1].Name)
		}
	})

	t.Run("it returns a single entry when querying a path", func(t *testing.T) {
		req := adminRequest("GET", fmt.Sprintf("/api/libraries/%d/entries?path=root.txt", lib.ID), nil)
		recorder := httptest.NewRecorder()
		testApp.Router.ServeHTTP(recorder, req)

		assertStatus(t, recorder, 200)
		var body []service.FileInfo
		assertJsonUnmarshal(t, recorder, &body)
		if len(body) != 1 {
			t.Errorf("Expected 1 file, but got %d", len(body))
		}
		if body[0].Name != "root.txt" {
			t.Errorf("Expected root.txt, but got: %s", body[0].Name)
		}
	})

	t.Run("it returns an empty list when querying a path that does not exist", func(t *testing.T) {
		req := adminRequest("GET", fmt.Sprintf("/api/libraries/%d/entries?path=%s", lib.ID, url.QueryEscape("/no/file/here")), nil)
		recorder := httptest.NewRecorder()
		testApp.Router.ServeHTTP(recorder, req)

		assertStatus(t, recorder, 200)
		var body []interface{}
		assertJsonUnmarshal(t, recorder, &body)
		if len(body) > 0 {
			t.Errorf("Expected an empty list but got: %s", recorder.Body.String())
		}
	})

	t.Run("it returns an empty list when querying the root path", func(t *testing.T) {
		req := adminRequest("GET", fmt.Sprintf("/api/libraries/%d/entries?path=", lib.ID), nil)
		recorder := httptest.NewRecorder()
		testApp.Router.ServeHTTP(recorder, req)

		assertStatus(t, recorder, 200)
		var body []interface{}
		assertJsonUnmarshal(t, recorder, &body)
		if len(body) > 0 {
			t.Errorf("Expected an empty list but got: %s", recorder.Body.String())
		}
	})

	t.Run("it requires authentication", func(t *testing.T) {
		req, _ := http.NewRequest("GET", fmt.Sprintf("/api/libraries/%d/entries?path=root.txt", lib.ID), nil)
		recorder := httptest.NewRecorder()
		testApp.Router.ServeHTTP(
			recorder,
			req,
		)
		assertStatus(t, recorder, 401)
	})
}

func TestDownloadEntry(t *testing.T) {
	// Prep filesystem
	tempDir := t.TempDir()
	_ = os.WriteFile(filepath.Join(tempDir, "gimme.txt"), []byte("Hello World!\n"), 0777)

	lib := model.Library{
		Type:       model.TypeGeneric,
		Name:       "Test Library",
		RootFolder: tempDir,
	}
	testApp.DB.Create(&lib)

	// Go!
	t.Run("it downloads a file", func(t *testing.T) {
		req := adminRequest("GET", fmt.Sprintf("/api/libraries/%d/entries/download?path=gimme.txt", lib.ID), nil)
		recorder := httptest.NewRecorder()
		testApp.Router.ServeHTTP(
			recorder,
			req,
		)

		contentType := recorder.Header().Get("Content-Type")
		if contentType != "text/plain" {
			t.Errorf("Expected text/plain but got: %s", contentType)
		}
		disposition := recorder.Header().Get("Content-Disposition")
		if disposition != "attachment; filename=\"gimme.txt\"" {
			t.Errorf("Expected [attachment; filename=\"gimme.txt\"] but got: %s", disposition)
		}
		body := recorder.Body.String()
		if body != "Hello World!\n" {
			t.Errorf("Expected [Hello World!] but got: %s", body)
		}
	})

	t.Run("it returns 404 if file not found", func(t *testing.T) {
		req := adminRequest("GET", fmt.Sprintf("/api/libraries/%d/entries/download?path=nope.txt", lib.ID), nil)
		recorder := httptest.NewRecorder()
		testApp.Router.ServeHTTP(
			recorder,
			req,
		)
		assertStatus(t, recorder, 404)
	})

	t.Run("it returns 404 if no access to library", func(t *testing.T) {
		req := noAccessUserRequest("GET", fmt.Sprintf("/api/libraries/%d/entries/download?path=gimme.txt", lib.ID), nil)
		recorder := httptest.NewRecorder()
		testApp.Router.ServeHTTP(
			recorder,
			req,
		)
		assertStatus(t, recorder, 404)
	})

	t.Run("it returns 400 if no path provided", func(t *testing.T) {
		req := adminRequest("GET", fmt.Sprintf("/api/libraries/%d/entries/download", lib.ID), nil)
		recorder := httptest.NewRecorder()
		testApp.Router.ServeHTTP(
			recorder,
			req,
		)
		assertStatus(t, recorder, 400)
	})

	t.Run("it requires authentication", func(t *testing.T) {
		req, _ := http.NewRequest("GET", fmt.Sprintf("/api/libraries/%d/entries/download?path=not_allowed", lib.ID), nil)
		recorder := httptest.NewRecorder()
		testApp.Router.ServeHTTP(
			recorder,
			req,
		)
		assertStatus(t, recorder, 401)
	})
}

func TestDeleteEntry(t *testing.T) {
	tempDir := t.TempDir()
	rootFile := filepath.Join(tempDir, "root.txt")
	_ = os.WriteFile(rootFile, []byte("I am root\n"), 0777)

	lib := model.Library{
		Type:       model.TypeGeneric,
		Name:       "Test Library",
		RootFolder: tempDir,
	}
	testApp.DB.Create(&lib)

	t.Run("it requires authentication", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", fmt.Sprintf("/api/libraries/%d/entries?path=root.txt", lib.ID), nil)
		recorder := httptest.NewRecorder()
		testApp.Router.ServeHTTP(
			recorder,
			req,
		)

		assertStatus(t, recorder, 401)
		if _, err := os.Stat(rootFile); err != nil {
			t.Errorf("File does not exist: %s", err)
		}
	})

	t.Run("it deletes a file", func(t *testing.T) {
		deleteFile := filepath.Join(tempDir, "delete_me.txt")
		_ = os.WriteFile(deleteFile, []byte("Delete me plz\n"), 0777)

		req := adminRequest("DELETE", fmt.Sprintf("/api/libraries/%d/entries?path=delete_me.txt", lib.ID), nil)
		recorder := httptest.NewRecorder()
		testApp.Router.ServeHTTP(
			recorder,
			req,
		)

		assertStatus(t, recorder, 204)
		if _, err := os.Stat(deleteFile); !os.IsNotExist(err) {
			t.Errorf("File still exists")
		}
	})

	t.Run("it deletes a folder", func(t *testing.T) {
		dir := filepath.Join(tempDir, "dir")
		_ = os.Mkdir(dir, 0777)
		_ = os.WriteFile(filepath.Join(tempDir, "delete_me.txt"), []byte("File in dir\n"), 0777)

		req := adminRequest("DELETE", fmt.Sprintf("/api/libraries/%d/entries?path=dir", lib.ID), nil)
		recorder := httptest.NewRecorder()
		testApp.Router.ServeHTTP(
			recorder,
			req,
		)

		assertStatus(t, recorder, 204)
		if _, err := os.Stat(dir); !os.IsNotExist(err) {
			t.Errorf("Folder still exists")
		}
	})

	t.Run("it does not error if the path does not exist", func(t *testing.T) {
		req := adminRequest("DELETE", fmt.Sprintf("/api/libraries/%d/entries?path=nope", lib.ID), nil)
		recorder := httptest.NewRecorder()
		testApp.Router.ServeHTTP(
			recorder,
			req,
		)

		assertStatus(t, recorder, 204)
	})
}
