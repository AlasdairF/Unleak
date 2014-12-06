##Unleak

Simple function which makes a copy of a slice and returns the copy, so as to avoid potential memory leaks.

It's useful to use this with the output of `bytes.Buffer` if the buffer is to be reset.
