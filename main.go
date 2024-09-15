package main



import (
	"fmt"
	"github.com/gocolly/colly"
)

// Colly : https://go-colly.org/
// Cobra : https://cobra.dev
// Tutorial: https://benjamincongdon.me/blog/2018/03/01/Scraping-the-Web-in-Golang-with-Colly-and-Goquery/ 

func main() {
	fmt.Println("Welcome to Comical.")
	getPagesByIssues("bone-1991", 1)
}

func getPagesByIssues(comic string, issueNum int) {
	c := colly.NewCollector()

	// On recieving HTML
	c.OnHTML(".page-chapter", func(e *colly.HTMLElement) {
		e.ForEach("img.lazy", func(_ int, img *colly.HTMLElement) {
			imgSrc := img.Attr("alt")
			fmt.Println(imgSrc)
		})
	})

	// On requesting a site
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	// Setup error handling
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request failed: ", err)
	})

	// Visit this site
	url := fmt.Sprintf("https://readcomic.me/comic/%s/issue-%d/all", comic, issueNum)
	err := c.Visit(url)
	if err != nil {
		fmt.Printf("Failed to visit the URL %s: %v\n",url, err)
	}
}
