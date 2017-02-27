// GO SOURCE FILE MACHINE GENERATED BY RAGEL; DO NOT EDIT

// Copyright 2014 The Cayley Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package nquads

import (
	"fmt"
	"unicode"

	"github.com/codelingo/cayley/quad"
)

%%{
	machine raw;

	action Escape {
        isEscaped = true
    }

    action StartSubject {
        subject = p
    }

    action StartPredicate {
        predicate = p
    }

    action StartObject {
        object = p
    }

    action StartLabel {
        label = p
    }

    action SetSubject {
        if subject < 0 {
            panic("unexpected parser state: subject start not set")
        }
        q.Subject = unEscapeRaw(data[subject:p], isEscaped)
        isEscaped = false
    }

    action SetPredicate {
        if predicate < 0 {
            panic("unexpected parser state: predicate start not set")
        }
        q.Predicate = unEscapeRaw(data[predicate:p], isEscaped)
        isEscaped = false
    }

    action SetObject {
        if object < 0 {
            panic("unexpected parser state: object start not set")
        }
        q.Object = unEscapeRaw(data[object:p], isEscaped)
        isEscaped = false
    }

    action SetLabel {
        if label < 0 {
            panic("unexpected parser state: label start not set")
        }
        q.Label = unEscapeRaw(data[label:p], isEscaped)
        isEscaped = false
    }

    action Return {
        return q, nil
    }

    action Comment {
    }

    action Error {
        if p < len(data) {
            if r := data[p]; r < unicode.MaxASCII {
                return q, fmt.Errorf("%v: unexpected rune %q at %d", quad.ErrInvalid, data[p], p)
            } else {
                return q, fmt.Errorf("%v: unexpected rune %q (\\u%04x) at %d", quad.ErrInvalid, data[p], data[p], p)
            }
        }
        return q, quad.ErrIncomplete
    }

	include nquads "nquads.rl";

    literal                 = STRING_LITERAL_QUOTE ('^^' IRIREF | LANGTAG)? ;

    subject                 = IRIREF | BLANK_NODE_LABEL ;
    predicate               = IRIREF ;
    object                  = IRIREF | BLANK_NODE_LABEL | literal ;
    graphLabel              = IRIREF | BLANK_NODE_LABEL ;

    statement := (
        whitespace*  subject    >StartSubject   %SetSubject
        whitespace*  predicate  >StartPredicate %SetPredicate
        whitespace*  object     >StartObject    %SetObject
        (whitespace* graphLabel >StartLabel     %SetLabel)?
        whitespace*  '.' whitespace* ('#' any*)? >Comment
     ) %Return @!Error ;

	write data;
}%%

// ParseRaw returns a valid quad.Quad or a non-nil error. ParseRaw does
// handle comments except where the comment placement does not prevent
// a complete valid quad.Quad from being defined.
func ParseRaw(statement string) (quad.Quad, error) {
	data := []rune(statement)

	var (
		cs, p int
		pe    = len(data)
		eof   = pe

		subject   = -1
		predicate = -1
		object    = -1
		label     = -1

		isEscaped bool

		q quad.Quad
	)

	%%write init;

	%%write exec;

	return quad.Quad{}, quad.ErrInvalid
}
