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

## Cautions

<b>Never</b> serve a public web service by this tool. It is only for presentation and provides full control of remote code execution.

## Note

GC optimization is currently disabled. Will use parameter to control this.
