# Garmin API Client

```go
func main() {
    client := garmin.NewClient()
    err := client.Login(os.Getenv("GARMIN_EMAIL"), os.Getenv("GARMIN_PASSWORD"))
    if err != nil {
        panic(err)
    }
    api := garmin.NewAPI(client)
    activities, err := api.Activity.Activities(new(garmin.ActivitySearch).
        WithStart(0).
        WithLimit(20).
        WithActivityType("running"))
    if err != nil {
        panic(err)
    }
    for _, a := range activities {
        fmt.Printf("%+v\n", a)
    }
}
```

Special thanks to [garth](https://github.com/matin/garth), a Garmin API client
written in python. When developing this library I was only able to test with a
Garmin Forerunner 256, if you are using a different device I would recommend
checking the raw json responses to make sure that the structs are complete.

# Other Notes

Download the Garmin Fit SDK from
[here](https://developer.garmin.com/fit/overview/). The architecture is loosely
based on githubs golang api client library.

# TODO

- Look into [this SDK](https://github.com/muktihari/fit) for the Garmin Fit protocol.
