package main

import (
	"flag"
	"github.com/upamune/go-esa/esa"
	"net/url"
	"os"
	"strconv"
	"time"
)

const EsaRequestInterval = 12

func main() {

	var err error
	var errorPosts []string
	var postUrl string

	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	esaTeamName := fs.String("e", "", "esa team name")
	esaToken := fs.String("eToken", "", "esa Access Token")
	startIdStr := fs.String("start-id", "", "Starting delete esa post ID")
	endIdStr := fs.String("end-id", "", "Ending delete esa post ID")
	fs.Parse(os.Args[1:])

	if *esaTeamName == "" {
		panic("Please type -e (esa team name)")
	}

	if *esaToken == "" {
		panic("Please type -eToken (esa Access Token)")
	}

	startId, err := strconv.Atoi(*startIdStr)
	if err != nil {
		panic("-start-id should be integer.")
	}

	endId, err := strconv.Atoi(*endIdStr)
	if err != nil {
		panic("-end-id should be integer.")
	}

	client := esa.NewClient(*esaToken)

	for i := startId; i <= endId; i++ {
		postUrl = esaPostUrl(*esaTeamName, i)
		print("Processing: " + postUrl + " ... ")

		err = client.Post.Delete(*esaTeamName, i)

		if err != nil {
			println("Failed!")
			errorPosts = append(errorPosts, postUrl)
		}

		println("Complete!")
		time.Sleep(EsaRequestInterval * time.Second)
	}

	printErrorPosts(errorPosts)
}

func esaPostUrl(teamName string, postID int) string {
	postUrl := url.URL{
		Scheme: "https",
		Host:   teamName + ".esa.io",
		Path:   "/posts/" + strconv.Itoa(postID),
	}

	return postUrl.String()
}

func printErrorPosts(errorPosts []string) {
	println("---------------------------")
	println("Failed Posts")
	for key, errorPost := range errorPosts {
		println(strconv.Itoa(key+1) + " : " + errorPost)
	}
	println("---------------------------")
}
