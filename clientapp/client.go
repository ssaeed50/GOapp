package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {

	//Open the file to upload

	file, err := os.Open("C:/Users/ssaeed/Desktop/GOproject/serverapp/")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Create a new multipart form

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add the file to the form
	part, err := writer.CreateFormFile("file", "Test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = io.Copy(part, file)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Add the file name to the form
	_ = writer.WriteField("name", "Test.txt")

	// Close the form
	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Send the POST request to the server
	req, err := http.NewRequest("POST", "http://localhost:8080/upload", body)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	//////this for http request////////
	/*client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	*/

	///////This for TLS configuration////////////

	config := &tls.Config{InsecureSkipVerify: true}
	client := &http.Client{Transport: &http.Transport{TLSClientConfig: config}}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// Read the response
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed to read response: %v", err)
	}
	fmt.Printf("Response: %s", string(respBody))

	// Check the response status code to determine whether the upload was successful or not
	if resp.StatusCode == 200 {
		fmt.Println("File uploaded successfully")
	} else {
		fmt.Println("File upload failed")
	}

}
