package legacy

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

type Context struct {
	context.Context
}

func NewContextWithLogger(ctx context.Context) Context {
	return Context{ctx}
}

func (c Context) Debugf(format string, args ...interface{}) {
	log.Debugf(c, format, args...)
}
func (c Context) Infof(format string, args ...interface{}) {
	log.Infof(c, format, args...)
}
func (c Context) Warningf(format string, args ...interface{}) {
	log.Warningf(c, format, args...)
}
func (c Context) Errorf(format string, args ...interface{}) {
	log.Errorf(c, format, args...)
}
func (c Context) Criticalf(format string, args ...interface{}) {
	log.Criticalf(c, format, args...)
}
