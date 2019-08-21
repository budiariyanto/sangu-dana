# Dana Sangu

## Usage blueprint

1. There is a type named `Client` (`linkaja.Client`) that should be instantiated through `NewClient` which hold any possible setting to the library.
2. There is a gateway classes which you will be using depending on whether you used. The gateway type need a Client instance.
3. Any activity (public token request) is done in the gateway level.

## Example

```go
    danaClient := dana.NewClient()
    danaClient.BaseUrl = "DANA_BASE_URL",
    ---
    ---

    coreGateway := dana.CoreGateway{
        Client: danaClient,
    }

    body := &dana.RequestBody{
        Order: {},
        MerchantId: "216620000000296294081",
        ---
        ---
        ---
    }

    res, _ := coreGateway.Order(req)
```
