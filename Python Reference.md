# Python Reference

## General Language Reference

[[#Collections]]
[[#Documentation]]
[[#Strings]]
[[#Style]]
[[#Types]]


_[[#Python Reference|Back to Top]]_

<br><br><br><br><br>
---

## Specialty Use Cases



_[[#Python Reference|Back to Top]]_

<br><br><br><br><br>
---


### Collections

[[#Lists]]


_[[#Python Reference|Back to Top]]_

<br><br><br><br><br>
---

### Documentation

[[#Retrieving Documentation]]


_[[#Python Reference|Back to Top]]_

<br><br><br><br><br>
---

### Strings

[[#Implementation]]


_[[#Python Reference|Back to Top]]_

<br><br><br><br><br>
---

### Style

[[#Naming]]
[[#Standard]]


_[[#Python Reference|Back to Top]]_

<br><br><br><br><br>
---

### Types



_[[#Python Reference|Back to Top]]_

<br><br><br><br><br>
---


#### Lists

[[#Creating Lists]]


_[[#Python Reference|Back to Top]]_

<br><br><br><br><br>
---


##### Creating Lists

[[#Create a List With Dynamic Content]]


_[[#Python Reference|Back to Top]]_

<br><br><br><br><br>
---



###### Create a List With Dynamic Content

In Python, it's a common idiom to create a list by iterating over a sequence using a technique called _list comprehension_, with the syntax `new_list = [expression] for <item> in <iterable>]` or `newlist = [expression] for <item> in <iterable> if _condition_ == True]`:

_Example:_
```py
newlist = [x for x in range(10) if x < 5]
```

_[[#Python Reference|Back to Top]]_

<br><br><br><br><br>
---

#### Retrieving Documentation

[[#Retrieve Documentation for a Class]]
[[#Retrieve Documentation for a Function]]


_[[#Python Reference|Back to Top]]_

<br><br><br><br><br>
---



###### Retrieve Documentation for a Class

These functions can be called from a script if you pass the call into `print()`, but for ease of use, call these from the REPL/Python Interpretor

_Example:_
```py
help(function_name)
dir(function_name)
```

_[[#Python Reference|Back to Top]]_

<br><br><br><br><br>
---

###### Retrieve Documentation for a Function

These functions can be called from a script if you pass the call into `print()`, but for ease of use, call these from the REPL/Python Interpretor

_Example:_
```py
help(function_name)
type(function_name)
```

_[[#Python Reference|Back to Top]]_

<br><br><br><br><br>
---

#### Implementation

[[#Encoding]]


_[[#Python Reference|Back to Top]]_

<br><br><br><br><br>
---



###### Encoding

By default, Python uses UTF-8 for string encoding

_[[#Python Reference|Back to Top]]_

<br><br><br><br><br>
---

#### Naming

[[#Naming Functions]]
[[#Naming Variables]]


_[[#Python Reference|Back to Top]]_

<br><br><br><br><br>
---



###### Naming Functions

Python function names should be lowercase with words separated by underscores:

_**Example:**_
```py
def my_multi_word_function():
```

_[[#Python Reference|Back to Top]]_

<br><br><br><br><br>
---

###### Naming Variables

Python variable names should be lowercase with words separated by underscores:

_Example:_
```py
my_multi_word_variable = 0
```

_[[#Python Reference|Back to Top]]_

<br><br><br><br><br>
---

###### Standard

[PEP 8](https://www.python.org/dev/peps/pep-0008/)

_[[#Python Reference|Back to Top]]_

<br><br><br><br><br>
---



