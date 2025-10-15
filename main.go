package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/antchfx/xmlquery"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gogo/protobuf/proto"
	lru "github.com/hashicorp/golang-lru"

	// add more here as blank imports or real usage
	_ "github.com/sirupsen/logrus"
	_ "github.com/valyala/fasthttp"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("run failed: %v", err)
	}
}

func run() error {
	fmt.Println("Demo app (Go 1.18 vuln-scan playground)")

	// Always call your dependency sandbox
	if err := useDepsForTesting(); err != nil {
		return fmt.Errorf("useDepsForTesting: %w", err)
	}

	fmt.Println("All dummy deps executed successfully")
	return nil
}

// useDepsForTesting is a dummy sandbox for dependency inclusion.
// Just keep adding new package usages here to make go.mod/go.sum track them.
func useDepsForTesting() error {
	fmt.Println("Running dummy dependency usage...")

	// --- Example 1: xmlquery ---
	xml := `<root><a>hi</a></root>`
	doc, err := xmlquery.Parse(strings.NewReader(xml))
	if err != nil {
		return fmt.Errorf("xmlquery.Parse: %w", err)
	}
	node := xmlquery.FindOne(doc, "//a")
	if node != nil {
		fmt.Println("xmlquery ok:", node.InnerText())
	}

	// --- Example 2: jwt-go ---
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return fmt.Errorf("jwt SignedString: %w", err)
	}
	fmt.Println("jwt token len:", len(tokenString))

	// --- Example 3: gogo/protobuf ---
	msg := proto.String("hello proto")
	if msg == nil {
		return fmt.Errorf("proto.String returned nil")
	}
	fmt.Println("proto msg:", *msg)

	// --- Example 4: golang-lru ---
	cache, err := lru.New(10)
	if err != nil {
		return fmt.Errorf("lru.New: %w", err)
	}
	cache.Add("foo", "bar")
	v, ok := cache.Get("foo")
	if !ok {
		return fmt.Errorf("lru: missing key after Add")
	}
	fmt.Println("lru value:", v)

	// --- Placeholder for new deps ---
	// Add imports at top and lightweight usage here to have them included in go.mod/go.sum.
	// Example:
	//   import "github.com/gin-gonic/gin"
	//   e := gin.New(); _ = e

	return nil
}
