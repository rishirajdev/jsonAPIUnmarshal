package controllers

import (
	"github.com/revel/revel"
	"github.com/shwoodard/jsonapi"
	"net/http"
	"fmt"
	
	


)

type App struct {
	*revel.Controller
}

type Address struct{
		Country string  `jsonapi:"attr,country" `
		State   string  `jsonapi:"attr,state" `
		City    string  `jsonapi:"attr,city" `
		Street  string  `jsonapi:"attr,street" `
		Zipcode string  `jsonapi:"attr,zipcode"`
}

type User struct {
		ID               int  	 `jsonapi:"primary,users"`
		FirstName        string  `jsonapi:"attr,first-name"`
		LastName         string  `jsonapi:"attr,last-name"`
		Username         string  `jsonapi:"attr,username"`
		Contact          string  `jsonapi:"attr,contact"`
		SecondaryContact string  `jsonapi:"attr,secContact"`
		Email            string  `jsonapi:"attr,email"`
		Company          string  `jsonapi:"attr,company"`
		Jobtitle         string  `jsonapi:"attr,job-title"`
		YearsOfExp       string  `jsonapi:"attr,yearsOfExp"`
		Description      string  `jsonapi:"attr,description"`
		Address          []*Address `jsonapi:"relation,address"`
}




type jsonApiResp []interface{}

func (r jsonApiResp) Apply(req *revel.Request, resp *revel.Response) {


	resp.WriteHeader(http.StatusOK, "application/json")

	if err := jsonapi.MarshalManyPayload(resp.Out, r); err != nil {
		http.Error(resp.Out, err.Error(), 500)

	}

}

func (c App) Index() revel.Result {

	   
	user := make([]interface{},0,1)

   
	user = append(user, testBlogForCreate(1))


	return jsonApiResp(user)

}

func (c App) TestUnmarshal() revel.Result{

	u := new(User)	
	
	if err := jsonapi.UnmarshalPayload(c.Request.Body, u); err != nil {
 		fmt.Println(err)
		return c.RenderText("could not parse request")
        	
    	}
	
	return c.RenderJson(u)


}



func testBlogForCreate(i int) *User{
	return &User{
		ID	 	:  1 * i,
		FirstName	: "ABC",
		LastName 	: "DEF",
		Username 	: "abc123",
		Contact  	: "12345678",
		SecondaryContact: "434545665",
		Email 	 	: "abc@abc.com",
		Company  	: "Torant",
		Jobtitle 	: "xyz",
		YearsOfExp 	: "5",
		Description 	: "hdgfsdfg",
		Address		: []*Address{&Address{
						Country :"INDIA",
						State	: "MH",
						City	:  "PU",
						Street 	: "ahbfshdf",		
						Zipcode : "12345",
					},
				},
				
		}
	
}