# README #

TwitterAPIアクセス用ラッパーライブラリ

# Usage #

    package main

    import (
      "bitbucket.org/nayuneko/go-twitter"
    )

    const (
      CONSUMER_KEY        = ""
      CONSUMER_SECRET     = ""
      ACCESS_TOKEN        = ""
      ACCESS_TOKEN_SECRET = ""
    )

    func main() {
       twitter_client := twitter.NewTwitter(CONSUMER_KEY, CONSUMER_SECRET, ACCESS_TOKEN, ACCESS_TOKEN_SECRET)
    }