#!/usr/bin/env python

# Este script genera la base de datos para la prueba.
# Son alrededor de medio millon de entradas, asi que mejor commitear esto
# que no el .sql que es bastante grande.

from math import floor
from hashlib import sha1
from random import choice
from MySQLdb import connect

LIVE = False

flag = "NCN" + sha1("Acortame esta").hexdigest().lower()
print flag
exit()

base_url = "http://shortener.ctf.noconname.org/"  # poner la barra al final

fd = open("database.sql", "w")
if LIVE:
    db = connect(host="127.0.0.1", user="root", passwd="toor", db="phurl")

script = """DROP DATABASE IF EXISTS phurl;
CREATE DATABASE phurl;
USE phurl;

CREATE TABLE phurl_settings (
  last_number bigint(20) unsigned NOT NULL default '0',
  KEY last_number (last_number)
);

CREATE TABLE phurl_urls (
  id int(10) unsigned NOT NULL auto_increment,
  url text CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  code varchar(20) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL default '',
  alias varchar(40) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL default '',
  date_added datetime NOT NULL default '0000-00-00 00:00:00',
  PRIMARY KEY  (id),
  UNIQUE KEY code (code),
  KEY alias (alias)
);

"""

fd.write(script)
if LIVE:
    c = db.cursor()
    c.execute(script)
    c.close()
    db.commit()

last_number = 1
used_codes = set()

#function generate_code($number) {
#    $out   = "";
#    $codes = "abcdefghjkmnpqrstuvwxyz23456789ABCDEFGHJKMNPQRSTUVWXYZ";
#
#    while ($number > 53) {
#        $key    = $number % 54;
#        $number = floor($number / 54) - 1;
#        $out    = $codes{$key}.$out;
#    }
#
#    return $codes{$number}.$out;
#}
def generate_code(number):
    out   = ""
    codes = "abcdefghjkmnpqrstuvwxyz23456789ABCDEFGHJKMNPQRSTUVWXYZ"
    while number > 53:
        key    = number % 54
        number = int(floor(number / 54)) - 1
        out    = codes[key] + out
    return codes[number] + out

def insert(url, alias = ''):
    global last_number
    number = last_number
    last_number += 1
    code = generate_code(number - 1)
    assert code not in used_codes, code
    used_codes.add(code)
    query = "INSERT INTO phurl_urls VALUES (%d,'%s','%s','%s',NOW());\n" % (number, url, code, alias)
    fd.write(query)
    if LIVE:
        c = db.cursor()
        c.execute(query)
        db.commit()

def random_str():
    codes = "abcdefghjkmnpqrstuvwxyz23456789ABCDEFGHJKMNPQRSTUVWXYZ"
    return ''.join(
        choice(codes)
        for i in xrange(20)
    )

def make_cycles(start, end):
    repetido = set()
    for longitud in xrange(start, end):
        ciclo = []
        for _ in xrange(longitud):
            alias = random_str()
            while alias in repetido or "NCN" in alias:
                alias = random_str()
            repetido.add(alias)
            ciclo.append(alias)
        for indice in xrange(len(ciclo) - 1):
            url = base_url + ciclo[indice + 1]
            insert(url, ciclo[indice])
        url = base_url + ciclo[-1]
        insert(url, ciclo[0])

################################################################################

# las letras
insert('http://ano.lolcathost.org/', 'ano')                         # a
insert('http://boards.4chan.org/b/')                                # b
insert('http://boards.4chan.org/c/')                                # c
insert('http://boards.4chan.org/d/')                                # d
insert('http://boards.4chan.org/e/')                                # e
insert('http://boards.4chan.org/f/')                                # f
insert('http://boards.4chan.org/g/')                                # g
insert('http://boards.4chan.org/h/')                                # h
#insert('http://instagram.com',)                                    # i
insert('http://www.java.com/', 'malware')                           # j
insert('http://krakow.jakdojade.pl/', 'best_city_in_the_world')     # k
#insert('http://lemonparty.org/')                                    # l
insert('http://en.wikipedia.org/wiki/M_(1931_film)')                # m
insert('http://boards.4chan.org/n/')                                # n
#insert('http://boards.4chan.org/o/')                                # o
insert('http://boards.4chan.org/p/')                                # p
insert('http://www.quickmeme.com/', 'meme')                         # q
insert('http://boards.4chan.org/r/')                                # r
insert('http://boards.4chan.org/s/')                                # s
insert('http://boards.4chan.org/t/')                                # t
insert('http://boards.4chan.org/u/')                                # u
insert('http://boards.4chan.org/v/')                                # v
insert('http://boards.4chan.org/w/')                                # w
insert('http://boards.4chan.org/x/')                                # x
insert('http://www.radare.org/y/', 'pancake')                       # y
insert('http://www.imdb.com/title/tt0065234/')                      # z

# las letras que faltan, solo para que les cueste mas deducir el algoritmo
insert('http://boards.4chan.org/i/', 'i')
insert('http://lemonparty.org/', 'l')
insert('http://boards.4chan.org/o/', 'o')

# ciclos de longitud variable, para que se les rompan los scripts
make_cycles(1, 700)

# el flag, mas o menos en el medio pero un toque mas cerca del final,
# aca el timestamp es el menor de todos a proposito
# el schema tambien esta mal a proposito para que destaque,
# y para que no se puedan crear flags falsas
number = last_number
last_number += 1
code = generate_code(number - 1)
assert code not in used_codes, code
used_codes.add(code)
query = "INSERT INTO phurl_urls VALUES (%d,'%s','%s','%s',FROM_UNIXTIME(0)+1);\n" % (number, 'ncn://' + flag, code, flag[3:])
fd.write(query)
if LIVE:
    c = db.cursor()
    c.execute(query)
    db.commit()

# ciclos de longitud variable, para que se les rompan los scripts
make_cycles(701, 1000)

# las del final, para trollear un poco :D
insert('http://www.goatse.info/', 'flag')       # LOL
insert('http://www.meatspin.cc/', 'admin')      # juankeerrrr
insert('http://www.youtube.com/watch?v=fe4fyhzS3UM', 'index.php~')
insert('http://www.youtube.com/watch?v=dQw4w9WgXcQ', 'README')
insert('http://www.youtube.com/watch?v=pIgZ7gMze7A', 'readme')
insert('http://twitter.com/txemaenfurecido', 'robots.txt')
insert('http://177etz3mauip3sw0he3ac6hp.wpengine.netdna-cdn.com/wp-content/uploads/cool-story-bro-500x499.jpg', 'CHANGELOG')
insert('http://www.youtube.com/watch?v=JmcA9LIIXWw', 'changelog')
insert('http://www.youtube.com/watch?v=9sJUDx7iEJw', 'LICENSE')
insert('http://www.youtube.com/watch?v=WGkNiRFwmOg', 'license')
insert('http://www.youtube.com/watch?v=vCH0An-z6XA', 'india')
insert('http://www.youtube.com/watch?v=WniL0xrH-JQ', 'pixel')
insert('http://www.youtube.com/watch?v=pDVORKo8rYs', 'byte')
insert('http://www.youtube.com/watch?v=FXGuOUGttKc', 'putos')
insert('http://www.urbandictionary.com/define.php?term=concha', 'concha')
insert('http://storify.com/nahsagar/opinones-respecto-al-sexo-en-grupo-el-inevitable-e', 'espadeo')
insert('http://www.youtube.com/watch?v=RZ_3MKomQzY', 'bitcoin')
insert('http://www.youtube.com/watch?v=dQ4maOru7Ks', 'nsa')
insert('http://www.youtube.com/watch?v=LJnghGBBP2Q', 'NSA')
insert('http://www.youtube.com/watch?v=CP8CNp-vksc', 'ubuntu')
insert('http://camisetasfrikis.es/', 'remera')

# XXX TODO Las cinco "oficiales" deberian ir con fecha del inicio del CTF
insert('http://truthism.com/', 'aliens')
insert('http://facebook.com/', 'caralibro')
insert('http://instagram.com/', 'hipster')
insert('http://news.ycombinator.com/', 'hacker')
insert('http://www.noconname.org/', 'ncn')

################################################################################

query = "\nINSERT INTO phurl_settings VALUES (%d);\n" % last_number
fd.write(query)
if LIVE:
    c = db.cursor()
    c.execute(query)
    db.commit()

fd.close()
if LIVE:
    db.close()

