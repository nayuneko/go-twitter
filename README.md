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

    func main() {
      twitterConfig := twitter.twitterConfig{
        ConsumerKey:    "",
        ConsumerSecret: "",
        AccessToken:    "",
        AccessSecret:   "",
      }
      twitterClient := twitter.NewTwitter(twitterConfig)
      // API options : https://dev.twitter.com/rest/reference/get/lists/statuses 
      opts := map[string]string{
        "list_id":     "1234567",
        "include_rts": "true",
        "count":       "50",
      }
      lists, err := twitterClient.ListsStatuses(opts)
      if err != nil {
        fmt.Println(err)
        os.Exit(1)
      }
      for _, receivedStatus := range lists {
        fmt.Printf(receivedStatus)
      }
    }
