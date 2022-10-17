package handlers

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"net/url"
	"encoding/json"
	"github.com/udaysonu/ober/models"
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
	{"about","/register","POST",[]postData{},http.StatusOK},
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
		}else if e.method=="POST"{
			// reqBody:="first_name=uday&email=udayyadusonu@1"

			reqBody:=url.Values{}
			reqBody.Add("first_name","uday")
			reqBody.Add("email","udaysonubakka123@gmail.com")
			req,_:=http.NewRequest("POST","/register",strings.NewReader(reqBody.Encode()))


			// req,_:=http.NewRequest("POST","/register",strings.NewReader(reqBody))
			req.Header.Set("Content-Type","application/x-www-form-urlencoded")
			rr:=httptest.NewRecorder()

			handler:=http.HandlerFunc(Repo.CheckPostRequest)
			handler.ServeHTTP(rr,req)
		}
	}
}

func TestCheckPost(t *testing.T){
			resp:=models.Driver{
				FirstName:"udaykkrna",
			}
			out,_:=json.MarshalIndent(resp,"","  ")
 			req,_:=http.NewRequest("POST","/post",strings.NewReader(string(out[:])))
			req.Header.Set("Content-Type","application/json")
			rr:=httptest.NewRecorder()

			handler:=http.HandlerFunc(Repo.PostTest)
			handler.ServeHTTP(rr,req)

			var response models.Driver
			_=json.Unmarshal([]byte(rr.Body.String()),&response)
			fmt.Println("response::",response)
}
func TestCheckSession(t *testing.T){
	driver:=models.Driver{
	DriverId :"1233",
	FirstName : "kiran",
	MiddleName :"",
	LastName :"",
	Email :"uday@gmail.com",
	Age :33,
	PhoneNumber :432432432,
	Password :"udayyareffdysonu@1",
	Latitude : 34.323,
	Longitude : 43.232,
	Bio : "gerat driver",
	}

	req,_:=http.NewRequest("GET","/make-reservation",nil)

	ctx := getCtx(req)
	req=req.WithContext(ctx)

	rr:= httptest.NewRecorder()

	session.Put(ctx,"session-data",driver)

	handler:=http.HandlerFunc(Repo.CheckSession)
	handler.ServeHTTP(rr,req)

	if rr.Code!=http.StatusOK{
		t.Errorf("Reservation handler returned wrong response code: go")
	}
}

func getCtx(req *http.Request)context.Context{
	ctx,err:=session.Load(req.Context(),req.Header.Get("X-Session"))
	if err!=nil{
		fmt.Println(err)
	}
	return ctx
}