## ROC API Load test script

#### Requires working installation of roc cli and vegeta tool 

install from https://github.com/tsenart/vegeta

Configuration file: stress.config

Configuration can be trumped by command line args

```
# for help
./roc_load_test.sh -h 

./roc_load_test.sh [-c connections] [-d maxdevices] [-e environment] [-r QPS] [-t duration]
```

