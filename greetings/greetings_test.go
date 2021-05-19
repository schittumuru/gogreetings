package greetings


import "testing"
import "regexp"


func TestHelloName(t *testing.T) {

	name := "sasi"
	want := regexp.MustCompile(`\b`+name+`\b`)

	message, err := Hello(name)

	if !want.MatchString(message) || err != nil {
		t.Fatalf(`Hello("Sasi") = %q, %v, want match for %#q, nil`, message, err, want)
	}

}

func TestHelloEmpty(t *testing.T) {
	message, err := Hello("")

	if message != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, message, err)
	}
}
