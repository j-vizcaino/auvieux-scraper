package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

const searchURLBase = "http://www.auvieuxcampeur.fr/catalogsearch/result/?q="

func searchURL(id string) string {
	return fmt.Sprintf("%s%s", searchURLBase, id)
}

type product struct {
	ID    string
	Name  string
	Price float32
}

func convertPrice(pstr string) (float64, error) {
	// Remove trailing ' €'
	re := regexp.MustCompile(`^(\d+),(\d+).*$`)
	s := re.ReplaceAllString(pstr, "$1.$2")
	return strconv.ParseFloat(s, 32)
}

func scrapeProduct(id string) (*product, error) {
	url := searchURL(id)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}

	name := doc.Find(".product-name > a").First().Text()
	priceStr := doc.Find(".regular-price .price").First().Text()

	price, err := convertPrice(priceStr)
	if err != nil {
		return nil, err
	}

	return &product{
		ID:    id,
		Name:  name,
		Price: float32(price),
	}, nil
}

func main() {
	products := os.Args[1:]

	if len(products) == 0 {
		fmt.Printf("Usage: auvieux-scraper product_id [product_id ...]")
		os.Exit(1)
	}

	fmt.Printf("Référence;Désignation;Prix TTC\n")
	for _, productID := range products {
		p, err := scrapeProduct(productID)
		if err != nil {
			fmt.Printf("Error getting product %s (%s)\n", productID, err)
			continue
		}
		fmt.Printf("%s;%s;%.2f\n", p.ID, p.Name, p.Price)
	}

}
