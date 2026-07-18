# 🥬 Apio

> A minimalist HTTP file reader and executor for your terminal.

**apio** is a simple open-source terminal tool designed for developers who prefer defining HTTP requests in plain text files (such as `.http` or `.rest` files) rather than using heavy GUI clients. 

It reads your HTTP files, parses and extracts the requests, and lets you execute them interactively inside a clean Terminal User Interface (TUI). *apio is a reader/executor; it does not modify your files.*

---

## Keybindings

Navigating and running requests in `apio` is straightforward:

| Key | Action |
| :--- | :--- |
| `↓` / `j` | Select the next request |
| `↑` / `k` | Select the previous request |
| `Enter` | Execute the selected request |
| `y` | Copy the response body to clipboard |
| `r` | Reload the HTTP files from disk (on-demand sync) |
| `c` | Clear the response (returns view to request details) |
| `f` | Select a different HTTP/REST file (if multiple files exist) |
| `h` / `?` | Toggle help screen |
| `q` / `Ctrl + C` | Quit apio |

---

## Supported `.http` File Format

`apio` supports standard plain-text HTTP client syntax. For example:

```http
# Global variables
@api = jsonplaceholder.typicode.com
@contentType = application/json

### Get a post
GET https://{{api}}/posts/1
Accept: {{contentType}}

### Create a new post
POST https://{{api}}/posts
Content-Type: {{contentType}}
Token: Bearer your-jwt-token-here

{
  "title": "Testing apio",
  "body": "Sent from the terminal",
  "userId": 1
}
```

---

## Project Status & Limitations

This is a lightweight open-source tool in its early stages:
- **Simple Lexer**: The parser is optimized for standard, everyday syntax (GET, POST, PUT, DELETE, headers, simple variables, comments, and JSON payloads).
- **Edge Cases**: More complex features or syntax from commercial/advanced REST clients may not be fully supported yet.

Contributions, issues, and pull requests are highly welcome to help improve the parser and parser coverage!

---

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## Instalation
Run `./build.sh` to build the binary and `./build.sh install` to install it locally at `~/.local/bin/apio`
(Don't forget to make the file executable if it is not executable by default)

---

### To-do

- [x] Do a better folder structure
- [ ] fix composed variables 
        @hostname=localhost
        @port=44320
        @host={{hostname}}:{{port}}
        GET https://{{host}}/api/search/tool
- [ ] The rest of syntax for the http files
    - [ ] shared variables, composed variables, http protocol in the same line as the request, etc.
        https://learn.microsoft.com/en-us/aspnet/core/test/http-files?view=aspnetcore-10.0

