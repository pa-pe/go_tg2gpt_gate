package response

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"
	"upserv/src/apperror"
)

type Cookie struct {
	Name    string
	Value   string
	Expires time.Time
}
type Cookies struct {
	list map[string]*Cookie
}

func (c *Cookies) Set(model *Cookie) {
	if len(model.Name) == 0 {
		return
	}
	if c.list == nil {
		c.list = make(map[string]*Cookie)
	}
	c.list[model.Name] = model
}

func (c *Cookies) GetAll() []*Cookie {
	var res []*Cookie
	for _, cc := range c.list {
		res = append(res, cc)
	}
	return res
}

type NoFilterList struct {
	List []interface{} `json:"List"`
}

type List struct {
	Limit  int           `json:"Limit"`
	Offset int           `json:"Offset"`
	Total  int64         `json:"Total"`
	List   []interface{} `json:"List"`
}

type errorResp struct {
	// Error message to display
	ErrorMsg string `json:"error"`
	// Internal apperror code
	Code int32 `json:"code"`
	// id to determinate what exacly was wrong by searching in logs.
	RequestId string `json:"request_id"`
}

func Error(e *apperror.IError, w http.ResponseWriter) {
	js, _ := json.Marshal(
		&errorResp{
			e.Msg,
			e.Code,
			e.RequestID,
		},
	)
	w.WriteHeader(e.HttpCode)
	_, _ = w.Write(js)
}

func Success(obj interface{}, w http.ResponseWriter) {
	if obj == nil {
		w.WriteHeader(http.StatusAccepted)
		return
	}
	writeCookiesFromObj(obj, w)

	js, err := json.Marshal(obj)
	if err != nil {
		e := apperror.InternalError(context.Background()).WithMsg(err.Error())
		Error(e, w)
		return
	}
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(js)
}

func writeCookiesFromObj(obj interface{}, w http.ResponseWriter) {
	element := reflect.ValueOf(obj).Elem()
	cookiesField := element.FieldByName("Cookies")
	if !cookiesField.IsValid() {
		return
	}
	cookiesStruct := cookiesField.Interface().(Cookies)
	cookiesList := cookiesStruct.GetAll()
	for _, val := range cookiesList {
		http.SetCookie(w, &http.Cookie{
			Name:    val.Name,
			Value:   val.Value,
			Expires: val.Expires,
		})
	}
}

func GetExpFields(m interface{}) []string {
	fields := make([]string, 0)
	val := reflect.ValueOf(m)
	for i := 0; i < val.Type().NumField(); i++ {
		t := val.Type().Field(i)
		if exportTag := t.Tag.Get("export"); exportTag != "" {
			if expRule := strings.Split(exportTag, ","); (len(expRule) > 1 && expRule[0] == "+") || exportTag == "+" {
				fields = append(fields, t.Name)
			}
		}
	}
	return fields
}

func ToExpSlice(m interface{}) []string {
	s := make([]string, 0)
	val := reflect.ValueOf(m)
	for i := 0; i < val.Type().NumField(); i++ {
		ty := val.Type().Field(i)
		if exportTag := ty.Tag.Get("export"); exportTag != "" {
			if expRule := strings.Split(exportTag, ","); (len(expRule) > 1 && expRule[0] == "+") || exportTag == "+" {
				f := reflect.Indirect(val).FieldByName(ty.Name)
				var res string

				switch f.Kind() {
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint:
					res = fmt.Sprint(f)
				case reflect.Float32, reflect.Float64:
					res = fmt.Sprint(f)
				case reflect.String:
					res = f.String()
				case reflect.Bool:
					res = fmt.Sprint(f.Bool())
				case reflect.Struct:
					if f.Type().String() == "time.Time" {
						inputs := make([]reflect.Value, 1)
						inputs[0] = reflect.ValueOf(expRule[1])
						timeSlice := f.MethodByName("Format").Call(inputs)
						if len(timeSlice) > 0 {
							res = fmt.Sprint(timeSlice[0])
						}
					} else {
						res = fmt.Sprint(f.FieldByName(expRule[1]))
					}
				case reflect.Ptr:
					if !f.IsNil() {
						switch f.Elem().Kind() {
						case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint:
							res = fmt.Sprint(f.Elem())
						case reflect.Float32, reflect.Float64:
							res = fmt.Sprint(f.Elem())
						case reflect.String:
							res = f.Elem().String()
						case reflect.Bool:
							res = fmt.Sprint(f.Elem().Bool())
						case reflect.Struct:
							if f.Elem().Type().String() == "time.Time" {
								inputs := make([]reflect.Value, 1)
								inputs[0] = reflect.ValueOf(expRule[1])
								timeSlice := f.Elem().MethodByName("Format").Call(inputs)
								if len(timeSlice) > 0 {
									res = fmt.Sprint(timeSlice[0])
								}
							} else {
								res = fmt.Sprint(f.Elem().FieldByName(expRule[1]))
							}
						}
					}
				}
				s = append(s, res)
			}
		}
	}

	return s
}
