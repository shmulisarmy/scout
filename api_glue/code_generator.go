package apiglue

import (
	"log"
	"os"
	"reflect"
	"runtime"
	"strconv"

	"github.com/gin-gonic/gin"
)

func generate_ts_route(file_name string, path_base_name string, params []reflect.Type) string {
	output := ""
	output += "export function " + path_base_name + "("
	for i := range params {
		output += "_" + strconv.Itoa(i) + ": " + typeToTSType(params[i], nil, nil) + ", "
	}
	_, file, line, _ := runtime.Caller(2)
	output += "){\n"
	output += "\t//LINK " + file + ":" + strconv.Itoa(line) + "\n"
	full_fetch_string := path_base_name
	for i := range params {
		full_fetch_string += "/${" + "_" + strconv.Itoa(i) + "}"
	}
	output += "\tfetch(`http://localhost:" + os.Getenv("PORT") + "/" + full_fetch_string + "`, {credentials: 'include'})\n"
	output += "\t.then(response => {\n"
	output += "\tif (response.headers.get(\"sync\")){\n"
	output += "\t\thandle_server_sync(JSON.parse(response.headers.get(\"sync\")))\n"
	output += "\t}\n"
	output += "\treturn response.json()})\n"
	output += "\t.then(data => console.log(data))\n"
	output += "}\n"
	return output
}

var to_add_to_ts_file = "import { handle_server_sync } from \"../apiglue/sync\";\n"

func Make_route(r *gin.Engine, path string, f interface{}) {
	running_status.route_made = true
	//f is a function that takes gc as param 1

	params := []reflect.Type{}
	for param := range reflect.ValueOf(f).Type().NumIn() {
		if param == 0 {
			continue
		}
		params = append(params, reflect.ValueOf(f).Type().In(param))
	}
	ts_code := generate_ts_route(Config.Src_folder+"/routes.ts", path, params)
	for i := range params {
		path += "/:_" + strconv.Itoa(i)
	}
	r.GET(path, func(gc *gin.Context) {
		built_up_params := []reflect.Value{
			reflect.ValueOf(gc),
		}
		for i := range params {
			switch params[i].Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				num, err := strconv.Atoi(gc.Param("_" + strconv.Itoa(i)))
				if err != nil {
					log.Panic("param type not supported: " + params[i].String())
				}
				built_up_params = append(built_up_params, reflect.ValueOf(int(num)))
			case reflect.String:
				built_up_params = append(built_up_params, reflect.ValueOf(gc.Param("_"+strconv.Itoa(i))))
			default:
				log.Panic("param type not supported: " + params[i].String())
			}
		}
		reflect.ValueOf(f).Call(built_up_params)
	})
	to_add_to_ts_file += ts_code

}
