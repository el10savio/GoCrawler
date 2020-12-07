# GoCrawler

A distributed asynchronus web crawler implemented using Go, Postgres, RabbitMQ and Docker

## Introduction

A web crawler traverses over a given webpage and find out the links present it. It then repeats the same process for each obtained link recursively indexes a series of pages, thus crawling over the sites. 

Here using a supervisor worker server model we utilize a set of distributed worker nodes to process each page and a supervisor to communicate with the client and aggregate the results. We utilize RabbitMQ to send messages to the worker nodes and Postgres to persist these results.

![Crawler](Crawler.png)


## Steps

To provision the cluster:

```
$ make provision
```

This creates a 3 node 1 supervisor GoCrawler cluster established in their own docker network.

To view the status of the cluster

```
$ make info
```

Now we can send requests to crawl, and view sitemaps to any page.

```
$ curl -X POST -d "{\"URL\": \"<URL>\"}" http://localhost:8050/spider/view
$ curl -X POST -d "{\"URL\": \"<URL>\"}" http://localhost:8050/spider/crawl
```

In the logs for each worker docker container, we can see the logs of the worker nodes parsing and processing each obtained URL.

To tear down the cluster and remove the built docker images:

```
$ make clean
```

## Example

After building the GoCrawler nodes we can now parse/view URLs from the GoSupervisor node

```
$ curl -X POST -d "{\"URL\": \"http://www.google.com\"}" http://localhost:8050/spider/view
$ curl -X POST -d "{\"URL\": \"http://www.google.com\"}" http://localhost:8050/spider/crawl
```

When reading the list of values in they aggregate the visited sitemap and respond with a JSON tree like sitemap.

```
$ curl -X POST -d "{\"URL\": \"http://www.google.com\"}" http://localhost:8050/spider/view
{
  "url": "http://www.google.com//",
  "links": [
    {
      "url": "http://www.google.com/about",
      "links": null
    },
    {
      "url": "http://www.google.com/contact",
      "links": null
    },
    ...
 }
```

We can then instantiate other new spiders to crawl other pages 

```
$ curl -X POST -d "{\"URL\": \"http://www.google.com\"}" http://localhost:8050/spider/view
> 200 OK
```

