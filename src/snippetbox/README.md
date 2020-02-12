### Project Structure:
Let’s take a moment to discuss what each of these directories will be used
for.
 
- The cmd directory will contain the application-specific code for the
executable applications in the project. For now we’ll have just one
executable application — the web application — which will live under
the cmd/web directory.


- The pkg directory will contain the ancillary non-application-specific
code used in the project. We’ll use it to hold potentially reusable code
like validation helpers and the SQL database models for the project.

- The ui directory will contain the user-interface assets used by the web
application. Specifically, the ui/html directory will contain HTML
templates, and the ui/static directory will contain static files (like
CSS and images).
So why are we using this structure?

- It scales really nicely if you want to add another executable
application to your project. For example, you might want to add a CLI
(Command Line Interface) to automate some administrative tasks in
the future. With this structure, you could create this CLI application
under cmd/cli and it will be able to import and reuse all the code
you’ve written under the pkg directory.

**So why are we using this structure?
There are two big benefits:

1. It gives a clean separation between Go and non-Go assets. All the Go
code we write will live exclusively under the cmd and pkg directories,
leaving the project root free to hold non-Go assets like UI files,
makefiles and module definitions (including our go.mod file). This can
make things easier to manage when it comes to building and
deploying your application in the future.
2. It scales really nicely if you want to add another executable
application to your project. For example, you might want to add a CLI
(Command Line Interface) to automate some administrative tasks in
the future. With this structure, you could create this CLI application
under cmd/cli and it will be able to import and reuse all the code
you’ve written under the pkg directory.

### HTML Templating and Inheritance :
- A quick note on naming conventions: Throughout this book we’ll use the
  convention <name>.<role>.tmpl for naming template files, where
  <role> is either page, partial or layout. Being able to distinguish the
  role of the template will help us when it comes to creating a cache of
  templates later in the book — but actually it doesn’t matter so much what
  naming convention or file extension you use. Go is flexible about this.

### Template Composition :



