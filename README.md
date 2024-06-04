## My objectives achieved thanks to blockchain
- Each block must be clearly defined:
    1. Short description
    1. Why I'm doing it
    1. What I would gain from doing it
    1. End date
    1. What I can do after

- A block is :
    1. Timestamp
    1. Index
    1. Objectives data (new type)
    1. Previous hash
    1. Hash

#### Each block will be strutured :

```go
type Block struct {
    Index int
    Timestamp string
    Data DataBlock
    PrevHash string
    Hash string
}

type DataBlock struct {
    ShortDesc string
    WhyDesc string
    GainDesc string
    EndDate string
    WhatsNext string
}
```

---
**With this project, my goal would be:**
- master golang
- deploy this instance on my server (or even deploy in p2p, why not on etherum or solana)
- visualize my objectives via a user-friendly interface
- get in touch with the various domain tools