# gosrt

Go bindings for the [srt](https://github.com/haivision/srt) network protocol (using cgo)

## Prerequisites

* Go
* SRT (see below)
* pkg-config

### Installing SRT

Currently, SRT doesn't install its headers correctly, so you must checkout the `addHeaders` branch from my fork at https://github.com/russelltg/srt:

```bash
git clone https://github.com/russelltg/srt -b addHeaders
```

Then build and install it:
```bash
cmake . -DCMAKE_INSTALL_PREFIX=<some sane place that pkg-config will find>
make
make install
```

If you need more detailed instructions for building SRT, see [the SRT documentation](https://github.com/Haivision/srt/blob/master/README.md#requirements).

After that, you can run `go get` and `go test` if you want. Should work like a charm.

If you have issues, feel free to create a github issue.
