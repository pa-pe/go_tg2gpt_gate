package request

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"golang.org/x/text/language"
	"net/http"
	"reflect"
	"time"
	"upserv/logger"
	"upserv/src/apperror"
)

type ListParams struct {
	Offset int `schema:"offset"`
	Limit  int `schema:"limit"`
}

var supportedLangs = []language.Tag{
	language.English, // The first language is used as fallback.
}

// types definition
type IRequest interface {
	Validate() *apperror.IError
	InitDefaults()
}

type Cookie struct {
	Name  string
	Value string
}

type Cookies struct {
	list map[string]*Cookie
}

func (c *Cookies) Get(name string) *Cookie {
	return c.list[name]
}

func (c *Cookies) GetAll() []*Cookie {
	var res []*Cookie
	for _, cc := range c.list {
		res = append(res, cc)
	}
	return res
}

// extract header language
func ExtractLang(r *http.Request) {
	if r.Header.Get("Lang") == "" {
		var matcher = language.NewMatcher(supportedLangs)
		tag, _ := language.MatchStrings(matcher, r.Header.Get("Accept-Language"))
		b, _ := tag.Base()
		r.Header.Set("Lang", b.String())
	}
	ctx := context.WithValue(r.Context(), "Lang", r.Header.Get("Lang"))
	nR := r.WithContext(ctx)
	*r = *nR
}

// main extract params from request method
func Extract(obj IRequest, r *http.Request) *apperror.IError {
	ctx := r.Context()
	decoder := schema.NewDecoder()
	err := r.ParseForm()
	//if hasField(obj, "ID") {
	//	data := mux.Vars(r)
	//	r.Form.Add("id", data["id"])
	//}
	data := mux.Vars(r)
	if data != nil && len(data) > 0 {
		for field, value := range data {
			r.Form.Add(field, value)
		}
	}
	if err != nil {
		logger.Log.Info("Parse form error")
		logger.Log.Debug(err.Error())
		return apperror.BadRequestGeneral(ctx)
	}
	decoder.RegisterConverter(time.Location{}, timeRegionConverter)
	decoder.RegisterConverter(time.Time{}, timeConverter)
	obj.InitDefaults()
	err = decoder.Decode(obj, r.Form)
	extractCookiesToObj(obj, r.Cookies())
	if err != nil {
		logger.Log.Info("Decode error")
		logger.Log.Debug(err.Error())
		return apperror.BadRequestOnParams(ctx).WithMsg(err.Error())
	}
	er := obj.Validate()
	if er != nil {
		logger.Log.Info("Validate error")
		logger.Log.Debug(er.Error())
		return apperror.BadRequestOnParams(ctx).WithMsg(er.Error())
	}
	return nil
}

func timeConverter(value string) reflect.Value {
	if v, err := time.Parse("2006-01-02T15:04", value); err == nil {
		return reflect.ValueOf(v)
	}
	return reflect.Value{}
}

func timeRegionConverter(value string) reflect.Value {
	if loc, err := time.LoadLocation(value); err == nil {
		if loc != nil {
			return reflect.ValueOf(*loc)
		}
	}
	return reflect.Value{}
}
func extractCookiesToObj(obj interface{}, cookies []*http.Cookie) {
	element := reflect.ValueOf(obj).Elem()
	cookiesField := element.FieldByName("Cookies")
	if !cookiesField.IsValid() {
		return
	}
	newCookies := Cookies{}
	newCookies.list = make(map[string]*Cookie)
	for _, val := range cookies {
		c := &Cookie{
			Name:  val.Name,
			Value: val.Value,
		}
		newCookies.list[val.Name] = c
	}
	if cookiesField.CanSet() {
		cookiesField.Set(reflect.ValueOf(newCookies))
	}
}
