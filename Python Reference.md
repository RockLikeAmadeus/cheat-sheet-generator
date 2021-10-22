# Python Reference

## General Language Reference

[[#Collections]]
[[#Documentation]]
[[#Strings]]
[[#Style]]
[[#Types]]

---

[[#Python Reference|Back to Top]]

---

## Specialty Use Cases


---

[[#Python Reference|Back to Top]]

---


### Collections

[[#Lists]]

---

[[#Python Reference|Back to Top]]

---

### Documentation

[[#Retrieving Documentation]]

---

[[#Python Reference|Back to Top]]

---

### Strings

[[#Implementation]]

---

[[#Python Reference|Back to Top]]

---

### Style

[[#Naming]]
[[#Standard]]

---

[[#Python Reference|Back to Top]]

---

### Types


---

[[#Python Reference|Back to Top]]

---


#### Lists

[[#Creating Lists]]

---

[[#Python Reference|Back to Top]]

---


##### Creating Lists

[[#Create a List With Dynamic Content]]

---

[[#Python Reference|Back to Top]]

---



###### Create a List With Dynamic Content

In Python, it's a common idiom to create a list by iterating over a sequence using a technique called _list comprehension_, with the syntax `new_list = [expression] for <item> in <iterable>]` or `newlist = [expression] for <item> in <iterable> if _condition_ == True]`:

_Example:_
```py
newlist = [x for x in range(10) if x < 5]
```

[[#Python Reference|Back to Top]]

---

#### Retrieving Documentation

[[#Retrieve Documentation for a Class]]
[[#Retrieve Documentation for a Function]]

---

[[#Python Reference|Back to Top]]

---



###### Retrieve Documentation for a Class

These functions can be called from a script if you pass the call into `print()`, but for ease of use, call these from the REPL/Python Interpretor

_Example:_
```py
help(function_name)
dir(function_name)
```

[[#Python Reference|Back to Top]]

---

###### Retrieve Documentation for a Function

These functions can be called from a script if you pass the call into `print()`, but for ease of use, call these from the REPL/Python Interpretor

_Example:_
```py
help(function_name)
type(function_name)
```

[[#Python Reference|Back to Top]]

---

#### Implementation

[[#Encoding]]

---

[[#Python Reference|Back to Top]]

---



###### Encoding

By default, Python uses UTF-8 for string encoding

[[#Python Reference|Back to Top]]

---

#### Naming

[[#Naming Functions]]
[[#Naming Variables]]

---

[[#Python Reference|Back to Top]]

---



###### Naming Functions

Python function names should be lowercase with words separated by underscores:

_**Example:**_
```py
def my_multi_word_function():
```

[[#Python Reference|Back to Top]]

---

###### Naming Variables

Python variable names should be lowercase with words separated by underscores:

_Example:_
```py
my_multi_word_variable = 0
```

[[#Python Reference|Back to Top]]

---

##### Standard

[PEP 8](https://www.python.org/dev/peps/pep-0008/)

[[#Python Reference|Back to Top]]

---



