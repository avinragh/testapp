package main

import (  
"fmt"
"net/http"
"os"
"github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/aws/session"
"github.com/aws/aws-sdk-go/service/s3"

)


func handler(w http.ResponseWriter, r *http.Request) {  
    svc := s3.New(session.New(),&aws.Config{Region:aws.String("us-east-1")})
	input := &s3.ListBucketsInput{}

	result, err := svc.ListBuckets(input)
	if err != nil {
      fmt.Println(err)
      os.Exit(1)  
        } else {
            fmt.Fprintf(w,"%s",result.GoString())
		}

}


func main() {  
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8086", nil)
}