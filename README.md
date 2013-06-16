# SEOCMS

SEOCMS is a Search engine optimized CMS, which is developed in Go programming language.

## Dependency

SEOCMS uses following 3rd frameworks:

* [beego](https://github.com/astaxie/beego/)
* [beedb](https://github.com/astaxie/beedb/)
* [Bootstrap](http://twitter.github.io/bootstrap/)

## SQL

    $ mysql -u root -p
    
    CREATE DATABASE seocms CHARACTER SET utf8 COLLATE utf8_general_ci;
    CREATE USER seocms@localhost IDENTIFIED BY 'helloworld';
    GRANT ALL PRIVILEGES ON seocms.* TO seocms@localhost;
    USE seocms;
    CREATE TABLE IF NOT EXISTS category (id INT NOT NULL AUTO_INCREMENT, name CHAR(20) NOT NULL, name_en CHAR(20) NOT NULL, description TEXT, alias CHAR(100), PRIMARY KEY (id));
    CREATE TABLE IF NOT EXISTS tag (id INT NOT NULL AUTO_INCREMENT, name CHAR(20) NOT NULL, name_en CHAR(20), description TEXT, alias CHAR(100), PRIMARY KEY (id));
    CREATE TABLE IF NOT EXISTS article (id INT NOT NULL AUTO_INCREMENT, title CHAR(200) NOT NULL, abstract TEXT, abstract_html TEXT, content TEXT, content_html TEXT, pubdate DATETIME NOT NULL, updated DATETIME NOT NULL, category INT NOT NULL, PRIMARY KEY (id), FOREIGN KEY (category) REFERENCES category(id));
    CREATE TABLE IF NOT EXISTS article_tags (id INT NOT NULL AUTO_INCREMENT, article INT NOT NULL, tag INT NOT NULL, PRIMARY KEY (id), FOREIGN KEY (article) REFERENCES article(id), FOREIGN KEY (tag) REFERENCES tag(id));
    CREATE TABLE IF NOT EXISTS user (id INT NOT NULL AUTO_INCREMENT, email CHAR(100) NOT NULL, name CHAR(50) NOT NULL, password CHAR(100) NOT NULL, created TIMESTAMP DEFAULT '0000-00-00 00:00:00', updated TIMESTAMP DEFAULT NOW() ON UPDATE NOW(), last_login TIMESTAMP, PRIMARY KEY (id));
    CREATE TABLE IF NOT EXISTS link (id INT NOT NULL AUTO_INCREMENT, name CHAR(50) NOT NULL, url CHAR(255) NOT NULL, description TEXT, PRIMARY KEY (id));
    CREATE TABLE IF NOT EXISTS site (id INT NOT NULL AUTO_INCREMENT, name CHAR(20) NOT NULL, content TEXT, PRIMARY KEY (id));

    quit

## Let's go

Build and start:

    $ go build main.go
    $ ./main

Visit <http://localhost:8056/> for article surfing.

Visit <http://localhost:8056/admin/> for article and category administration.

## DEMO

[学车网](http://xueche.haijia.net.cn/)
