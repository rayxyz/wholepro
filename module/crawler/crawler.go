package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/mvdan/xurls"
)

var (
	maxmatches = 1 << 10
	hrefs      []*string
	filePath   string
	// pattern    = "[http://|][www|][a-z|0-9]{1,}.{1,}[com|net|cn]"
	pattern = `www.[a-z|0-9]{1,}.{1,}[com|net|cn]`
)

func init() {
	log.Println("Initiating...")
	filePath = path.Join(os.Getenv("home"), "hrefs.txt")
}

// func handler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("The web crawler is working..."))
// 	err := craw(w, r)
// 	if err != nil {
// 		fmt.Fprintln(w, "Start web crawler failed.")
// 		log.Fatal(err)
// 	}
// }

func craw() error {
	crawch := make(chan int)
	go func() {
		doCraw("http://www.csdn.net", crawch)
	}()
	<-crawch
	log.Println("File path: ", filePath)
	p, err := os.Create(filePath)
	if err != nil {
		log.Println("Create hrefs.txt error.")
		return err
	}
	if p != nil {
		log.Println("Created a file with path: ")
		defer p.Close()
	}

	var hrefsContent string
	for _, href := range hrefs {
		hrefsContent += *href + "\n"
	}
	len, err := p.WriteString(hrefsContent)
	if err != nil {
		log.Println("Write hrefs content error.")
		return err
	}
	log.Printf("%d bytes has been written.", len)
	return nil
}

func doCraw(href string, crawch chan int) {
	log.Println("Now crawling ", href)
	resp, err := http.Get(href)
	if err != nil {
		log.Print("Crawling failed.")
		return
	}
	if resp != nil {
		defer resp.Body.Close()

		buf, _ := ioutil.ReadAll(resp.Body)
		bufStr := string(buf)
		// matches := regexp.MustCompile(pattern).FindAllString(bufStr, maxmatches)
		matches := xurls.Strict.FindAllString(bufStr, -1)
		if matches == nil {
			log.Println("No matches.")
			return
		}
		for i := 0; i < len(matches); i++ {
			exists := checkHrefExists(matches[i])
			if exists {
				continue
			}
			hrefs = append(hrefs, &matches[i])
			crawchx := make(chan int)
			go func() {
				doCraw(matches[i], crawchx)
				crawchx <- 2
			}()
			<-crawchx
		}
	}
	crawch <- 1
}

func checkHrefExists(hrefToCheck string) bool {
	if hrefToCheck == "" {
		return true
	}
	for _, href := range hrefs {
		if hrefToCheck == *href {
			return true
		}
	}
	return false
}

func main() {
	// http.HandleFunc("/go", handler)
	// http.ListenAndServe(":8888", nil)
	log.Println("Crawler initiation complete.")
	err := craw()
	if err != nil {
		log.Println("Error in main of crawler.")
	}
}
