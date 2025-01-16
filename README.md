# Money-go

The Money-go Golang library is used for dealing with money objects based on crypto or fiat currencies, This is partially inspired by the Ruby Money gem and largely inspired by the implementation by one of the best to ever do it.

## Instructions

Before using this library, make sure to initialise the currencies using one of the below methods.

- `InitFromFile`: expects to be passed the path to a json file which contains an array of currency objects. See an example in [examples/currencies.json](examples/currencies.json).
- `InitFromJsonString`: expects to be passed a json string.
- `InitFromCurrencies`: expects to be passed a list of golang currency structs.
