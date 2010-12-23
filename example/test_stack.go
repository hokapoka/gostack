package main

import(
	"fmt"
	"http"
	"gostack"
	"io/ioutil"
	"json"
	"compress/gzip"
)

func main(){

	testBadges()
}

func testQuestions(){
	r, _, err := http.Get( "http://api.stackoverflow.com/1.0/questions?key=change_me&answers=true" )
	if err != nil {
		fmt.Println(err.String())
		return
	}

	d, err := gzip.NewReader( r.Body )
	if err != nil {
		fmt.Println(err.String())
		return
	}


	var qr gostack.QuestionsResult
	b, _ := ioutil.ReadAll( d ) 

	err = json.Unmarshal(b, &qr)
	if err != nil {
		fmt.Println(err.String())
		return
	}

	fmt.Println(qr)

}

func testBadges(){
	r, _, err := http.Get( "http://api.stackoverflow.com/1.0/badges?key=change_me" )
	if err != nil {
		fmt.Println(err.String())
		return
	}

	d, err := gzip.NewReader( r.Body )
	if err != nil {
		fmt.Println(err.String())
		return
	}

	var br gostack.BadgesResult
	b, _ := ioutil.ReadAll( d ) 

	err = json.Unmarshal(b, &br)
	if err != nil {
		fmt.Println(err.String())
		return
	}

	fmt.Println(br)

}
