from http.server import BaseHTTPRequestHandler, HTTPServer
import json
import os
import socket

PORT = int(os.environ.get("PORT", 5000))
APP_ENV = os.environ.get("APP_ENV", "development")

class Handler(BaseHTTPRequestHandler):
    def do_GET(self):
        if self.path == "/health":
            self.send_response(200)
            self.send_header("Content-Type", "application/json")
            self.end_headers()
            response = {"status": "healthy", "service": "python-demo-api"}
            self.wfile.write(json.dumps(response).encode("utf-8"))
            return

        self.send_response(200)
        self.send_header("Content-Type", "application/json")
        self.end_headers()
        payload = {
            "message": "Hello World from Version 2.0 - Fast Cache Rebuild!",
            "environment": APP_ENV,
            "container_hostname": socket.gethostname(),
            "port": PORT
        }
        self.wfile.write(json.dumps(payload, indent=2).encode("utf-8"))

def run():
    server_address = ("", PORT)
    httpd = HTTPServer(server_address, Handler)
    print(f"Server running on port {PORT} [Env: {APP_ENV}]...")
    httpd.serve_forever()

if __name__ == "__main__":
    run()
