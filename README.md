# Simple Local Search Engine Benchmark: Python vs Go

I created this project mainly to evaluate the benefits and costs of using different languages. It’s often said that Python is fast for development but slow at runtime compared to more statically typed languages. I chose Go not because it’s slow to develop in — in fact, it’s quite fast — but simply because I wanted to try it :D

The program is a very simple and inefficient search engine: it splits text into words using whitespace, counts the occurrences of each word in each file, and then sorts the files by frequency. I used it to search for the word “python” in the Python 3 documentation.

Here are the results:

Python

```
hyperfine "python3 main.py '../../../doc/python-3.13-docs-html' 'python'"
Benchmark 1: python3 main.py '../../../doc/python-3.13-docs-html' 'python'
Time (mean ± σ): 5.728 s ± 0.052 s [User: 5.496 s, System: 0.212 s]
Range (min … max): 5.664 s … 5.814 s 10 runs
```

Go with goroutine

```
hyperfine './search_engine "../../../doc/python-3.13-docs-html" "python"'
Benchmark 1: ./search_engine "../../../doc/python-3.13-docs-html" "python"
Time (mean ± σ): 3.856 s ± 0.076 s [User: 9.125 s, System: 2.408 s]
Range (min … max): 3.758 s … 4.020 s 10 runs
```

This result is actually surprising, since the Go version uses goroutines to read files and process them concurrently, but the performance difference isn’t huge. I suspect this is because the hard drive still has to read files sequentially.

The Python code is more concise, but it actually took longer than I expected to write. Python code can be very dense, especially with list comprehensions — if you cram too much logic into one line, it becomes error-prone. Because Python lacks mandatory static typing (you can use type hints, but they’re optional), you have to keep a lot in your head when reading or modifying the code. Also, Python has some [surprises](https://guiyuanju.github.io/blog-ns/2025/06/20/Python-Default-Parameters.html) for a newbie like me, which cost even more time (a skill issue, admittedly).

The Go code was written after the Python version, so the design phase was minimal, making development quite smooth. One notable difference is that when writing Go, I tend to think in smaller, more fragmented units of logic. The code is broken up into smaller pieces compared to the Python version. I also have to handle things like type definitions and error checks, but this actually makes it easier for me to focus on one part of the logic at a time.
