package main

func callbackMap() error {
  
  pokeapiClient := pokeapi.NewClient()

  resp, err := pokeapiClient.ListLocationAreas()
  if err != nil {
    log.Fatal(err)

  }
  fmt.Println("Location ares:")
  for _, area := range resp.Results{

  }
}
