// Copyright (c) 2015 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package utils

import (
	"bytes"
	"io"
	"io/ioutil"

	l4g "github.com/alecthomas/log4go"
)

// DebugReader logs the content of the io.Reader and returns a new io.Reader
// with the same content as the received io.Reader.
// If you pass reader by reference, it won't be re-created unless the loglevel
// includes Debug.
// If an error is returned, the reader is consumed an cannot be read again.
func DebugReader(reader io.Reader, message string) (io.Reader, error) {
	var err error
	l4g.Debug(func() string {
		content, err := ioutil.ReadAll(reader)
		if err != nil {
			return ""
		}

		reader = bytes.NewReader(content)

		return message + string(content)
	})

	return reader, err
}
