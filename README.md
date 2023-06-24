# properties #

A simple properties file encoder/decoder.

## usage ##

```go
# encoding
if err := properties.NewEncoder(writer).Encode(v); err != nil {
// ...
}

# decoding
if err := properties.NewDecoder(reader).Decode(v); err != nil {
// ...
}
```
