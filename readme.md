# dateconvert/eurodate

Pretty basic golang package to convert from American date formats to European date formats.

## using the code
```
import "github.com/fafrd/dateconvert/eurodate"
```
```
output, err := eurodate.Format("12-13-2014")
if err != nil {
        fmt.Println(err)
} else {
        fmt.Println(output)
}
```

## using the command line
first, get it:
```
go get -v github.com/fafrd/dateconvert
```
usage:
```
dateconvert 12-13-2004
13-12-04

dateconvert 12-13-03
13-12-03

dateconvert 12-13
13-12-18

dateconvert 13-13
parsing time "13-13-2018": month out of range
```

also works with slashes as inputs:
```
dateconvert 12/13
13-12-18
```
