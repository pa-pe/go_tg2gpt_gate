package logger

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.elastic.co/ecslogrus"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
	"upserv/config"
)

var accessLogger = logrus.New()

func initAccessLog(filePath string, env string) {
	// setting up output to file
	accessLogger.SetOutput(ioutil.Discard) // Default discard all output
	if filePath != "" {
		f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			LaunchLog("error opening file:")
			LaunchLog(err.Error())
		} else {
			accessLogger.SetLevel(logrus.TraceLevel)
			accessLogger.SetFormatter(&ecslogrus.Formatter{})
			accessLogger.WithField("Env", env)
			accessLogger.SetOutput(f)
		}
	}
}

func NewRequestLog(w http.ResponseWriter, r *http.Request) *requestLog {
	ls := &requestLog{
		OriginalWriter: w,
		Request:        r,
		fields:         make(map[string]interface{}),
		startTime:      time.Now(),
	}
	return ls
}

type requestLog struct {
	OriginalWriter http.ResponseWriter
	Request        *http.Request
	fields         logrus.Fields
	startTime      time.Time
}

func (l *requestLog) GetResponseWriter() ResponseLogWriter {
	return ResponseLogWriter{
		original: l.OriginalWriter,
		logger:   l,
	}
}

func (l *requestLog) GetRequest() *http.Request {
	return l.Request
}

func (l *requestLog) AddField(key string, value interface{}) {
	l.fields[key] = value
}

func (l *requestLog) Commit() {
	_ = l.Request.ParseForm()
	reqT := logRequest{
		l.Request.Header,
		l.Request.RemoteAddr,
		l.Request.Method,
		l.Request.RequestURI,
		l.Request.Form.Encode(),
	}
	l.AddField("Request", reqT)
	l.AddField("RequestID", l.Request.Header.Get(config.RequestIdKey))
	l.Request.WithContext(context.WithValue(l.Request.Context(), "RequestID", l.Request.Header.Get(config.RequestIdKey)))
	l.AddField("Response.Headers", l.OriginalWriter.Header())
	if l.fields["Response.Body"] == "" {
		l.AddField("Response.Body", "")
	}
	if l.Request.Method == "OPTIONS" {
		return
	}
	accessLogger.WithFields(l.fields).Trace("Trace caller")
}

type logRequest struct {
	Headers    interface{}
	RemoteAddr string
	Method     string
	RequestURI string
	Form       string
}

// ResponseLogWriter Type implements interface @http.ResponseWriter also added logs of response
// all what this implementation do is add logs before wright output
type ResponseLogWriter struct {
	original   http.ResponseWriter
	logger     *requestLog
	statusCode int
}

func (w ResponseLogWriter) Header() http.Header {
	return w.original.Header()
}

func (w ResponseLogWriter) Write(b []byte) (int, error) {
	var sb strings.Builder
	responseLength := config.GetInt("logger", "response_length")
	if responseLength > 0 && len(b) > responseLength {
		sb.Write(b[:responseLength])
	} else {
		sb.Write(b)
	}
	w.logger.AddField("ResponseBody", sb.String())
	return w.original.Write(b)
}

func (w ResponseLogWriter) WriteHeader(statusCode int) {
	d := time.Since(w.logger.startTime)
	w.logger.AddField("ResponseCode", statusCode)
	w.logger.AddField("Response.Time", d.Milliseconds())
	w.Header().Set("Execute-time", fmt.Sprintf("%dms", d.Milliseconds()))
	w.original.WriteHeader(statusCode)
}
