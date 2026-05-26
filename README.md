# 🚀 Paper-To-Go (Go Language Playground & Low-Level Sandbox)

A personal learning playground and dev-log to break away from the JS/TS ecosystem and truly master the raw "hands-on" mechanics of Go—including pointers, memory management, generics, and concurrency.

The core philosophy of this repo is: **"Write code that just works first, then rack your brain to refactor it into a 100-point elegant masterpiece."** This space is dedicated to throwing in any spicy low-level topics Go has to offer—performance tuning, custom data structures, memory pools, network parsing, and more.

## 🛠️ Research & Implementation Checklist (ToDo & Done)

### 1. Data Structures & Memory (Manipulating Memory Layouts)
- [x] **Doubly Linked List**: Implemented safe insertions and deletions using `comparable` generics and a clean, symmetrical pointer structure. (Edge case nightmare conquered!)
- [ ] **Stack & Queue**: Compare memory allocation efficiency between slice-based vs. pointer-based implementations.
- [ ] **Object Pool (Memory Pool)**: Reuse memory blocks instead of allocating to the heap with `&Node{}` on every call to reduce Garbage Collection (GC) overhead.
- [ ] **Binary Search Tree (BST)**: Master multi-directional pointer operations and node-deletion algorithms.
- [ ] **HashMap & Set**: Create a custom hash function and handle collisions manually without relying on Go's built-in `map`.

### 2. Storage Engine & Core Database (Under the Hood of NoSQL)
- [ ] **SSTable (Sorted String Table)**: Flush in-memory data sequentially to disk files while maintaining a sorted state.
- [ ] **LSM-Tree**: Build a mini NoSQL engine where memory buffers (MemTable) flush to SSTable files, and merge files in the background using compaction.
- [ ] **Bloom Filter**: Implement ultra-fast in-memory filtering using bitwise operations to check if data exists before scanning heavy disk files.

### 3. Network & Infrastructure (Tearing Down Network Layers)
- [ ] **Primitive HTTP Server**: Bypass the standard `net/http` package entirely. Open raw TCP sockets via `net.Listen`, manually parse HTTP request strings, and return raw responses.
- [ ] **Custom Socket Communication Library**: Design a custom protocol that prefixes data packets with a 4-byte size header to solve TCP packet fragmentation (Sticky Packet issues).
- [ ] **Mini WebSockets**: Upgrade an HTTP connection to a persistent TCP socket and handle full-duplex bi-directional data streams.
- [ ] **Reverse Proxy**: Build a mini Nginx-like feature that receives external requests, forwards them to internal servers, and routes responses back.

### 4. Concurrency & Performance (Squeezing Every Drop of CPU Power)
- [ ] **Worker Pool Pattern**: Leverage Goroutines and Channels to build a safe, pipelined system where a fixed set of workers distribute and process massive tasks concurrently.
- [ ] **Custom Mutex / SpinLock**: Implement a thread-safe lock from scratch using atomic operations (CAS) from the `sync/atomic` package instead of `sync.Mutex`.
- [ ] **Concurrent Map**: Design a high-performance map data structure that minimizes lock contention when multiple Goroutines read/write simultaneously.

### 5. Low-Level & OS Virtualization (Deconstructing the Runtime)
- [ ] **Garbage Collector Simulator**: Mimic Go's Tri-color Marking GC algorithm to track object references and sweep unreferenced memory.
- [ ] **Mini Virtual Machine / Interpreter**: Build a small VM engine that interprets basic bytecodes (e.g., ADD, SUB, PUSH, POP) using a virtual stack and program counters.
- [ ] **Custom String Class & Coding Utils**: Build a custom string struct managing mutable/immutable properties and gather handy utility functions for competitive programming.

## 🧠 Cheat Sheet: Core Rules Learned So Far

### The Laws of Go Pointer Manipulation
1. **Mutate the Actual Target**: Setting `targetNode = nil` only clears the temporary local variable pointer. To actually clear a node from the list, you must directly mutate the source fields: `l.Head = nil` or `l.Tail = nil`.
2. **Symmetrical Architecture**: If your code is drowning in nested `if` blocks, separate the logic into a distinct "front-link step" and "back-link step." Making them symmetrical via `if-else` blocks allows edge cases to be absorbed naturally.
3. **Naming Convention**: Packages
