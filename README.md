# This Package is a very simple JSON response getter with Go
The name i chose come from Es gehts which means in german it is ok
easy peasy only import package and then :
sgeht.SReq(url)
only give url as string inside SReq function.
it will retuen for you a json response file
then you can save it like this:
d:=sgeht.SReq(url)
and then print it:
fmt.Println(d)
Easy peasy print the json file :)
