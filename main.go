package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const githubSearchURL = "https://api.github.com/search/repositories"

type GitHubSearchResponse struct {
	TotalCount int          `json:"total_count"`
	Items      []GitHubRepo `json:"items"`
}

type GitHubRepo struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"html_url"`
}

func main() {

	// start
	welcome()

	//flags
	fileName := flag.String("f", "", "file name")
	accessToken := flag.String("t", "", "GitHub access token")
	flag.Parse()

	if *fileName == "" {
		fmt.Println("Please specify a file name using the -f flag.")
		os.Exit(1)
	}

	if *accessToken == "" {
		fmt.Println("Please specify a GitHub access token using the -t flag.")
		os.Exit(1)
	}

	// Read the search term from the file
	searchBytes, err := ioutil.ReadFile(*fileName)
	if err != nil {
		panic(err)
	}
	searchTerm := string(searchBytes)

	searchURL := fmt.Sprintf("%s?q=%s", githubSearchURL, searchTerm)
	req, err := http.NewRequest("GET", searchURL, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("token %s", *accessToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var searchResponse GitHubSearchResponse
	err = json.Unmarshal(body, &searchResponse)
	if err != nil {
		panic(err)
	}

	// Print the search results
	fmt.Printf("Total count: %d\n", searchResponse.TotalCount)
	for _, item := range searchResponse.Items {
		fmt.Printf("Name: %s\n", item.Name)
		fmt.Printf("Description: %s\n", item.Description)
		fmt.Printf("URL: %s\n", item.URL)
	}

	// Write the search results to a JSON file
	file, err := json.MarshalIndent(searchResponse, "", "    ")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("search_results.json", file, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Search results written to search_results.json")
}

func welcome() {

	//

	been := `
          _____                    _____                _____                    _____                    _____                    _____                            _____                    _____                    _____                    _____                    _____                    _____          
         /\    \                  /\    \              /\    \                  /\    \                  /\    \                  /\    \                          /\    \                  /\    \                  /\    \                  /\    \                  /\    \                  /\    \         
        /::\    \                /::\    \            /::\    \                /::\____\                /::\____\                /::\    \                        /::\    \                /::\    \                /::\    \                /::\    \                /::\    \                /::\____\        
       /::::\    \               \:::\    \           \:::\    \              /:::/    /               /:::/    /               /::::\    \                      /::::\    \              /::::\    \              /::::\    \              /::::\    \              /::::\    \              /:::/    /        
      /::::::\    \               \:::\    \           \:::\    \            /:::/    /               /:::/    /               /::::::\    \                    /::::::\    \            /::::::\    \            /::::::\    \            /::::::\    \            /::::::\    \            /:::/    /         
     /:::/\:::\    \               \:::\    \           \:::\    \          /:::/    /               /:::/    /               /:::/\:::\    \                  /:::/\:::\    \          /:::/\:::\    \          /:::/\:::\    \          /:::/\:::\    \          /:::/\:::\    \          /:::/    /          
    /:::/  \:::\    \               \:::\    \           \:::\    \        /:::/____/               /:::/    /               /:::/__\:::\    \                /:::/__\:::\    \        /:::/__\:::\    \        /:::/__\:::\    \        /:::/__\:::\    \        /:::/  \:::\    \        /:::/____/           
   /:::/    \:::\    \              /::::\    \          /::::\    \      /::::\    \              /:::/    /               /::::\   \:::\    \               \:::\   \:::\    \      /::::\   \:::\    \      /::::\   \:::\    \      /::::\   \:::\    \      /:::/    \:::\    \      /::::\    \           
  /:::/    / \:::\    \    ____    /::::::\    \        /::::::\    \    /::::::\    \   _____    /:::/    /      _____    /::::::\   \:::\    \            ___\:::\   \:::\    \    /::::::\   \:::\    \    /::::::\   \:::\    \    /::::::\   \:::\    \    /:::/    / \:::\    \    /::::::\    \   _____  
 /:::/    /   \:::\ ___\  /\   \  /:::/\:::\    \      /:::/\:::\    \  /:::/\:::\    \ /\    \  /:::/____/      /\    \  /:::/\:::\   \:::\ ___\          /\   \:::\   \:::\    \  /:::/\:::\   \:::\    \  /:::/\:::\   \:::\    \  /:::/\:::\   \:::\____\  /:::/    /   \:::\    \  /:::/\:::\    \ /\    \ 
/:::/____/  ___\:::|    |/::\   \/:::/  \:::\____\    /:::/  \:::\____\/:::/  \:::\    /::\____\|:::|    /      /::\____\/:::/__\:::\   \:::|    |        /::\   \:::\   \:::\____\/:::/__\:::\   \:::\____\/:::/  \:::\   \:::\____\/:::/  \:::\   \:::|    |/:::/____/     \:::\____\/:::/  \:::\    /::\____\
\:::\    \ /\  /:::|____|\:::\  /:::/    \::/    /   /:::/    \::/    /\::/    \:::\  /:::/    /|:::|____\     /:::/    /\:::\   \:::\  /:::|____|        \:::\   \:::\   \::/    /\:::\   \:::\   \::/    /\::/    \:::\  /:::/    /\::/   |::::\  /:::|____|\:::\    \      \::/    /\::/    \:::\  /:::/    /
 \:::\    /::\ \::/    /  \:::\/:::/    / \/____/   /:::/    / \/____/  \/____/ \:::\/:::/    /  \:::\    \   /:::/    /  \:::\   \:::\/:::/    /          \:::\   \:::\   \/____/  \:::\   \:::\   \/____/  \/____/ \:::\/:::/    /  \/____|:::::\/:::/    /  \:::\    \      \/____/  \/____/ \:::\/:::/    / 
  \:::\   \:::\ \/____/    \::::::/    /           /:::/    /                    \::::::/    /    \:::\    \ /:::/    /    \:::\   \::::::/    /            \:::\   \:::\    \       \:::\   \:::\    \               \::::::/    /         |:::::::::/    /    \:::\    \                       \::::::/    /  
   \:::\   \:::\____\       \::::/____/           /:::/    /                      \::::/    /      \:::\    /:::/    /      \:::\   \::::/    /              \:::\   \:::\____\       \:::\   \:::\____\               \::::/    /          |::|\::::/    /      \:::\    \                       \::::/    /   
    \:::\  /:::/    /        \:::\    \           \::/    /                       /:::/    /        \:::\__/:::/    /        \:::\  /:::/    /                \:::\  /:::/    /        \:::\   \::/    /               /:::/    /           |::| \::/____/        \:::\    \                      /:::/    /    
     \:::\/:::/    /          \:::\    \           \/____/                       /:::/    /          \::::::::/    /          \:::\/:::/    /                  \:::\/:::/    /          \:::\   \/____/               /:::/    /            |::|  ~|               \:::\    \                    /:::/    /     
      \::::::/    /            \:::\    \                                       /:::/    /            \::::::/    /            \::::::/    /                    \::::::/    /            \:::\    \                  /:::/    /             |::|   |                \:::\    \                  /:::/    /      
       \::::/    /              \:::\____\                                     /:::/    /              \::::/    /              \::::/    /                      \::::/    /              \:::\____\                /:::/    /              \::|   |                 \:::\____\                /:::/    /       
        \::/____/                \::/    /                                     \::/    /                \::/____/                \::/____/                        \::/    /                \::/    /                \::/    /                \:|   |                  \::/    /                \::/    /        
                                  \/____/                                       \/____/                  ~~                       ~~                               \/____/                  \/____/                  \/____/                  \|___|                   \/____/                  \/____/
`

	fmt.Printf("%v", been)
	fmt.Printf("\t\t v 1.0 - 2023\n\n")
	fmt.Println("Are You Ready ? :-) ")
	time.Sleep(2 * time.Second)
	fmt.Println("3")
	time.Sleep(2 * time.Second)
	fmt.Println("2")
	time.Sleep(1 * time.Second)
	fmt.Println("▄︻̷̿┻̿═━一")
}
