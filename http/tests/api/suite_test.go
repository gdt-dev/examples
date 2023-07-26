// Use and distribution licensed under the Apache license version 2.
//
// See the COPYING file in the root project directory for full text.

package api_test

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/gdt-dev/gdt"
	gdttypes "github.com/gdt-dev/gdt/types"
	gdthttp "github.com/gdt-dev/http"
	"github.com/stretchr/testify/require"

	"github.com/gdt-dev/examples/http/api"
)

const (
	dataFilePath = "testdata/fixtures.json"
)

func currentDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

type dataset struct {
	Authors    interface{}
	Publishers interface{}
	Books      []*api.Book
}

func data() *dataset {
	f, err := os.Open(dataFilePath)
	if err != nil {
		panic(err)
	}
	data := &dataset{}
	if err = json.NewDecoder(f).Decode(&data); err != nil {
		panic(err)
	}
	return data
}

func dataFixture() gdttypes.Fixture {
	f, err := os.Open(dataFilePath)
	if err != nil {
		panic(err)
	}
	f.Seek(0, io.SeekStart)
	fix, err := gdt.NewJSONFixture(f)
	if err != nil {
		panic(err)
	}
	return fix
}

func TestHTTP(t *testing.T) {
	require := require.New(t)

	// Register an HTTP server fixture that spins up the API service on a
	// random port on localhost
	logger := log.New(os.Stdout, "books_api_http: ", log.LstdFlags)
	c := api.NewControllerWithBooks(logger, data().Books)

	ctx := gdt.NewContext()
	apiFixture := gdthttp.NewServerFixture(c.Router(), false /* useTLS */)
	ctx = gdt.RegisterFixture(ctx, "books_api", apiFixture)
	ctx = gdt.RegisterFixture(ctx, "books_data", dataFixture())

	// Construct a new gdt.TestSuite from the directory of this file
	s, err := gdt.From(currentDir())
	require.Nil(err)

	s.Run(ctx, t)
}

func TestHTTPS(t *testing.T) {
	require := require.New(t)

	// Register an HTTPS server fixture that spins up the API service on a
	// random port on localhost and a well-known cert for localhost/127.0.0.1
	logger := log.New(os.Stdout, "books_api_https: ", log.LstdFlags)
	c := api.NewControllerWithBooks(logger, data().Books)

	ctx := gdt.NewContext()
	apiFixture := gdthttp.NewServerFixture(c.Router(), true /* useTLS */)
	ctx = gdt.RegisterFixture(ctx, "books_api", apiFixture)
	ctx = gdt.RegisterFixture(ctx, "books_data", dataFixture())

	// Construct a new gdt.TestSuite from the directory of this file
	s, err := gdt.From(currentDir())
	require.Nil(err)

	s.Run(ctx, t)
}
