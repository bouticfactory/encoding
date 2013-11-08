// Copyright 2013 outofpluto.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package macroman implements the MacRomanEncoding data encoding
// as used in Adobe's PostScript and PDF document formats.
// https://www.adobe.com/devnet/pdf/pdf_reference.html

/*
		Source strings should be UTF-8

		Usage:

		func main() {
				src := []byte("Caractères spéciaux : ® et $")
				dst := make([]byte, len(src))
				macroman.Encode(dst, src)
				// use dst in a pdf generator
		}
*/
package macroman
