# `gostruct`: Convert example JSON objects to Go Struct definitions

This is an imperfect but handy little tool to quickly convert JSON objects into Go struct definitions. For example:

```json
{
  "id": "623f6018-6cf0-431e-84a5-43aeb8be4372",
  "extra_data": null,
  "created": 1405637075
}
```

becomes

```go
type NewObject struct {
        Id              string `json:"id"`
        ExtraData               interface{} `json:"extra_data"`
        Created         float64 `json:"created"`
}
```

## Installation

Haven't done a release yet, so for now just run `go install -v github.com/seanph/gostruct@latest`.


## Usage

`gostruct` can be used via either piped input or direct input.

**Pipe**
```bash
echo '{"id": 5, "name": "test"}' | gostruct
```

**Direct**
```bash
gostruct --dump
```
Then paste in the JSON data and hit `Ctrl+D` to signal EOF.

## Caveats

- Currently only supports `map[string]interface{}` shaped JSON objects. Arrays or alternative maps won't work.
- Currently outputs ints and floats as `float64` as this is the default Go behaviour. Worth manually checking back over these!
- Formatting isn't great, but your editor should fix this.
