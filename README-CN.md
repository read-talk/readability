# Readability 内容分析算法

Readability，它通过遍历Dom对象，通过标签和常用文字的加减权，来重新整合出页面的内容。

# 学习自以下仓库

## [网页文章标题和正文抽取工具](https://github.com/ying32/readability)   

> 最近一年没有新的提交

Golang版本是根据 [readabiliity for node.js](https://github.com/luin/readability) 以及 [readability for python](https://github.com/kingwkb/readability) 所改写，并加入了些自己的，比如支持gzip等。

## [go-readability](https://github.com/mauidude/go-readability) 

> 最近几年没有新的提交

this is library for extracting the main content off of an HTML page. 

This library implements the readability algorithm created by arc90 labs 

and was heavily inspired by https://github.com/cantino/ruby-readability.           

## [Go-Readability](https://github.com/go-shiori/go-readability) 

this is a Go package that find the main readable content and the metadata from a HTML page. 

It works by removing clutter like buttons, ads, background images, script, etc.
     
This package is based from [Readability.js by Mozilla](https://github.com/mozilla/readability), 

and written line by line to make sure it looks and works as similar as possible. 

This way, hopefully all web page that can be parsed by Readability.js are parse-able by go-readability as well.

## [新闻网页正文通用抽取器 Beta 版](https://github.com/kingname/GeneralNewsExtractor)

开发这个项目，源自于我在知网发现了一篇关于自动化抽取新闻类网站正文的算法论文——[《基于文本及符号密度的网页正文提取方法》](https://kns.cnki.net/KCMS/detail/detail.aspx?dbcode=CJFQ&dbname=CJFDLAST2019&filename=GWDZ201908029&v=MDY4MTRxVHJXTTFGckNVUkxPZmJ1Wm5GQ2poVXJyQklqclBkTEc0SDlqTXA0OUhiWVI4ZVgxTHV4WVM3RGgxVDM=)）

这篇论文中描述的算法看起来简洁清晰，并且符合逻辑。但由于论文中只讲了算法原理，并没有具体的语言实现，所以我使用 Python 根据论文实现了这个抽取器。

并分别使用今日头条、网易新闻、游民星空、观察者网、凤凰网、腾讯新闻、ReadHub、新浪新闻做了测试，发现提取效果非常出色，几乎能够达到100%的准确率。

## 其他类似语言的实现

- PHP版本 https://github.com/feelinglucky/php-readability
- Java版本 https://github.com/wuman/JReadability
- Node版本 https://www.npmjs.org/package/node-readability
- Ruby版本 https://github.com/cantino/ruby-readability

