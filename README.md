createsend-go
=============

createsend-go is a [Go](http://golang.org) library for accessing the [Campaign
Monitor API](http://www.campaignmonitor.com/api/).

Running the tests
-----------------

To run the tests:

```
go test ./createsend
```

To run the included example (in `createsend/example_test.go`), set your Campaign
Monitor API key in the `API_KEY` environment variable (available in Account
Settings).

```
API_KEY=your-api-key go test ./createsend
```

Acknowledgements
----------------

The library's architecture and testing code are adapted from
[go-github](https://github.com/google/go-github).