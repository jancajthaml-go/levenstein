## Performant implementation of Levenstein distance

[![Go Report Card](https://goreportcard.com/badge/jancajthaml-go/levenstein)](https://goreportcard.com/report/jancajthaml-go/levenstein)

Algorithm will measure the minimum number of single-character edits (insertions, deletions or substitutions) required to change one word into the other resulting in integer weight of strings "simmilarity".

![formula](https://wikimedia.org/api/rest_v1/media/math/render/svg/f0a48ecfc9852c042382fdc33c19e11a16948e85)

## Practical usage

Increasing fidelity of search algorithms, e.g if I search for "Mony Lawndaring", and I want system to bring back results for "Money Laundering".

### Usage ###

```
import "github.com/jancajthaml-go/levenstein"

levenstein.Distance("aba", "bba")
```

### Performance ###

```
BenchmarkLevensteinParallel-4     20000000        68.4 ns/op
BenchmarkLevensteinSerial-4       10000000        129 ns/op
```

test on your own by running `make benchmark`

## Resources

* [Wikipedia - Levenstein distance](https://en.wikipedia.org/wiki/Levenshtein_distance)
