package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type postData struct{
	key string
	value string
}

var theTests = []struct{
	name string
	url string
	method string
	params []postData
	expectedStatusCode int
}{
	{"home","/","GET",[]postData{},http.StatusOK},
	{"about","/register","GET",[]postData{},http.StatusOK},
	{"about","/drivers","GET",[]postData{},http.StatusOK},
}

func TestHandler(t *testing.T) {
	mux := goRoutes()
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	for _,e:= range theTests{
		if e.method=="GET"{
			resp,err:=ts.Client().Get(ts.URL+e.url)
			if err!=nil{
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode!=http.StatusOK{
				t.Errorf("for %s, expected %d but got %d",e.name,e.expectedStatusCode,resp.StatusCode)
			}
		}
	}
}