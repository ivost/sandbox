A place for scripts.

## cloudwatch_incoming

Outputs the amount of incoming data to CloudWatch for a specified log group.

Requires the `aws` command-line utility and a valid AWS token.

Example usage:

```
$ ./cloudwatch_incoming "roc_staging"
roc_staging
2019-10-15	25.2947 MB
2019-10-14	207.403 MB
2019-10-13	18.1725 MB
2019-10-12	91.9644 MB
2019-10-11	77119.5 MB
2019-10-10	185.376 MB
2019-10-09	102.353 MB
2019-10-08	88.3494 MB
2019-10-07	1321.47 MB
2019-10-06	19.3199 MB
2019-10-05	65.3586 MB
2019-10-04	54.0941 MB
2019-10-03	1361.39 MB
2019-10-02	696.491 MB
2019-10-01	3578.32 MB
```
