# Golang Talks

A collection of presentations, articles, and sample code about Go (Golang), presented by Chee Leong ([@klrkdekira](https://github.com/klrkdekira)) primarily at Golang Meetup Kuala Lumpur / Go Malaysia.

---

## Table of Contents

- [How to View and Run Locally](#how-to-view-and-run-locally)
- [Presentations by Year](#presentations-by-year)
  - [2020](#2020)
  - [2019](#2019)
  - [2018](#2018)
  - [2017](#2017)
  - [2014](#2014)
- [Repository Structure](#repository-structure)
- [License](#license)

---

## How to View and Run Locally

These slides are written in the format used by the official Go [`present`](https://pkg.go.dev/golang.org/x/tools/cmd/present) tool. Running `present` locally provides interactive slide navigation and allows executing Go code snippets directly in your browser.

### 1. Run with `go tool` (Recommended)

Since Go 1.24+, executable tools can be tracked directly in `go.mod` using the `tool` directive. The `present` tool is pre-configured in this repository's [`go.mod`](go.mod).

To launch the presentation server, simply run:

```bash
go tool present
```

Or specify a custom address/port:

```bash
go tool present -http=:3999
```

No manual installation into `$PATH` is required—Go resolves, builds, and executes `present` automatically.

---

### Alternative: Run via `go install`

If running outside this module or using an older Go version:

```bash
go install golang.org/x/tools/cmd/present@latest
present -http=:3999
```

---

### 2. Open in Browser

Open your browser and navigate to:

```text
http://localhost:3999
```

Click on any `.slide` or `.article` file to view the presentation.

#### Navigation Shortcuts in `present`:
- `Page Down` / `Right Arrow` / `Space`: Next slide
- `Page Up` / `Left Arrow`: Previous slide
- `N`: Toggle speaker notes (if present)
- Click **Run** on embedded code snippets to compile and execute them on your machine.

> **Note on Online Viewing:**
> You can also view the slides online via [go-talks.appspot.com](https://go-talks.appspot.com). Links are provided below. If you encounter rate-limiting errors (HTTP 403/500 from GitHub API) on the public `go-talks` service, viewing the slides locally using `present` is recommended.

---

## Presentations by Year

### 2020

- **Go Kuala Lumpur October 2020**
  - Links: [Slide (go-talks)](https://go-talks.appspot.com/github.com/klrkdekira/golang-talks/2020/october/main.slide) | [Source File](2020/october/main.slide) | [Code Samples](2020/october/codes/)
  - **Topics:** Go 1.15.3 release notes, Go Developer Survey 2020, HashiCorp Go libraries, acceptance of the static file embedding proposal (`//go:embed`), Prisma Client Go, and rqlite distributed database.

- **Go Kuala Lumpur September 2020**
  - Links: [Slide (go-talks)](https://go-talks.appspot.com/github.com/klrkdekira/golang-talks/2020/september/main.slide) | [Source File](2020/september/main.slide) | [Code Samples](2020/september/codes/)
  - **Topics:** Community meetup reboot, Go 1.15.2, Go Developer Survey results, revised Generics draft design and playground, embedded static assets draft, open sourcing of `pkg.go.dev`, GORM 2.0, and the new Protocol Buffers API (`google.golang.org/protobuf`).

- **Golang Meetup Kuala Lumpur February 2020**
  - Links: [Slide (go-talks)](https://go-talks.appspot.com/github.com/klrkdekira/golang-talks/2020/february/main.slide) | [Source File](2020/february/main.slide)
  - **Topics:** Go 1.13.8 release, community swag, `pigo` pure Go face detection, Kubethanos chaos pod killer, static file embedding proposal for `cmd/go`, sqlc compile-time SQL query compiler, and dynamic instrumentation.

### 2019

- **Optimization Tips and Tricks** (November 2019)
  - Links: [Slide (go-talks)](https://go-talks.appspot.com/github.com/klrkdekira/golang-talks/2019/compiler/main.slide) | [Source File](2019/compiler/main.slide) | [Code Samples](2019/compiler/codes/)
  - **Topics:** Compiler-level optimizations and memory tuning: allocation benchmarking (`go test -bench=. -benchmem`), escape analysis (`-gcflags="-m"`), stack vs. heap allocation, slice bounds check elimination, string-to-byte conversion overhead, and reducing GC pressure with `sync.Pool`.

- **Golang Meetup November 2019**
  - Links: [Slide (go-talks)](https://go-talks.appspot.com/github.com/klrkdekira/golang-talks/2019/november/main.slide) | [Source File](2019/november/main.slide)
  - **Topics:** Celebrating Go turns 10, Go 1.13.4 release, launch of `go.dev`, Go Developer Survey 2019, Sourcegraph Go style guide, and Pkger static file embedding.

- **Golang Meetup October 2019**
  - Links: [Slide (go-talks)](https://go-talks.appspot.com/github.com/klrkdekira/golang-talks/2019/october/main.slide) | [Source File](2019/october/main.slide)
  - **Topics:** Go 1.13.3 release, Ristretto memory-bound cache, Watermill event-driven library, Gizmo microservice toolkit, v8go JavaScript engine bindings, and `benchdraw`.

- **Golang Meetup September 2019**
  - Links: [Slide (go-talks)](https://go-talks.appspot.com/github.com/klrkdekira/golang-talks/2019/september/main.slide) | [Source File](2019/september/main.slide)
  - **Topics:** Go 1.13 release breakdown (Go module mirror and checksum database enabled by default, number literals, error wrapping with `errors.Is`/`As`), gopy, gorilla/websocket, Dgraph, GoLand 2019.3 EAP, TinyGo, and fasthttp.

- **Delve - Golang Debugger** (August 2019)
  - Links: [Slide (go-talks)](https://go-talks.appspot.com/github.com/klrkdekira/golang-talks/2019/delve/main.slide) | [Source File](2019/delve/main.slide)
  - **Topics:** Live demo and presentation on debugging Go applications using Delve (`dlv`).

- **Golang Meetup August 2019**
  - Links: [Slide (go-talks)](https://go-talks.appspot.com/github.com/klrkdekira/golang-talks/2019/august/main.slide) | [Source File](2019/august/main.slide)
  - **Topics:** Go 1.12.7 release, retiring the `try` error handling proposal, contracts and early generics draft design for Go 2, Russ Cox's proposal process, Yaegi Go interpreter, and TinyGo 0.7.0.

- **Different Faces of JSON Processing** (June 2019)
  - Links: [Slide (go-talks)](https://go-talks.appspot.com/github.com/klrkdekira/golang-talks/2019/json/main.slide) | [Source File](2019/json/main.slide) | [Code Samples](2019/json/codes/)
  - **Topics:** Overcoming JSON unmarshaling challenges in Go: handling heterogeneous/dynamic JSON without generics, limitations of `map[string]interface{}`, numeric types with `json.Number`, deferred unmarshaling using `json.RawMessage`, custom unmarshalers, and reflection techniques.

- **Golang Meetup June 2019**
  - Links: [Slide (go-talks)](https://go-talks.appspot.com/github.com/klrkdekira/golang-talks/2019/june/main.slide) | [Source File](2019/june/main.slide)
  - **Topics:** Go 1.12.6 release, GopenPGP encryption library, Go Playground enhancements, compile-time dependency injection with Wire, and community talks.

- **A Short Introduction to Metrics** (April 2019)
  - Links: [Slide (go-talks)](https://go-talks.appspot.com/github.com/klrkdekira/golang-talks/2019/metrics/main.slide) | [Source File](2019/metrics/main.slide) | [Code Samples](2019/metrics/codes/)
  - **Topics:** Application observability fundamentals, Google SRE concepts, time-series data storage, standard `expvar`, instrumenting Go web applications with Prometheus metrics, and dashboarding with Grafana.

- **Golang Meetup April 2019**
  - Links: [Slide (go-talks)](https://go-talks.appspot.com/github.com/klrkdekira/golang-talks/2019/april/main.slide) | [Source File](2019/april/main.slide)
  - **Topics:** Go 1.12.4 update, Go 1.13 error handling preview, Caddy 1.0 release, Jingo fast JSON encoder, overview of standard Go tooling, and community resources.

### 2018

- **Golang Meetup Kuala Lumpur November 2018**
  - Links: [Slide (go-talks)](https://go-talks.appspot.com/github.com/klrkdekira/golang-talks/2018/november/main.slide) | [Source File](2018/november/main.slide)
  - **Topics:** 2018 Go User Survey, Rob Pike on Go 2 draft specifications, GopherLua, visualizing pprof with speedscope, analyzing Docker layers with `dive`, GraphQL-Go, and Go adoption at Mudah.my.

- **Golang File Processing** (A Quick Start to Go)
  - Links: [Slide (go-talks)](https://go-talks.appspot.com/github.com/klrkdekira/golang-talks/2018/file-processing/main.slide) | [Source File](2018/file-processing/main.slide) | [Code Samples](2018/file-processing/codes/)
  - **Topics:** Patterns for high-throughput file processing in Go: `os.Open`, line scanning with `bufio.Scanner`, channel-based buffering, and concurrent producer-consumer pipelines.

- **Building Utilities with Golang**
  - Links: [Slide (go-talks)](https://go-talks.appspot.com/github.com/klrkdekira/golang-talks/2018/cli-tools/main.slide) | [Source File](2018/cli-tools/main.slide) | [Code Samples](2018/cli-tools/codes/)
  - **Topics:** Building and maintaining robust CLI utilities: profiling black-box programs with `net/http/pprof`, identifying CPU and memory bottlenecks (`wasteCycle`), benchmarking, and preventive testing.

### 2017

- **The Show of Profile** (Quick Guide into Golang Profiling)
  - Links: [Slide (go-talks)](https://go-talks.appspot.com/github.com/klrkdekira/golang-talks/2017/profiler/main.slide) | [Source File](2017/profiler/main.slide) | [Code Samples](2017/profiler/codes/)
  - **Topics:** Profiling live Go web applications using `net/http/pprof`, CPU and memory profile capture (`go tool pprof`), interpreting profile graphs, and writing benchmarks and unit tests to prevent performance regressions.

- **Quick Golang** (A Quick Start to Go: Tips and Tricks)
  - Links: [Slide (go-talks)](https://go-talks.appspot.com/github.com/klrkdekira/golang-talks/2017/quick-golang/main.slide) | [Article (go-talks)](https://go-talks.appspot.com/github.com/klrkdekira/golang-talks/2017/quick-golang/main.article) | [Source Slide](2017/quick-golang/main.slide) | [Source Article](2017/quick-golang/main.article) | [Code Samples](2017/quick-golang/codes/)
  - **Topics:** Fast ramp-up on Go: compilation, static linking, cross-compilation, standard tooling (`go fmt`, `godoc`, `go test`), benchmarking, profiling with pprof and expvar, and Prometheus metrics.

### 2014

- **Golang 101 for Programmers**
  - Links: [Slide (go-talks)](https://go-talks.appspot.com/github.com/klrkdekira/golang-talks/2014/golang-101/101.slide) | [Article (go-talks)](https://go-talks.appspot.com/github.com/klrkdekira/golang-talks/2014/golang-101/101.article) | [Source Slide](2014/golang-101/101.slide) | [Source Article](2014/golang-101/101.article) | [Code Snippets](2014/golang-101/snippets/)
  - **Topics:** Introductory guide to Go: design philosophy, syntax basics, variables and types, control flow, functions, closures, pointers, structs, methods, interfaces, defer/panic/recover, goroutines, and channels with concurrency examples.

---

## Repository Structure

```text
golang-talks/
├── 2014/
│   └── golang-101/          # Golang 101 presentation (.slide, .article, code snippets)
├── 2017/
│   ├── profiler/            # Go profiling guide (.slide, code samples)
│   └── quick-golang/        # Quick start guide (.slide, .article, code samples)
├── 2018/
│   ├── cli-tools/           # Building utilities and profiling (.slide, code samples)
│   ├── file-processing/     # Efficient file processing patterns (.slide, code samples)
│   └── november/            # Meetup slides November 2018 (.slide, assets)
├── 2019/
│   ├── april/               # Meetup slides April 2019 (.slide, assets)
│   ├── august/              # Meetup slides August 2019 (.slide, assets)
│   ├── compiler/            # Optimization tips & tricks (.slide, code benchmarks)
│   ├── delve/               # Delve debugger slides (.slide)
│   ├── json/                # JSON processing techniques (.slide, code samples)
│   ├── june/                # Meetup slides June 2019 (.slide, assets)
│   ├── metrics/             # Introduction to metrics (.slide, Prometheus/expvar code)
│   ├── november/            # Meetup slides November 2019 (.slide, assets)
│   ├── october/             # Meetup slides October 2019 (.slide, assets)
│   └── september/           # Meetup slides September 2019 (.slide, assets)
├── 2020/
│   ├── february/            # Meetup slides February 2020 (.slide, assets)
│   ├── october/             # Meetup slides October 2020 (.slide, code samples, assets)
│   └── september/           # Meetup slides September 2020 (.slide, code samples, assets)
├── LICENSE                  # BSD 2-Clause License
├── README.md                # Documentation and catalog
├── go.mod                   # Go module with tool dependency for present
└── go.sum                   # Checksums for tool dependencies
```

---

## License

This repository is licensed under the [BSD 2-Clause License](LICENSE).
