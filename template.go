package main

import "html/template"

var html = template.Must(template.New("webtail").Parse(`
<html>
<head>
    <title>WebTail</title>
    <script src="http://cdn.staticfile.org/jquery/2.0.0/jquery.min.js"></script>
	<script>
	$(document).ready(function() {
		if (!!window.EventSource) {
			var source = new EventSource('/log');
			source.addEventListener('message', function(e) {
				if (e.data.indexOf("error") >= 0) {
					msg = "<div class='msg error'>" + e.data + "</div>";
				} else if (e.data.indexOf("warning") >= 0) {
					msg = "<div class='msg warning'>" + e.data + "</div>";
				} else {
					msg = "<div class='msg'>" + e.data + "</div>";
				}
				$('#messages').append(msg);
				$('html, body').animate({scrollTop:$(document).height()}, 'slow');
			}, false);
		} else {
			alert("NOT SUPPORTED");
		}
	});
    </script>
	<style>
		body {
			background-color: black;
			color: white;
			font:13px/1.4 monaco, "Courier New", Courier, monospace;
			margin: 0px;
			padding: 10px 20px;
		}
		h1 {
			background-color: #222;
			color: greenyellow;
			font-size: 1.2em;
			font-weight: 600;
			position: fixed;
			width: 100%;
			margin: 0;
			top: 0;
			left: 0;
			padding: 5px 20px;
		}
		#messages {
			margin: 30px 0px 10px 0px;
			padding: 10px 0px;
			line-height: 1.5em;
			color: #999;
		}
		div.error { color: red; }
		div.warning { color: orange; }
	</style>
    </head>
    <body>
		<h1>WebTail</h1>
		<div id="messages"></div>
	</body>
</html>
`))
