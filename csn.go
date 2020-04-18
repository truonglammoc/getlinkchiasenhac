package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func main(){
	casi := os.Args[1]
	downloadAll(casi)
}
func stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}
func unique(arrString []string) []string {
	cleaned := []string{}
	for _, value := range arrString {
		if !stringInSlice(value, cleaned) {
			cleaned = append(cleaned, value)
		}
	}
	return cleaned
}
func downloadAll(casi string) {
	name := strings.Replace(casi, " ", "_", -1)
	os.Mkdir(name,777)
	links := getMusicOfCasi(casi)
	for _, value := range links {
		qualityLink := getBestLink(value)
		cmd := exec.Command("wget", "-nc", "-P", name, qualityLink)
		cmd.Run()
	}
}
func fileGetContent(url string) (string) {
	resp, err := http.Get(url)
	if err == nil {
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			return string(body)
		}
	}
	return ""
}
func getLinkMusic(url string) []string {
	content := fileGetContent(url)
	regex := regexp.MustCompile("[^\"' ]+\\.(m4a|mp3|flac)")
	result := unique(regex.FindAllString(content, -1))
	return result
}
func getBestLink(url string) string {
	links := getLinkMusic(url)
	return links[len(links) - 1]
}
func getMusicOfCasi(casi string) []string {
	url := fmt.Sprintf("http://search.chiasenhac.vn/search.php?s=%s&mode=artist&order=quality&cat=music", strings.Replace(casi, " ", "+", -1))
	content := fileGetContent(url)
	regex := regexp.MustCompile("https:\\/\\/beta\\.chiasenhac\\.vn\\/mp3[^\"]+\\.html")
	result := regex.FindAllString(content, -1)
	return result
}
