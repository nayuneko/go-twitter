# README #

TwitterAPIアクセス用ラッパーライブラリ
（いまのところリスト取得のみ）

# Usage #

    package main

    import (
      "fmt"
      "os"

      twitter "github.com/nayuneko/go-twitter"
    )

    func main() {
      twitterConfig := twitter.Config{
        ConsumerKey:    "",
        ConsumerSecret: "",
        AccessToken:    "",
        AccessSecret:   "",
      }
      twitterClient := twitter.New(&twitterConfig)
      // API options : https://dev.twitter.com/rest/reference/get/lists/statuses
      opts := map[string]string{
        "list_id":     "",
        "include_rts": "true",
        "count":       "50",
      }
      lists, err := twitterClient.ListsStatuses(opts)
      if err != nil {
        fmt.Println(err)
        os.Exit(1)
      }
      for _, rs := range lists {
        fmt.Println(rs)
      }
    }
