#!/bin/sh
exec scala -cp ".:lib/commons-codec-1.10.jar:lib/commons-logging-1.2.jar:lib/fluent-hc-4.5.4.jar:lib/gson-2.8.0.jar:lib/httpclient-4.5.4.jar:lib/httpclient-cache-4.5.4.jar:lib/httpclient-win-4.5.4.jar:lib/httpcore-4.4.7.jar:lib/httpmime-4.5.4.jar:lib/jna-4.4.0.jar:lib/jna-platform-4.4.0.jar" "$0" "$@"
!#

import com.google.gson.Gson
import org.apache.http.client.methods.HttpPost
import org.apache.http.entity.StringEntity
import org.apache.http.impl.client.DefaultHttpClient
import scala.io.Source


case class SlackMessage(text: String, username: String, icon_emoji: String)

val WEBHOOK = args(0)

for(message <- Source.fromFile("../messages.txt").getLines) {

  val msg = new SlackMessage(message, "Mr Shipit", ":shipit:")
  val msgAsJson = new Gson().toJson(msg)
  println(msgAsJson)

  // create an HttpPost object
  val post = new HttpPost(WEBHOOK)

  // set the Content-type
  post.setHeader("Content-type", "application/json")

  // set size of payload
  // post.setHeader("Content-Length", len(msgAsJson))

  // add the JSON as a StringEntity
  post.setEntity(new StringEntity(msgAsJson))

  // send the post request
  val response = (new DefaultHttpClient).execute(post)

  // print the response headers
  println("--- HEADERS ---")
  response.getAllHeaders.foreach(arg => println(arg))

}
