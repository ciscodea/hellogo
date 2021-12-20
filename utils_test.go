package utils

import "testing"

func TestHelloWorld(t *testing.T) {
	testFunc := HelloWorld()
	expected := "Hello World from hellogo!"

	if testFunc != expected {
		t.Errorf("Message was incorrect, got %s expected %s", testFunc, expected)
	}
}

func TestHelloWorldName(t *testing.T) {
	tables := []struct {
		name   string
		expect string
	}{
		{"Pedro", "Hello Pedro!"},
		{"Yoselin", "Hello Yoselin!"},
	}

	for _, item := range tables {
		message := HelloWorldName(item.name)
		if message != item.expect {
			t.Errorf("Message was incorrect, got %s expected %s", message, item.expect)
		}
	}
}
