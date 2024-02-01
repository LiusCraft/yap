/*
 * Copyright (c) 2023 The GoPlus Authors (goplus.org). All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package yap

import (
	"html/template"
	"io/fs"
	"path/filepath"

	"github.com/goplus/yap/internal/templ"
)

// template delimiter, default is {{ }}
type Delims struct {
	Left  string
	Right string
}

// Template is the representation of a parsed template. The *parse.Tree
// field is exported only for use by html/template and should be treated
// as unexported by all other clients.
type Template struct {
	*template.Template
	delims Delims
}

// NewTemplate allocates a new, undefined template with the given name.
func NewTemplate(name string) Template {
	return Template{template.New(name), Delims{"{{", "}}"}}
}

func (t Template) Parse(text string) (ret Template, err error) {
	t.Template.Delims(t.delims.Left, t.delims.Right)
	ret.Template, err = t.Template.Parse(templ.Translate(text, t.delims.Left, t.delims.Right))
	return
}

func ParseFSFile(f fs.FS, file string, delims Delims) (t Template, err error) {
	b, err := fs.ReadFile(f, file)
	if err != nil {
		return
	}
	name := filepath.Base(file)
	t = NewTemplate(name)
	t.delims = delims
	return t.Parse(string(b))
}
