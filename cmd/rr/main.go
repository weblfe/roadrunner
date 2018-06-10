// MIT License
//
// Copyright (c) 2018 SpiralScout
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	rr "github.com/spiral/roadrunner/cmd/rr/cmd"

	// services (plugins)
	"github.com/spiral/roadrunner/http"
	"github.com/spiral/roadrunner/rpc"
	"github.com/spiral/roadrunner/static"

	// cli plugins
	_ "github.com/spiral/roadrunner/cmd/rr/http"
	"github.com/spiral/roadrunner/debug"
)

func main() {
	// provides ability to make local connection to services
	rr.Container.Register(rpc.Name, &rpc.Service{})

	// http serving
	rr.Container.Register(http.Name, &http.Service{})

	// serving static files
	rr.Container.Register(static.Name, &static.Service{})

	// provides additional verbosity
	rr.Container.Register(debug.Name, &debug.Service{Logger: rr.Logger})

	// you can register additional commands using cmd.CLI
	rr.Execute()
}
