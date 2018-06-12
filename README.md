# go-distribution

A Golang model of distribution of popularity. For example, this is used when a zipf-like distribution of popularity in the web access or a gamma-like distribution of popularity in the Video-on-Demand in YouTube.

`go get github.com/roamer7038/go-distribution`

- zipf
- gamma

##  API

### Creating a Distribution

`zipf := distribution.NewZipf(n, alpha)`

`gamma := distribution.NewGamma(n, k, thita)`

### Properties

- `n` : the number of elements.
- `alpha` : the value of the exponent characterizing the distribution.
- `k` : a shape parameter.
- `thita`: a scale parameter θ.

### Probability Functions

- `Pdf(x)` : the probability density function, which describes a probability taking of the specified rank.
- `Cdf(x)` : the cumulative distribution function, which describes a probability up to specified order.

