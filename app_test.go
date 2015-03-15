package app_test

import (
	"bytes"
	. "github.com/natlownes/capitalization_as_a_service"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCapitalizationMicroservice(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CapitalizationMicroservice Suite")
}

var _ = Describe("CapitalizationService", func() {
	var ()

	It("should respond with capital letters", func() {
		url := "http://example.com/capitalize?arg=hello how are you?"
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			Fail(err.Error())
		}

		w := httptest.NewRecorder()
		CapitalizeHandler(w, req)

		expected := "HELLO HOW ARE YOU?"
		Expect(w.Body.String()).To(Equal(expected))
	})

	It("should respond with empty string if no query string", func() {
		url := "http://example.com/capitalize"
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			Fail(err.Error())
		}

		w := httptest.NewRecorder()
		CapitalizeHandler(w, req)

		expected := ""
		Expect(w.Body.String()).To(Equal(expected))
	})

	It("should respond with empty string if empty arg param", func() {
		url := "http://example.com/capitalize?arg="
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			Fail(err.Error())
		}

		w := httptest.NewRecorder()
		CapitalizeHandler(w, req)

		expected := ""
		Expect(w.Body.String()).To(Equal(expected))
	})

	It("should respond with HTTP 200", func() {
		url := "http://example.com/capitalize?arg=dogs are huge"
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			Fail(err.Error())
		}

		w := httptest.NewRecorder()
		CapitalizeHandler(w, req)

		Expect(w.Code).To(Equal(200))
	})

	It("should handle an HTTP POST", func() {
		url := "http://example.com/capitalize"
		body := []byte(`dogs are very, very, very huge!!!`)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
		if err != nil {
			Fail(err.Error())
		}

		w := httptest.NewRecorder()
		CapitalizeHandler(w, req)
		expected := "DOGS ARE VERY, VERY, VERY HUGE!!!"

		Expect(w.Body.String()).To(Equal(expected))
	})

	It("should HTTP 400 an HTTP POST with empty body", func() {
		url := "http://example.com/capitalize"
		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			Fail(err.Error())
		}

		w := httptest.NewRecorder()
		CapitalizeHandler(w, req)
		expected := "HTTP Body required\n"

		Expect(w.Code).To(Equal(400))
		Expect(w.Body.String()).To(Equal(expected))
	})

	It("should have a content-type of text/plain", func() {
		url := "http://example.com/capitalize?arg=dogs are huge"
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			Fail(err.Error())
		}

		w := httptest.NewRecorder()
		CapitalizeHandler(w, req)
		expected := "text/plain; charset=utf-8"

		Expect(w.HeaderMap.Get("content-type")).To(Equal(expected))
	})
})
