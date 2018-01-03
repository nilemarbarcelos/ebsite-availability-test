[![Go Report Card](https://goreportcard.com/badge/github.com/nilemarbarcelos/website-availability-test)](https://goreportcard.com/report/github.com/nilemarbarcelos/website-availability-test)
[![GoDoc](https://godoc.org/github.com/nilemarbarcelos/website-availability-test?status.svg)](https://godoc.org/github.com/nilemarbarcelos/website-availability-test)


# website-availability-test

* Edit websites.txt and add the websites you want to test

```
--------------------------
1 - Availability test
2 - Show logs
0 - Exit
--------------------------
```

# Triggering tests
```
02/01/2018 22:38:45 Availability testing: 0 https://golang.org/
02/01/2018 22:38:45 https://golang.org/ OK
02/01/2018 22:38:45 Availability testing: 1 https://github.com/
02/01/2018 22:38:46 https://github.com/ OK
02/01/2018 22:38:46 Availability testing: 2 https://www.ign.com/
02/01/2018 22:38:47 https://www.ign.com/ OK
02/01/2018 22:38:47 Availability testing: 3 https://www.cnn.com/
02/01/2018 22:38:47 https://www.cnn.com/ OK
02/01/2018 22:38:47 Availability testing: 4 https://www.foxnews.com/
02/01/2018 22:38:48 https://www.foxnews.com/ OK
02/01/2018 22:38:48 Availability testing: 5 https://www.google.com/
02/01/2018 22:38:50 https://www.google.com/ OK
02/01/2018 22:38:50 Availability testing: 6 https://www.theverge.com/
02/01/2018 22:38:50 https://www.theverge.com/ OK
02/01/2018 22:38:50 Availability testing: 7 https://www.twitter.com/
02/01/2018 22:38:52 https://www.twitter.com/ OK
02/01/2018 22:38:52 Availability testing: 8 https://www.facebook.com/
02/01/2018 22:38:52 https://www.facebook.com/ OK
02/01/2018 22:38:52 Availability testing: 9 https://www.instagram.com/
02/01/2018 22:38:53 https://www.instagram.com/ OK
```

# Showing logs
```
02/01/2018 22:38:45 - https://golang.org/ - OK
02/01/2018 22:38:46 - https://github.com/ - OK
02/01/2018 22:38:47 - https://www.ign.com/ - OK
02/01/2018 22:38:47 - https://www.cnn.com/ - OK
02/01/2018 22:38:48 - https://www.foxnews.com/ - OK
02/01/2018 22:38:50 - https://www.google.com/ - OK
02/01/2018 22:38:50 - https://www.theverge.com/ - OK
02/01/2018 22:38:52 - https://www.twitter.com/ - OK
02/01/2018 22:38:52 - https://www.facebook.com/ - OK
02/01/2018 22:38:53 - https://www.instagram.com/ - OK
```
