# Mypresent

A redistribution of [present](https://github.com/golang/tools) from golang tools.

## Install

```
go get -u -v github.com/leexun/mypresent
```

## Extra features

### Inline HTML

```
* _
```html
<p class="slogan slogan-15 bold">About me</p>
```htmlend

* _
```html
<p class="slogan slogan-10">First line,</p>
<p class="slogan">second line.</p>
```htmlend
```

### Monaco editor

```
.monaco ./code/example.go
```

## Note

GC optimization is currently disabled. Will use parameter to control this.