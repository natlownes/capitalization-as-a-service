package app

import (
	"fmt"
	"net/http"
	"strings"
)

func CapitalizeHandler(w http.ResponseWriter, r *http.Request) {
	input := r.FormValue("arg")
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, strings.ToUpper(input))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	body := `
  <html>
    <head>
      <title>
        Capitalization Microsevice Documentation
      </title>
      <style type="text/css">
        body {
          margin-left:   15%;
          margin-right:  15%;
          padding-bottom: 20px;
        }

        a {
          color:            black;
          text-decoration:  none;
        }

        a:hover {
          color:            FF4500;
          text-decoration:  none;
        }

        hr {
          border: 1px solid black;
        }

        blockquote {
          padding:      20px;
          background:   lightgrey;
          font-family:  monospace;
        }
      </style>
    </head>
    <body>
      <h1><strong>CAPITALIZATION MICROSERVICE</strong></h1>

      <p>
        Everyone knows that microservices are resilent because they exist and
        words are said about them.  Trying to manage capitalization across
        different architectures, languages, and character encodings in a
        performant manner has been a massive stumbling block for modern
        distributed systems and distributed teams.
      </p>

      <p>
        We are excited to announce a powerful, flexible, heavily tested, and
        completely free capitalization microservice for use in any and all of
        your distributed systems.
      </p>

      <hr />

      <h3>
        GET <a href="/capitalize?arg=h" target="_blank">/capitalize?arg=h</a>
      </h3>

      <p>
        Responds with the capitalized version of the letter you request:
      </p>

      <blockquote>
        H
      </blockquote>

      <hr />

      <h3>
        GET <a href="/capitalize?arg=hello%2C%20how%20are%20you%3F" target="_blank">
          /capitalize?arg=hello, how are you?
        </a>
      </h3>

      <p>
        You are not restricted to capitalizing one letter at a time (though you
        may find it more performant to do so in a distributed system) - you can
        also request an entire string to be capitalized.
      </p>

      <blockquote>
        HELLO, HOW ARE YOU?
      </blockquote>

    </body>
  </html>
  `
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprint(w, body)
}

func CapitalizationService() {
	http.HandleFunc("/capitalize", CapitalizeHandler)
	http.HandleFunc("/", HomeHandler)
}
