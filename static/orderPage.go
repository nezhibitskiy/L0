package static

import (
	"L0/internal"
	"fmt"
	"reflect"
)

const header = "<!DOCTYPE html>\n" +
	"<html lang=\"en\">\n" +
	"<head>\n" +
	"<meta charset=\"UTF-8\">\n" +
	"<title>WB L0</title>\n" +
	"<style>\n" +
	".menu {\n" +
	"list-style: none;\n" +
	"padding: 0;\n" +
	"border: 1px solid rgba(0,0,0, .2);\n" +
	"}\n" +
	".menu li {\n" +
	"overflow: hidden;\n" +
	"padding: 6px 10px;\n" +
	"font-size: 20px;\n" +
	"}\n" +
	".menu li:first-child {\n" +
	"font-weight: bold;\n" +
	"padding: 15px 0 10px 15px;\n" +
	"margin-bottom: 10px;\n" +
	"border-bottom: 1px solid rgba(0,0,0, .2);\n" +
	"border-bottom-left-radius: 10px;\n" +
	"border-bottom-right-radius: 10px;\n" +
	"color: #679bb7;\n" +
	"font-size: 24px;\n" +
	"box-shadow: 0 10px 20px -5px rgba(0,0,0, .2);\n" +
	"}\n" +
	".menu li:first-child:before {\n" +
	"content: \"\\2749\";\n" +
	"margin-right: 10px;\n" +
	"}\n" +
	".menu span {\n" +
	"float: left;\n" +
	"width: 75%;\n" +
	"color: #7C7D7F;\n" +
	"}\n" +
	".menu em {\n" +
	"float: right;\n" +
	"color: #9c836e;\n" +
	"font-weight: bold;\n" +
	"}\n" +
	"</style>\n" +
	"</head>\n" +
	"<body>\n" + "\t<form action=\"\" method=\"get\">\n" +
	"ID: <input type=\"text\" name=\"uid\">\n" +
	"<input type=\"submit\" value=\"uid\">\n" +
	"</form>"

func generateList(data interface{}) string {
	val := reflect.ValueOf(data).Elem()
	htmlString := "<ul class=\"menu\">\n"

	for i := 0; i < val.NumField(); i++ {
		htmlString = fmt.Sprint(htmlString+"<li><span>", val.Type().Field(i).Name,
			"</span><em>", val.Field(i).Interface(), "</em></li>\n")
	}
	htmlString += "</ul>\n"
	return htmlString
}

func GeneratePage(order *internal.Order) string {
	return header + generateList(order) + "</body>\n</html>"
}

func GenerateNotFound() string {
	return header + "<h1> Order not found </h1>\n" + "</body>\n</html>"
}
