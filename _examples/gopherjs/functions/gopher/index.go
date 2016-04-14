package main

import (
	"github.com/gopherjs/gopherjs/js"
)

// the signature for a lambda handler callback
type lambda_callback func(err, success interface{})

// the signature for a lambda handler
type lambda_handler func(ev js.M, ctx js.M, cb lambda_callback)

// export lambda handler
func exports_handle(handle lambda_handler) {
	js.Module.Get("exports").Set("handle", handle)
}

// log to console
func console_log(parts ...interface{}) {
	js.Global.Get("console").Call("log", parts...)
}

// Starting point for compiling JS code
func main() {
	//
	// exports.handle = function(event, context, callback) {
	//   callback(errorResponse, successResponse);
	// }
	//

	exports_handle(func(event js.M, context js.M, callback lambda_callback) {
		console_log(">>> hello gopher world")
		console_log("event=%j context=%j callback=%j", event, context, callback)

		callback(nil, js.M{"hello": "goodbye"})
	})
}
