package main

import (
	"flag"
	"log"
	"net/http"
)

const html = `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
</head>
<body>
	<h1>CORS Example</h1>
	<div id="output"></div>
	<script>
		document.addEventListener('DOMContentLoaded', async () => {
		    try {
				const response = await fetch('http://localhost:4000/v1/tokens/authentication', {
					method: 'POST',
					headers: {
						'Content-Type': 'application/json',
					},
					body: JSON.stringify({
						email: 'alice@example.com',
						password: 'pa55word',
					})
				});
				const text = await response.text();
				document.getElementById('output').innerText = text;
			} catch (err) {
				document.getElementById('output').innerText = err.message;
			}
		});
	</script>
</body>
</html>
`

func main() {
	addr := flag.String("addr", ":9000", "Server address")
	flag.Parse()

	log.Printf("Starting server on %s", *addr)

	err := http.ListenAndServe(*addr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(html))
	}))

	log.Fatal(err)
}
