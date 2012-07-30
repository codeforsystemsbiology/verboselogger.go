/*
   Copyright (C) 2003-2011 Institute for Systems Biology
                           Seattle, Washington, USA.

   This library is free software; you can redistribute it and/or
   modify it under the terms of the GNU Lesser General Public
   License as published by the Free Software Foundation; either
   version 2.1 of the License, or (at your option) any later version.

   This library is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
   Lesser General Public License for more details.

   You should have received a copy of the GNU Lesser General Public
   License along with this library; if not, write to the Free Software
   Foundation, Inc., 59 Temple Place, Suite 330, Boston, MA 02111-1307  USA

*/
package log4go

import (
	"fmt"
	"io"
	"log"
	"os"
)

// constructs a new verbose logger that wraps a standard out logger
func NewVerboseLogger(verbose bool, out io.Writer, prefix string) *VerboseLogger {
	if out == nil {
		out = os.Stdout
	}
	flags := log.LstdFlags | log.Lshortfile
	return &VerboseLogger{log.New(out, prefix, flags), verbose}
}

// extends log.Logger to add functions for verbose logging control and warning
type VerboseLogger struct {
	*log.Logger
	Verbose bool
}

func (this *VerboseLogger) Debug(format string, a ...interface{}) {
	if this.Verbose {
		this.Output(2, fmt.Sprintf(format, a...))
	}
}
func (this *VerboseLogger) Warn(err error) {
	this.Output(2, fmt.Sprintf("[WARN] %v", err))
}
