In Python, it's a common idiom to create a list by iterating over a sequence using a technique called _list comprehension_, with the syntax `new_list = [expression] for <item> in <iterable>]` or `newlist = [expression] for <item> in <iterable> if _condition_ == True]`:

_Example:_
```py
newlist = [x for x in range(10) if x < 5]
```