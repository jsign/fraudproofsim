# Fraud-Proof Blockchain Simulator

## Motivation
This is the repo corresponding to [Simulating a fraud-proof blockchain
](https://medium.com/@jsign.uy/simulating-fraud-proof-blockchain-network-24bc55b1237c) article. 

## CLI
```
$ go run main.go
It permits to compare, solve and verify fraud-proof networks.
Usage:
  fraudproofsim [command]
Available Commands:
  compare     Compares the Standard and Enhanced models
  help        Help about any command
  solve       Solves c for k, s and p
  verifypaper Verifies setups calculated in the paper
Flags:
      --enhanced   run an Enhanced Model
  -h, --help       help for fraudproofsim
      --n int      number of iterations to run per instance (default 500)
Use "fraudproofsim [command] --help" for more information about a command.
```

### _verifypaper_ command
```
$ go run main.go help solve
It solves c for k, s and p (p, within a threshold)
Usage:
  fraudproofsim solve [k] [s] [p] [threshold?] [flags]
Flags:
  -h, --help   help for solve
Global Flags:
      --enhanced   run an Enhanced Model
      --n int      number of iterations to run per instance (default 500)
$ go run main.go verifypaper
k=16, s=50, c=28 => p=1 37ms
k=16, s=20, c=69 => p=0.994 28ms
k=16, s=10, c=138 => p=0.988 37ms
k=16, s=5, c=275 => p=0.986 37ms
k=16, s=2, c=690 => p=0.99 63ms
k=32, s=50, c=112 => p=0.996 137ms
k=32, s=20, c=280 => p=0.994 131ms
k=32, s=10, c=561 => p=0.988 136ms
k=32, s=5, c=1122 => p=0.992 143ms
k=32, s=2, c=2805 => p=0.994 175ms
k=64, s=50, c=451 => p=0.996 464ms
k=64, s=20, c=1129 => p=0.996 536ms
k=64, s=10, c=2258 => p=0.992 510ms
k=64, s=5, c=4516 => p=0.988 527ms
k=64, s=2, c=11289 => p=0.996 679ms
k=128, s=50, c=1811 => p=0.992 2193ms
k=128, s=20, c=4500 => p=0.702 2068ms
exit status 2
```

### _solve_ command
```
$ go run main.go help solve
It solves c for k, s and p (p, within a threshold)
Usage:
  fraudproofsim solve [k] [s] [p] [threshold?] [flags]
Flags:
  -h, --help   help for solve
Global Flags:
      --enhanced   run an Enhanced Model
      --n int      number of iterations to run per instance (default 500)
      
$ go run main.go solve 64 10 .99 0.005
Solving for (k:64, s:10, p:0.99, threshold:0.005)
[1, 16384]: c=8192 p=1
[1, 8192]: c=4096 p=1
[1, 4096]: c=2048 p=0
[2048, 4096]: c=3072 p=1
[2048, 3072]: c=2560 p=1
[2048, 2560]: c=2304 p=1
[2048, 2304]: c=2176 p=0.002
[2176, 2304]: c=2240 p=0.902
[2240, 2304]: c=2272 p=1
[2240, 2272]: c=2256 p=0.994
Solution c=2256 with p=0.994 (4900ms)
```

### _compare_ command
```
$ go run main.go help compare
Compares Standard and Enhanced model to understand their impact on soundness
Usage:
  fraudproofsim compare [k] [s] [#points] [flags]
Flags:
  -h, --help   help for compare
Global Flags:
      --enhanced   run an Enhanced Model
      --n int      number of iterations to run per instance (default 500)
      
$ go run main.go compare 64 10 25
Solving c for (k: 64, s: 10) with precision .99+-.005:
[1, 16384]: c=8192 p=1
[1, 8192]: c=4096 p=1
[1, 4096]: c=2048 p=0
[2048, 4096]: c=3072 p=1
[2048, 3072]: c=2560 p=1
[2048, 2560]: c=2304 p=1
[2048, 2304]: c=2176 p=0
[2176, 2304]: c=2240 p=0.896
[2240, 2304]: c=2272 p=0.998
[2240, 2272]: c=2256 p=0.99
Found solution c=2256, now generating 25 points in [.50*c,1.5*c]=[1128, 3384]:
0%
3%
7%
11%
15%
19%
23%
27%
31%
35%
39%
43%
47%
51%
55%
59%
63%
67%
71%
75%
79%
83%
87%
91%
95%
99%
Plotted in plot.png
```
Output .png file:
<dl>
<img src="https://cdn-images-1.medium.com/max/800/1*ezgV1GyG9Lnss7zFhQ5JLw.png"></src>
</dl>
