package main

func TestCleanInput(t *testing.T) {
  cases := []struct{
    input string
    expected []string
  }{
    {
      input: "hello world",
      expeted: []string{
	"hello",
	"world",
      },
    },
  }
}

for _, cs := range cases {
  actual := cleanInput(cs.input)
  if len(actual) != len(cs.expected){
    t.Errorf("The lengths are not equal: %v vs %V",
      len(actual),
      len(cs.expected),
    )
    continue
  }
}
