# README #

TwitterAPIアクセス用ラッパーライブラリ
（いまのところリスト取得のみ）

# Usage #

    package main

    import (
      "bitbucket.org/nayuneko/go-twitter"
      "fmt"
      "os"
    )

    const (
      CONSUMER_KEY        = ""
      CONSUMER_SECRET     = ""
      ACCESS_TOKEN        = ""
      ACCESS_TOKEN_SECRET = ""
    )

    func main() {
      twitter_client := twitter.NewTwitter(CONSUMER_KEY, CONSUMER_SECRET, ACCESS_TOKEN, ACCESS_TOKEN_SECRET)
      // API options : https://dev.twitter.com/rest/reference/get/lists/statuses 
      opts := map[string]string{
        "list_id":     "1234567",
        "include_rts": "true",
        "count":       "50",
      }
      lists, err := twitter_client.ListsStatuses(opts)
      if err != nil {
        fmt.Println(err)
        os.Exit(1)
      }
      ...
    }